package main

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayerSpritesheets struct {
	downSpriteSheet  Spritesheet
	upSpriteSheet    Spritesheet
	leftSpriteSheet  Spritesheet
	rightSpriteSheet Spritesheet
}

type PlayerAnimations struct {
	downAnimation  Animation
	upAnimation    Animation
	leftAnimation  Animation
	rightAnimation Animation
}

type Player struct {
	sheet          *ebiten.Image
	allSheets      PlayerSpritesheets
	allAnimations  PlayerAnimations
	curSpritesheet Spritesheet
	curAnimation   Animation
	sprite         Sprite //present image based on animDirection/animation
	pos            Vec2i
	dest           Vec2i
	heading        Vec2i
	speed          float64
	animDir        int
	elemBar        elementBar
	spellCap       int
	spellCount     int
	knownSpells    SpellBook
	activeOrbs     string //current string of active orbs, player input to be parsed
	currentOrb     rune   //most recent/last orb
}

func createPlayer(sht *ebiten.Image, position Vec2i) Player {
	spritesheets := PlayerSpritesheets{
		downSpriteSheet:  createSpritesheet(newVec2i(0, 0), newVec2i(80, 79), 2, sht),
		upSpriteSheet:    createSpritesheet(newVec2i(0, 80), newVec2i(80, 159), 2, sht),
		leftSpriteSheet:  createSpritesheet(newVec2i(0, 160), newVec2i(80, 239), 2, sht),
		rightSpriteSheet: createSpritesheet(newVec2i(0, 240), newVec2i(80, 319), 2, sht),
	}
	animations := PlayerAnimations{
		downAnimation:  createAnimation(spritesheets.downSpriteSheet, sht),
		upAnimation:    createAnimation(spritesheets.upSpriteSheet, sht),
		leftAnimation:  createAnimation(spritesheets.rightSpriteSheet, sht),
		rightAnimation: createAnimation(spritesheets.rightSpriteSheet, sht),
	}
	return Player{
		sheet:          sht,
		allSheets:      spritesheets,
		allAnimations:  animations,
		curAnimation:   animations.downAnimation,
		curSpritesheet: spritesheets.downSpriteSheet,
		sprite: Sprite{
			ptA:  Vec2i{0, 0},
			ptB:  Vec2i{40, 80},
			size: Vec2i{40, 80},
		},
		pos:     position, //center of base of character sprite
		dest:    position,
		heading: position,
		speed:   10.0,
		animDir: 0,
		elemBar: elementBar{
			orbs:       "",
			size:       0,
			orbSprites: []Sprite{},
			sheet:      sht,
		},
		spellCount:  0,
		spellCap:    10,
		knownSpells: buildSpellBook(),
		activeOrbs:  "",  //prob redundant due to p.elemBar.orbs !!!!!!!!!!!
		currentOrb:  '0', //0 rune meaning no recorded player input yet?
	}
}

func (p *Player) update() {
	p.curAnimation.update(0.75)                  //parameter is animation speed
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) { //!!!!!!!!!!!!!!clean this up, just one call of addChild, otherwise order so most commonly used keys go first
		p.elemBar.addOrb(ebiten.KeyQ)
		p.currentOrb = Fire
		orbBoard.orbHit = 'F'
	} else if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		p.elemBar.addOrb(ebiten.KeyW)
		p.currentOrb = Water
		orbBoard.orbHit = 'W'
	} else if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		p.elemBar.addOrb(ebiten.KeyE)
		p.currentOrb = Earth
		orbBoard.orbHit = 'E'
	} else if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		p.elemBar.addOrb(ebiten.KeyR)
		p.currentOrb = Lightning
		orbBoard.orbHit = 'L'
	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		//activate spell, or fizzle
		for _, knownSpell := range p.knownSpells.spellList {
			if knownSpell.sequence.MatchString(p.elemBar.orbs) {
				knownSpell.effect()
			}
		}

		p.elemBar.orbs = ""
		p.elemBar.size = 0
		p.currentOrb = '0'
	}
	p.heading = Vec2i{p.dest.x - p.pos.x, p.dest.y - p.pos.y}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		p.dest.x, p.dest.y = ebiten.CursorPosition()
		p.heading = Vec2i{p.dest.x - p.pos.x, p.dest.y - p.pos.y}

		//deciding animDirection of sprite
		if math.Abs(float64(p.heading.x)) > math.Abs(float64(p.heading.y)) { //means either l
			if p.heading.x > 0 { //facing right
				p.animDir = 3
			} else {
				p.animDir = 2 // facing left
			}
		} else if p.heading.y < 0 { //facing up -- REMEMBER LOW Y IS TOP OF SCREEN!!!!!
			p.animDir = 1
		} else {
			p.animDir = 0 // facing down
		}

	}
	headingMag := math.Sqrt(float64(p.heading.x*p.heading.x + p.heading.y*p.heading.y)) //make this an approximation for speed??

	if headingMag > p.speed {
		p.heading.x = int(float64(p.heading.x) * (p.speed / headingMag)) //normalize then mult by desired scale of ten
		p.heading.y = int(float64(p.heading.y) * (p.speed / headingMag)) //	-type as float then int due to rounding to 0
	}

	if (p.pos.x != p.dest.x) || (p.pos.y != p.dest.y) {
		switch p.animDir { //0 down, 1 up, 2 left, 3 right
		case 0:
			p.curAnimation = p.allAnimations.downAnimation
			p.curSpritesheet = p.allSheets.downSpriteSheet
		case 1:
			p.curAnimation = p.allAnimations.upAnimation
			p.curSpritesheet = p.allSheets.upSpriteSheet
		case 2:
			p.curAnimation = p.allAnimations.leftAnimation
			p.curSpritesheet = p.allSheets.leftSpriteSheet
		case 3:
			p.curAnimation = p.allAnimations.rightAnimation
			p.curSpritesheet = p.allSheets.rightSpriteSheet
		}
		p.pos.x += p.heading.x
		p.pos.y += p.heading.y
	}
}

func (p *Player) renderPlayer(screen *ebiten.Image) {
	/*options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(p.pos.x-(p.sprite.size.x/2)), float64(p.pos.y-(p.sprite.size.y)))
	rect := image.Rect(p.sprite.ptA.x, p.sprite.ptA.y, p.sprite.ptB.x, p.sprite.ptB.y) */
	player.elemBar.renderElemBar(screen)
	//screen.DrawImage(p.sheet.SubImage(rect).(*ebiten.Image), options)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.pos.x-(p.curSpritesheet.size.x/(p.curAnimation.spritesheet.numberOfSprites*2))), float64(p.pos.y-(p.curSpritesheet.size.y)))
	op.Filter = ebiten.FilterNearest // Maybe fix rotation grossness?

	// fmt.Printf(p.curSpritesheet.sprites.)
	currentFrame := p.curSpritesheet.sprites[p.curAnimation.currentFrame]
	//currentFrame := p.curAnimation.spritesheet.sprites[p.curAnimation.currentFrame]
	subImageRect := image.Rect(
		currentFrame.ptA.x,
		currentFrame.ptA.y,
		currentFrame.ptB.x,
		currentFrame.ptB.y,
	)

	screen.DrawImage(p.sheet.SubImage(subImageRect).(*ebiten.Image), op) // Draw player
}
