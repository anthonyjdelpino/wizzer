package main

import (
	"fmt"
	"regexp"
)

type SpellEffect func() error

type Spell struct {
	sequence *regexp.Regexp
	effect   SpellEffect
}

type SpellBook struct {
	spellList []Spell
	size      int
}

func buildSpellBook() SpellBook {
	spellList := []Spell{}
	fmt.Print("tree initialized!\n")
	return SpellBook{spellList: spellList, size: 0}
}

func (sBook *SpellBook) appendSpell(newSpell Spell) {
	sBook.spellList = append(sBook.spellList, newSpell)
	sBook.size++
	fmt.Printf("spell \"" + newSpell.sequence.String() + "\" appended!\n")
}
