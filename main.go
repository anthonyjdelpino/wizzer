package main

/*TODO
-add spell functions
-add spell updates
-2nd spell list, to revert to spells known
-spell addition
*/
import (
	_ "image/png"
	"log"
	"regexp"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	SCREENWIDTH  = 640
	SCREENHEIGHT = 480
)

var (
	gameInitialized bool = false
	img             *ebiten.Image
	bg              *ebiten.Image
	bracket         *ebiten.Image
	gameRef         *Game
	player          Player
)

type Game struct {
	player Player
}

func /*(g *Game)*/ init() {
	var err error
	//gameRef = g
	img, _, _ = ebitenutil.NewImageFromFile("playerSheet.png")
	bg, _, err = ebitenutil.NewImageFromFile("bg.png")
	bracket, _, _ = ebitenutil.NewImageFromFile("bracket.png")
	player = createPlayer(img, newVec2i(100, 100))

	//example spell
	spellRegexExample := regexp.MustCompile("^FLF$")
	spellExample := Spell{
		sequence: spellRegexExample,
		effect: func() error {
			player.pos = newVec2i(200, 200)
			return nil
		},
	}
	player.knownSpells.appendSpell(spellExample)

	spellRegexExample2 := regexp.MustCompile("^FWEL$")
	spellExample2 := Spell{
		sequence: spellRegexExample2,
		effect: func() error {
			player.pos = newVec2i(500, 300)
			return nil
		},
	}

	player.knownSpells.appendSpell(spellExample2)

	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update() error {
	player.update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	player.renderPlayer(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {
	ebiten.SetWindowSize(1200, 900)
	ebiten.SetWindowTitle("Wizzer DEV")
	if err := ebiten.RunGame(&Game{
		player: player,
	}); err != nil {
		log.Fatal(err)
	}
}
