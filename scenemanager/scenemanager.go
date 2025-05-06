package scenemanager

import (
	"github.com/Ceinl/corgiEngine/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

var ActiveScene string

type SceneManager struct {
	Scenes []scene.Scene
	ActiveScene *scene.Scene
}

func NewSceneManager(scene scene.Scene) *SceneManager {
	sm := SceneManager{ActiveScene: &scene}	
	sm.AddScene(&scene)

	return &sm
}

func (sm *SceneManager) AddScene(s *scene.Scene) {
	sm.Scenes = append(sm.Scenes, *s)
}

func (sm *SceneManager) Update() {
	if sm.ActiveScene != nil {
		if sm.ActiveScene.Update() != "" {
			sm.SetActiveScene(sm.ActiveScene.Update())
		}
	}
}

func (sm *SceneManager) Draw(screen *ebiten.Image) {
	if sm.ActiveScene != nil {
		sm.ActiveScene.Draw(screen)
	}
}

func (sm *SceneManager) SetActiveScene(name string) {
	for _, s := range sm.Scenes {
		if s.Name == name {
			sm.ActiveScene = &s
			ActiveScene = name
			return
		}
	}
}
