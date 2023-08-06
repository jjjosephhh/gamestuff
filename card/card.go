package card

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/gamestuff/util"
)

type Card struct {
	TextureFront  *rl.Texture2D
	TextureBack   *rl.Texture2D
	Pos           *rl.Vector2
	Width         int32
	Height        int32
	RotFlip       float32
	RotFlipTarget float32
	SpeedFlip     float32
}

func NewCard(
	pathFront, pathBack string,
	position *rl.Vector2,
) *Card {
	textureFront := rl.LoadTexture(pathFront)
	textureBack := rl.LoadTexture(pathBack)
	return &Card{
		TextureFront:  &textureFront,
		TextureBack:   &textureBack,
		Pos:           position,
		Width:         textureFront.Width,
		Height:        textureFront.Height,
		RotFlip:       0,
		RotFlipTarget: 0,
		SpeedFlip:     1000,
	}
}
func (c *Card) Unload() {
	rl.UnloadTexture(*c.TextureFront)
	rl.UnloadTexture(*c.TextureBack)
}

func (c *Card) Draw() {
	dt := rl.GetFrameTime()
	if math.Abs(float64(c.RotFlipTarget)-float64(c.RotFlip)) < float64(c.SpeedFlip*dt) {
		c.RotFlip = c.RotFlipTarget
	}
	if c.RotFlipTarget > c.RotFlip {
		c.RotFlip += c.SpeedFlip * dt
	} else if c.RotFlipTarget < c.RotFlip {
		c.RotFlip -= c.SpeedFlip * dt
	}

	if c.RotFlip < 90 {
		newWidth := float32(c.Width) * (90 - c.RotFlip) / 90
		rl.DrawTextureRec(
			*c.TextureFront,
			rl.NewRectangle(
				float32(c.Width)+(float32(c.Width)-newWidth)/2,
				0,
				newWidth,
				float32(c.Height),
			),
			rl.NewVector2(
				c.Pos.X+(float32(c.Width)-newWidth)/2,
				c.Pos.Y,
			),
			rl.White,
		)
	} else {
		newWidth := float32(c.Width) * (c.RotFlip - 90) / 90
		rl.DrawTextureRec(
			*c.TextureBack,
			rl.NewRectangle(
				(float32(c.Width)-newWidth)/2,
				0,
				newWidth,
				float32(c.Height),
			),
			rl.NewVector2(
				c.Pos.X+(float32(c.Width)-newWidth)/2,
				c.Pos.Y,
			),
			rl.White,
		)
	}
}

func (c *Card) Clicked(posMouse *rl.Vector2) bool {
	if util.Between(posMouse.X, c.Pos.X, c.Pos.X+float32(c.Width)) && util.Between(posMouse.Y, c.Pos.Y, c.Pos.Y+float32(c.Height)) {
		return true
	}
	return false
}

func (c *Card) Flip() {
	if c.RotFlip == 0 && c.RotFlipTarget == 0 {
		c.RotFlipTarget = 180
	} else if c.RotFlip == 180 && c.RotFlipTarget == 180 {
		c.RotFlipTarget = 0
	}
}

func (c *Card) DrawTargetPath(posMouse *rl.Vector2, textureCrosshair *rl.Texture2D) {
	cardCenter := rl.NewVector2(
		c.Pos.X+float32(c.Width)/2,
		c.Pos.Y+float32(c.Height)/2,
	)
	// Calculate the direction vector between the start and end vectors
	direction := rl.Vector2Subtract(*posMouse, cardCenter)
	// Calculate the distance between the start and end vectors
	distance := rl.Vector2Length(direction)
	// Normalize the direction vector
	direction = rl.Vector2Normalize(direction)
	// Loop to calculate and store the intermediate vectors
	var prevVector *rl.Vector2
	for i := float32(0); i <= distance; i += 20 {
		intermediateVector := rl.Vector2Add(cardCenter, rl.Vector2Scale(direction, i))
		if prevVector != nil && int(i/20)%2 == 0 {
			rl.DrawLineEx(
				*prevVector,
				intermediateVector,
				7,
				rl.Pink,
			)
		}
		prevVector = &intermediateVector
	}
	rl.DrawTexture(
		*textureCrosshair,
		int32(posMouse.X)-textureCrosshair.Width/2,
		int32(posMouse.Y)-textureCrosshair.Height/2,
		rl.White,
	)
}
