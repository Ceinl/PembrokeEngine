package corgi

import (
	"log"

	"github.com/Ceinl/PembrokeEngine/scenemanager"
	"github.com/hajimehoshi/ebiten/v2"
)

type CorgiGame struct {
	screenWidth  int
	screenHeight int
	manager      scenemanager.SceneManager
}

var game *CorgiGame

func SetFullscreen(fullscreen bool) {
	ebiten.SetFullscreen(fullscreen)
}

func InitGame(screenWidth int, screenHeight int) {

	game = &CorgiGame{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		manager:      scenemanager.SceneManager{},
	}
}

func Run(title string) {

	ebiten.SetWindowSize(game.screenWidth, game.screenHeight)
	ebiten.SetWindowTitle(title)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}

type Game struct{}

func (g *Game) Update() error {
	game.manager.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	game.manager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}
