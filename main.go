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
	"math"
	"regexp"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	SCREENWIDTH  = 800
	SCREENHEIGHT = 600
)

var (
	gameInitialized bool = false
	img             *ebiten.Image
	bg              *ebiten.Image
	bracket         *ebiten.Image
	//offBoard        *ebiten.Image
	//onBoard         *ebiten.Image
	orbBoard OrbBoard
	gameRef  *Game
	player   Player
)

type Game struct {
	player Player
}

func /*(g *Game)*/ init() {
	var err error
	//gameRef = g
	img, _, _ = ebitenutil.NewImageFromFile("assets/playerSheet.png")
	bg, _, err = ebitenutil.NewImageFromFile("assets/bg.png")

	offBoard, _, _ := ebitenutil.NewImageFromFile("assets/orbBoard.png")
	onBoard, _, _ := ebitenutil.NewImageFromFile("assets/orbBoardActivated.png")
	orbBoard = createOrbBoard(offBoard, onBoard)

	bracket, _, _ = ebitenutil.NewImageFromFile("assets/bracket.png")
	player = createPlayer(img, newVec2i(100, 100))

	//example spell
	spellRegexExample := regexp.MustCompile("^LFL$")
	spellExample := Spell{
		sequence: spellRegexExample,
		effect: func() error {
			teleportDist := 300.0
			mposx, mposy := ebiten.CursorPosition()
			heading := Vec2i{mposx - player.pos.x, mposy - player.pos.y}

			//deciding spriteDirection of sprite
			if math.Abs(float64(heading.x)) > math.Abs(float64(heading.y)) { //means either l
				if heading.x > 0 { //facing right
					player.spriteDir = 3
				} else {
					player.spriteDir = 2 // facing left
				}
			} else if heading.y < 0 { //facing up -- REMEMBER LOW Y IS TOP OF SCREEN!!!!!
				player.spriteDir = 1
			} else {
				player.spriteDir = 0 // facing down
			}
			headingMag := math.Sqrt(float64(heading.x*heading.x + heading.y*heading.y))

			if headingMag > teleportDist {
				heading.x = int(float64(heading.x) * (teleportDist / headingMag)) //normalize then mult by desired scale of ten
				heading.y = int(float64(heading.y) * (teleportDist / headingMag)) //	-type as float then int due to rounding to 0
			}

			player.sprite.ptA = newVec2i(player.spriteDir*40, 0)
			player.sprite.ptB = newVec2i((player.spriteDir*40)+40, 80)
			player.pos.x += heading.x
			player.pos.y += heading.y
			player.dest = player.pos
			return nil
		},
	}
	player.knownSpells.appendSpell(spellExample)

	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update() error {
	player.update()
	//orbBoard.update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	orbBoard.renderBoard(screen)
	player.renderPlayer(screen)
	//orbBoard.renderBoard(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
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
