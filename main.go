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

	posFemale := rl.NewVector2(float32(screenWidth)/4, float32(screenHeight)/2)
	posMale := rl.NewVector2(3*float32(screenWidth)/4, float32(screenHeight)/2)

	cardFemale := card.NewCard(
		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_version1/Civilian_card_version1_pic1.png",
		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_back/Civilian_card_back.png",
		&posFemale,
	)
	defer cardFemale.Unload()
	cardMale := card.NewCard(
		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_version1/Civilian_card_version1_pic2.png",
		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_back/Civilian_card_back.png",
		&posMale,
	)
	defer cardMale.Unload()

	var cardSelected *card.Card

	for !rl.WindowShouldClose() {
		posMouse := rl.GetMousePosition()
		// rl.UpdateMusicStream(*songPinkVenom.Music)
		// rl.UpdateMusicStream(*songShutDown.Music)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Beige)

		cardFemale.Draw()
		cardMale.Draw()

		if cardFemale == cardSelected {
			cardFemale.DrawTargetPath(&posMouse, &crosshair177)
		} else if cardMale == cardSelected {
			cardMale.DrawTargetPath(&posMouse, &crosshair177)
		}

		rl.EndMode2D()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if cardFemale.Clicked(&posMouse) {
				if cardSelected == cardFemale {
					cardSelected = nil
				} else {
					cardSelected = cardFemale
				}
				cardFemale.Flip()
				// if songPinkVenom.IsPlaying() {
				// 	songPinkVenom.Stop()
				// } else {
				// 	songPinkVenom.PlayRandom(playTime)
				// }
			}

			if cardMale.Clicked(&posMouse) {
				if cardSelected == cardMale {
					cardSelected = nil
				} else {
					cardSelected = cardMale
				}
				cardMale.Flip()
				// if songShutDown.IsPlaying() {
				// 	songShutDown.Stop()
				// } else {
				// 	songShutDown.PlayRandom(playTime)
				// }
			}
		}

	}
}
