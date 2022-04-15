package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type OrbBoard struct {
	onSheet  *ebiten.Image
	offSheet *ebiten.Image
	orbHit   rune //what to highlight in rendering, means being hit on keyboard
}

//func (board *orbBoard) renderBoard(screen *ebiten.Image)

func createOrbBoard(off, on *ebiten.Image) OrbBoard {
	return OrbBoard{
		onSheet:  on,
		offSheet: off,
		orbHit:   ' ',
	}
}

// update func
func (o *OrbBoard) renderBoard(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(2.0, 2.0)
	options.GeoM.Translate(0, SCREENHEIGHT-80)

	//rectOff := image.Rect(0, 20, 80, 40)

	//screen.DrawImage(o.offSheet.SubImage(rectOff).(*ebiten.Image), options)
	screen.DrawImage(o.offSheet, options)

	switch o.orbHit {
	case '0':
		break
	case 'F':
		rectF := image.Rect(0, 0, 20, 40)
		optionsF := &ebiten.DrawImageOptions{}
		optionsF.GeoM.Scale(2.0, 2.0)
		optionsF.GeoM.Translate(0, SCREENHEIGHT-80)
		screen.DrawImage(o.onSheet.SubImage(rectF).(*ebiten.Image), optionsF)
	case 'W':
		rectW := image.Rect(20, 0, 40, 40)
		optionsW := &ebiten.DrawImageOptions{}
		optionsW.GeoM.Scale(2.0, 2.0)
		optionsW.GeoM.Translate(40, SCREENHEIGHT-80)
		screen.DrawImage(o.onSheet.SubImage(rectW).(*ebiten.Image), optionsW)
	case 'E':
		rectE := image.Rect(40, 0, 60, 40)
		optionsE := &ebiten.DrawImageOptions{}
		optionsE.GeoM.Scale(2.0, 2.0)
		optionsE.GeoM.Translate(80, SCREENHEIGHT-80)
		screen.DrawImage(o.onSheet.SubImage(rectE).(*ebiten.Image), optionsE)
	case 'L':
		rectL := image.Rect(60, 0, 80, 40)
		optionsL := &ebiten.DrawImageOptions{}
		optionsL.GeoM.Scale(2.0, 2.0)
		optionsL.GeoM.Translate(120, SCREENHEIGHT-80)
		screen.DrawImage(o.onSheet.SubImage(rectL).(*ebiten.Image), optionsL)
	}
	o.orbHit = '0'
}
