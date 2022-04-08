package main

import (
	"fmt"
	"regexp"
)

//MAKE SURE TO USE regexp Longest() function!!!!

type SpellEffect func() error

type Spell struct {
	sequence *regexp.Regexp
	effect   SpellEffect
}

type SpellBook struct {
	spellList []Spell
	size      int
}

// type Spell struct {
// 	sequence []Orb
// }

// type spellTree struct {
// 	root *spellNode
// 	size int
// 	//height int
// }

// type spellNode struct {
// 	value       Orb
// 	isRoot      bool
// 	isFinal     bool //is this the end of a spell?
// 	isActive    bool //has the player's element list ogne through this node?
// 	parent      *spellNode
// 	children    [4]*spellNode
// 	spellEffect func()
// }

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

// func (parent *spellNode) addChild(child spellNode) { //CHECK FUNCTION SYNTAX, NEED FUNCTION FOR TACKING ONTO TREE, NEED TARGET NODE
// 	child.parent = parent
// 	//parent.children = append(parent.children, &child) //adding the child to the list of parent node's children
// 	if parent.children[child.value] == nil {
// 		parent.children[child.value] = &child
// 	}
// 	//fmt.Print("child added!\n")
// }

// func (node *spellNode) renderTree(screen *ebiten.Image) {
// 	var label string
// 	//fmt.Print("1")
// 	if node == nil {
// 		fmt.Print("nil")
// 		return
// 	}
// 	//fmt.Print("2")
// 	if node.isRoot {
// 		label = "root "
// 	} else {
// 		switch node.value {
// 		case Fire:
// 			label = "fire "
// 		case Water:
// 			label = "water "
// 		case Earth:
// 			label = "earth "
// 		case Lightning:
// 			label = "lightning "
// 		}
// 	}
// 	//ebitenutil.DebugPrint(screen, label)
// 	fmt.Print(label)
// 	//if node.children != nil {
// 	for _, child := range node.children {
// 		if child != nil {
// 			child.renderTree(screen)
// 		}
// 	}
// 	fmt.Print("\n")
// 	//}
// }
