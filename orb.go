package main

import (
	"image"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//type Orb rune

const (
	// Fire 1
	Fire rune = 'F'
	// 	// Water 2
	Water rune = 'W'
	// 	// Earth 3
	Earth rune = 'E'
	// 	// Lightning 4
	Lightning rune = 'L'
)

type elementBar struct {
	orbs       string
	size       int //number of Orbs
	orbSprites []Sprite
	sheet      *ebiten.Image
}

func createElementBar() elementBar {
	return elementBar{
		orbs:       "",
		size:       0,
		orbSprites: []Sprite{},
		sheet:      &ebiten.Image{},
	}
}

func (e *elementBar) addOrb(key ebiten.Key) {
	if e.size == 0 {
		e.sheet, _, _ = ebitenutil.NewImageFromFile("assets/orb.png")
	}

	var sb strings.Builder
	sb.WriteString(e.orbs)
	switch key {
	case ebiten.KeyQ:
		sb.WriteRune('F')
	case ebiten.KeyW:
		sb.WriteRune('W')
	case ebiten.KeyE:
		sb.WriteRune('E')
	case ebiten.KeyR:
		sb.WriteRune('L')
	}
	e.orbs = sb.String()

	e.orbSprites = append(e.orbSprites, Sprite{
		ptA:  Vec2i{0, 0},
		ptB:  Vec2i{10, 10},
		size: Vec2i{10, 10},
	})
	e.size++

}

func (e *elementBar) renderElemBar(screen *ebiten.Image) {
	bracketRect1 := image.Rect(0, 0, 10, 20)
	bracketRect2 := image.Rect(10, 0, 20, 20)
	orbRect := image.Rect(0, 0, 10, 10)

	bracketY := player.pos.y + 5
	optionsBracket1 := &ebiten.DrawImageOptions{}
	leftBracketX := player.pos.x - 10 - 10*e.size - 2 //-2 to separate bracket some
	//leftBracketX := player.pos.x - 5 - 10*e.size - 2 //-2 to separate bracket some
	optionsBracket1.GeoM.Translate(float64(leftBracketX), float64(bracketY))
	optionsBracket2 := &ebiten.DrawImageOptions{}
	//optionsBracket2.GeoM.Translate(float64(player.pos.x+5+10*e.size+2), float64(bracketY)) //+2 to separate brackets some
	optionsBracket2.GeoM.Translate(float64(player.pos.x+10*e.size+2), float64(bracketY))
	optionsOrb := &ebiten.DrawImageOptions{} //to be modified and used for each of the orbs
	optionsOrb.GeoM.Translate(float64(leftBracketX+10), float64(bracketY))

	for i, orb := range e.orbs { //iterate through each of the orbs in element bar
		if i != 0 {
			optionsOrb.GeoM.Translate(float64(20), 0.0)
		}
		switch orb { //MODIFY THE SPELL TREE HERE no, call update on spell tree in player
		case Fire:
			orbRect = image.Rect(0, 0, 20, 20)
		case Water:
			orbRect = image.Rect(20, 0, 40, 20)
		case Earth:
			orbRect = image.Rect(40, 0, 60, 20)
		case Lightning:
			orbRect = image.Rect(60, 0, 80, 20)
		}
		//optionsOrb.GeoM.Translate(float64(20), 0.0)
		//TODO: now based on type of element render the correct sprite
		//	modifying orbRect, multiply 10 by the number of element
		screen.DrawImage(e.sheet.SubImage(orbRect).(*ebiten.Image), optionsOrb)
	}
	screen.DrawImage(bracket.SubImage(bracketRect1).(*ebiten.Image), optionsBracket1)
	screen.DrawImage(bracket.SubImage(bracketRect2).(*ebiten.Image), optionsBracket2)
}
