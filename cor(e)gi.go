package corgi

import (
	"log"

	"github.com/cein13/corgiEngine/scenemanager"
	"github.com/hajimehoshi/ebiten/v2"
)

type CorgiGame struct {

	screenWidth  int
	screenHeight int
	manager scenemanager.SceneManager
}

var game *CorgiGame

func SetFullscreen (fullscreen bool) {
	ebiten.SetFullscreen(fullscreen)
}

func InitGame(screenWidth int, screenHeight int , sm scenemanager.SceneManager) {

	game = &CorgiGame{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		manager:      sm, 
	}
}

func StartGame(title string,screenWidth int, screenHeight int) {
	
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(title)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}


type Game struct {}

func (g *Game) Update() error {
	game.manager.Update()	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	game.manager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenWidth , screenHeight
}
