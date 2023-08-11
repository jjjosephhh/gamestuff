package main

import (
	"fmt"

	rgui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/gamestuff/card"
)

const screenWidth int32 = 1000
const screenHeight int32 = 1000

// const playTime float32 = 3

func main() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - models loading")
	defer rl.CloseWindow()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)

	// crosshair177 := rl.LoadTexture("assets/images/kenney_crosshairPack/PNG/White Retina/crosshair177.png")

	// songPinkVenom := song.NewSong("assets/audio/pink-venom.mp3")
	// defer songPinkVenom.Unload()
	// songShutDown := song.NewSong("assets/audio/shut-down.mp3")
	// defer songShutDown.Unload()

	camera := rl.NewCamera2D(
		rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2), // Camera offset
		rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2), // Camera offset
		0.0, // Rotation angle in degrees (no rotation initially)
		1.0, // Zoom level (normal zoom initially)
	)

	// posFriendly := rl.NewVector2(float32(screenWidth)/4, float32(screenHeight)/2)
	// posEnemy := rl.NewVector2(3*float32(screenWidth)/4, 20)

	// cardFriendly := card.NewCard(
	// 	"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_version1/Civilian_card_version1_pic1.png",
	// 	"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_back/Civilian_card_back.png",
	// 	&posFriendly,
	// 	card.Friendly,
	// )
	// defer cardFriendly.Unload()

	// cardEnemy := card.NewCard(
	// 	"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_version1/Civilian_card_version1_pic2.png",
	// 	"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_back/Civilian_card_back.png",
	// 	&posEnemy,
	// 	card.Enemy,
	// )
	// defer cardEnemy.Unload()

	hand := CreateHand()
	defer UnloadHand(hand)

	// cards = append(cards, cardFriendly, cardEnemy)

	var cardSelected *card.Card

	btnClicked := false

	for !rl.WindowShouldClose() {
		posMouse := rl.GetMousePosition()
		// rl.UpdateMusicStream(*songPinkVenom.Music)
		// rl.UpdateMusicStream(*songShutDown.Music)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Beige)

		// var cardHovered *card.Card
		RepositionHand(hand)
		for _, c := range hand {
			c.Draw()
			// if c.MousedOver(&posMouse) {
			// 	cardHovered = c
			// }
		}

		// for _, c := range cards {
		// 	if c == cardSelected {
		// 		c.DrawTargetPath(&posMouse, &crosshair177, cardHovered)
		// 	}
		// }

		rl.EndMode2D()

		btnClicked = rgui.Button(rl.NewRectangle(float32(screenWidth/2-40), float32(screenHeight/2-20), 80, 40), "Click Me!")

		rl.DrawFPS(10, 10)

		rl.EndDrawing()

		if rl.IsKeyPressed(rl.KeyQ) {
			if cardSelected != nil {
				cardSelected = nil
			}
		}

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			for i := len(hand) - 1; i >= 0; i-- {
				c := hand[i]
				if c.MousedOver(&posMouse) {
					if cardSelected == c {
						cardSelected = nil
					} else {
						cardSelected = c
					}
					c.Flip()
					break
				}
			}
		}

		if btnClicked {
			btnClicked = false
			fmt.Println("Button was clicked!", len(hand))
			hand = AppendToHand(hand)
		}
	}
}

func CreateHand() []*card.Card {
	cards := make([]*card.Card, 0)
	for i := 0; i < 4; i++ {
		pos := rl.NewVector2(float32(100*i), float32(screenHeight))
		c := card.NewCard(
			"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_version1/Civilian_card_version1_pic1.png",
			"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_back/Civilian_card_back.png",
			&pos,
			card.Friendly,
		)

		c.Pos.Y -= float32(c.Height)
		c.Pos.Y += 50

		cards = append(cards, c)
	}
	return cards
}

func AppendToHand(hand []*card.Card) []*card.Card {
	pos := rl.NewVector2(float32(100*len(hand)), float32(screenHeight))
	c := card.NewCard(
		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_version1/Civilian_card_version1_pic1.png",
		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_back/Civilian_card_back.png",
		&pos,
		card.Friendly,
	)

	c.Pos.Y -= float32(c.Height)
	c.Pos.Y += 50
	return append(hand, c)
}

func RepositionHand(hand []*card.Card) {
	if len(hand) < 2 {
		return
	}

	distCards := float32(len(hand)) * float32(hand[0].Width)
	if distCards > float32(screenWidth) {
		offset := (float32(screenWidth) - float32(hand[0].Width)) / float32(len(hand)-1)
		for i, c := range hand {
			c.Pos.X = float32(i) * offset
		}
	} else {
		for i, c := range hand {
			c.Pos.X = float32(i) * float32(hand[0].Width)
		}
	}
}

func UnloadHand(hand []*card.Card) {
	for _, c := range hand {
		c.Unload()
	}
}
