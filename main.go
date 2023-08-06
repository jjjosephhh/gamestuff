package main

import (
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

	crosshair177 := rl.LoadTexture("assets/images/kenney_crosshairPack/PNG/White Retina/crosshair177.png")

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

	posFriendly := rl.NewVector2(float32(screenWidth)/4, float32(screenHeight)/2)
	posEnemy := rl.NewVector2(3*float32(screenWidth)/4, 20)

	cardFriendly := card.NewCard(
		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_version1/Civilian_card_version1_pic1.png",
		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_back/Civilian_card_back.png",
		&posFriendly,
		card.Friendly,
	)
	defer cardFriendly.Unload()
	cardEnemy := card.NewCard(
		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_version1/Civilian_card_version1_pic2.png",
		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_back/Civilian_card_back.png",
		&posEnemy,
		card.Enemy,
	)
	defer cardEnemy.Unload()

	cards := make([]*card.Card, 0)
	cards = append(cards, cardFriendly, cardEnemy)

	var cardSelected *card.Card

	for !rl.WindowShouldClose() {
		posMouse := rl.GetMousePosition()
		// rl.UpdateMusicStream(*songPinkVenom.Music)
		// rl.UpdateMusicStream(*songShutDown.Music)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Beige)

		var cardHovered *card.Card
		for _, c := range cards {
			c.Draw()
			if c.MousedOver(&posMouse) {
				cardHovered = c
			}
		}
		for _, c := range cards {
			if c == cardSelected {
				c.DrawTargetPath(&posMouse, &crosshair177, cardHovered)
			}
		}

		rl.EndMode2D()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			for _, c := range cards {
				if c.MousedOver(&posMouse) {
					if cardSelected == c {
						cardSelected = nil
					} else {
						cardSelected = c
					}
					cardFriendly.Flip()
				}
			}
		}

	}
}
