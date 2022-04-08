package main

import (
	"fmt"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Sprite struct {
	ptA Vec2i
	ptB Vec2i

	size Vec2i
}

type Player struct {
	sheet       *ebiten.Image //sprite sheet
	sprite      Sprite        //present image based on spriteDirection/animation
	pos         Vec2i
	dest        Vec2i
	heading     Vec2i
	speed       float64
	spriteDir   int
	elemBar     elementBar
	spellCap    int
	spellCount  int
	knownSpells SpellBook
	activeOrbs  string //current string of active orbs, player input to be parsed
	currentOrb  rune   //most recent/last orb
}

func createPlayer(sht *ebiten.Image, position Vec2i) Player {
	//spellList := spellTree{}
	//fireNode := spellNode{isRoot: false, value: Fire}
	//testSpell.root.addChild(spellNode{isRoot: false, value: Fire})
	//testSpell.root.addChild(spellNode{value: Water})
	return Player{
		sheet: sht,
		sprite: Sprite{
			ptA:  Vec2i{0, 0},
			ptB:  Vec2i{40, 80},
			size: Vec2i{40, 80},
		},
		pos:       position, //center of base of character sprite
		dest:      position,
		heading:   position,
		speed:     10.0,
		spriteDir: 0,
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

func (p *Player) update() { //UPDATE THE PLAYER'S 2 SPELL TREES HERE
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) { //!!!!!!!!!!!!!!clean this up, just one call of addChild, otherwise order so most commonly used keys go first
		p.elemBar.addOrb(ebiten.KeyQ)
		p.currentOrb = Fire

	} else if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		p.elemBar.addOrb(ebiten.KeyW)
		p.currentOrb = Water

	} else if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		p.elemBar.addOrb(ebiten.KeyE)
		p.currentOrb = Earth

		fmt.Printf("orbs: %d\n", p.elemBar.size)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		p.elemBar.addOrb(ebiten.KeyR)
		p.currentOrb = Lightning

	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		//activate spell, or fizzle
		fmt.Printf("cur bar: " + p.elemBar.orbs + "\n")

		for _, knownSpell := range p.knownSpells.spellList {
			fmt.Printf(knownSpell.sequence.String() + "\n")
			if knownSpell.sequence.MatchString(p.elemBar.orbs) {
				fmt.Printf("match!\n")
				knownSpell.effect()
			}
		}

		p.elemBar.orbs = "" //empty string
		p.elemBar.size = 0
		p.currentOrb = '0'
		//p.elemBar.orbs = ""
	}
	p.heading = Vec2i{p.dest.x - p.pos.x, p.dest.y - p.pos.y}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		p.dest.x, p.dest.y = ebiten.CursorPosition()
		p.heading = Vec2i{p.dest.x - p.pos.x, p.dest.y - p.pos.y}

		//deciding spriteDirection of sprite
		if math.Abs(float64(p.heading.x)) > math.Abs(float64(p.heading.y)) { //means either l
			if p.heading.x > 0 { //facing right
				p.spriteDir = 3
			} else {
				p.spriteDir = 2 // facing left
			}
		} else if p.heading.y < 0 { //facing up -- REMEMBER LOW Y IS TOP OF SCREEN!!!!!
			p.spriteDir = 1
		} else {
			p.spriteDir = 0 // facing down
		}

	}
	headingMag := math.Sqrt(float64(p.heading.x*p.heading.x + p.heading.y*p.heading.y)) //make this an approximation for speed??

	if headingMag > p.speed {
		p.heading.x = int(float64(p.heading.x) * (p.speed / headingMag)) //normalize then mult by desired scale of ten
		p.heading.y = int(float64(p.heading.y) * (p.speed / headingMag)) //	-type as float then int due to rounding to 0
	}

	if (p.pos.x != p.dest.x) || (p.pos.y != p.dest.y) {
		p.sprite.ptA = newVec2i(p.spriteDir*40, 0)
		p.sprite.ptB = newVec2i((p.spriteDir*40)+40, 80)
		p.pos.x += p.heading.x
		p.pos.y += p.heading.y
	}
}

func (p *Player) renderPlayer(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(p.pos.x-(p.sprite.size.x/2)), float64(p.pos.y-(p.sprite.size.y)))
	rect := image.Rect(p.sprite.ptA.x, p.sprite.ptA.y, p.sprite.ptB.x, p.sprite.ptB.y)
	player.elemBar.renderElemBar(screen)
	screen.DrawImage(p.sheet.SubImage(rect).(*ebiten.Image), options)
}
