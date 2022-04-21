package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite stores information on where a sprite subimage is located
type Sprite struct {
	ptA  Vec2i
	ptB  Vec2i
	size Vec2i
}

func createSprite(ptA Vec2i, ptB Vec2i, size Vec2i, image *ebiten.Image) Sprite {
	return Sprite{ptA, ptB, size}
}

func (s *Sprite) getBounds() image.Rectangle {
	return image.Rect(s.ptA.x, s.ptA.y, s.ptB.x, s.ptB.y)
}

// Spritesheet is a collection of Sprite's
type Spritesheet struct {
	sprites         []Sprite
	numberOfSprites int
	ptA             Vec2i
	size            Vec2i
}

func createSpritesheet(ptA Vec2i, ptB Vec2i, numberOfSprites int, image *ebiten.Image) Spritesheet {
	// Create sprite slice
	sprites := make([]Sprite, numberOfSprites)
	// Calculate size of entire sheet
	size := newVec2i(ptB.x-ptA.x, ptB.y-ptA.y)
	// Calculate size of single sprite
	spriteSize := newVec2i(size.x/numberOfSprites, size.y)
	for i := 0; i < numberOfSprites; i++ { //THIS MOVES ACROSS THE SPRITESHEET TO FORM ANIMATION (L to R)
		// Calculate start and end positions of current sprite
		spriteptA := newVec2i(
			i*spriteSize.x+ptA.x,
			ptA.y,
		)
		spriteptB := newVec2i(
			i*spriteSize.x+ptA.x+spriteSize.x,
			spriteSize.y+ptA.y,
		)
		// Add sprite to slice
		sprites[i] = Sprite{spriteptA, spriteptB, spriteSize}
	}
	return Spritesheet{
		sprites,
		numberOfSprites,
		ptA,
		size,
	}
}

func (s *Spritesheet) ptB() Vec2i {
	return newVec2i(s.ptA.x+s.size.x, s.ptA.y+s.size.y)
}

func (s *Spritesheet) getBounds() image.Rectangle {
	return image.Rect(
		s.ptA.x,
		s.ptA.y,
		s.ptA.x+s.size.x,
		s.ptA.y+s.size.y,
	)
}
