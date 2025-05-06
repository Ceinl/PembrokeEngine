package scene

import (
	"github.com/Ceinl/CorgiEngine/gameobject"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	Name    string
	Objects []*gameobject.GameObject
}

func NewScene(name string) *Scene {
	return &Scene{
		Name:    name,
		Objects: make([]*gameobject.GameObject, 0),
	}
}

func (s *Scene) Update() string {
	for _, o := range s.Objects {
		if o.IsActive {
			return o.Act.Update()
		}
	}
	return ""
}

func (s *Scene) Draw(screen *ebiten.Image) {
	for _, o := range s.Objects {
		img, opts := o.Sprite.Get()
		if o.IsVisible || img != nil || opts != nil {
			o.Act.Draw(screen)
		}
	}
}

func (s *Scene) Init() {
	for _, o := range s.Objects {
		o.Act.Init()
	}
}
