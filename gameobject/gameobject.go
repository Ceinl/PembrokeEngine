package gameobject

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject struct {
	Position  Position 
	Sprite    Sprite   

	IsVisible bool
	IsActive  bool

	Act Action
}

func NewGameObject(action Action) *GameObject {
	return &GameObject{
		IsVisible: true,
		IsActive:  true,
		Act:       action,
	}
}

// Position block
type Position struct {
	X *float64
	Y *float64
	Z *int
}

//

// Sprite block
type Sprite struct {
	img  *ebiten.Image
	opts *ebiten.DrawImageOptions
}

func (s *Sprite) Set(img *ebiten.Image, opts *ebiten.DrawImageOptions) {
	s.img = img
	s.opts = opts
}

func (s *Sprite) Get() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return s.img, s.opts
}

func(o *GameObject) StaticDraw(screen *ebiten.Image) {
	img, opts := o.Sprite.Get()
	opts.GeoM.Reset()
	if o.Position.X != nil && o.Position.Y != nil { 
		opts.GeoM.Translate(*o.Position.X, *o.Position.Y)
	}
	screen.DrawImage(img, opts)
}

//

// Action block
type Action interface {
	Draw(image *ebiten.Image)
	Update() string
	Init()
}

//
