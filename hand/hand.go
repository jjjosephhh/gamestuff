package hand

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/gamestuff/card"
	"github.com/jjjosephhh/gamestuff/constants"
)

type Hand struct {
	Cards []*card.Card
}

func NewHand() *Hand {
	cards := make([]*card.Card, 0)
	return &Hand{
		Cards: cards,
	}
	// for i := 0; i < 4; i++ {
	// 	pos := rl.NewVector2(float32(100*i), float32(screenHeight))
	// 	c := card.NewCard(
	// 		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_version1/Civilian_card_version1_pic1.png",
	// 		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_back/Civilian_card_back.png",
	// 		&pos,
	// 		card.Friendly,
	// 	)

	// 	c.Pos.Y -= float32(c.Height)
	// 	c.Pos.Y += 50

	// 	cards = append(cards, c)
	// }
	// return cards
}

func (hand *Hand) Append() {
	pos := rl.NewVector2(
		float32(constants.ScreenWidth),
		float32(constants.ScreenHeight),
	)
	posTarget := rl.NewVector2(
		float32(constants.ScreenWidth),
		float32(constants.ScreenHeight),
	)
	c := card.NewCard(
		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_version1/Civilian_card_version1_pic1.png",
		"assets/images/free-npc-quest-tcg-cards-pixel-art/PNG/Cards_color1/Civilian_card_back/Civilian_card_back.png",
		&pos,
		&posTarget,
		card.Friendly,
	)
	c.SetCardHandY()
	hand.Cards = append(hand.Cards, c)
	hand.Reposition()
	for _, c := range hand.Cards {
		fmt.Println("------------------------------------")
		fmt.Println("Card Pos", c.Pos.X, c.Pos.Y)
		fmt.Println("Card PosTarget", c.PosTarget.X, c.PosTarget.Y)
	}
}

func (hand *Hand) Reposition() {
	if len(hand.Cards) < 1 {
		return
	}

	distCards := float32(len(hand.Cards)) * float32(hand.Cards[0].Width)
	if distCards > float32(constants.ScreenWidth) {
		offset := (float32(constants.ScreenWidth) - float32(hand.Cards[0].Width)) / float32(len(hand.Cards)-1)
		for i, c := range hand.Cards {
			c.PosTarget.X = float32(i) * offset
		}
	} else {
		for i, c := range hand.Cards {
			c.PosTarget.X = float32(i) * float32(hand.Cards[0].Width)
		}
	}
}

func (hand *Hand) Draw() {
	for _, c := range hand.Cards {
		c.Move()
		c.Draw()
		// if c.MousedOver(&posMouse) {
		// 	cardHovered = c
		// }
	}
}

func (hand *Hand) Unload() {
	for _, c := range hand.Cards {
		c.Unload()
	}
}
