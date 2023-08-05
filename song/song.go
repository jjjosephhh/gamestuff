package song

import (
	"context"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Song struct {
	Music                *rl.Music
	Length               float32
	CancelTimedMusicStop *context.CancelFunc
}

func NewSong(filePath string) *Song {
	music := rl.LoadMusicStream(filePath)
	length := rl.GetMusicTimeLength(music)
	return &Song{
		Music:  &music,
		Length: length,
	}
}
func (s *Song) Unload() {
	if s.Music == nil {
		return
	}
	rl.UnloadMusicStream(*s.Music)
}

func (s *Song) IsPlaying() bool {
	if s.Music == nil {
		return false
	}
	return rl.IsMusicStreamPlaying(*s.Music)
}

func (s *Song) Stop() {
	if s.Music == nil {
		return
	}
	rl.StopMusicStream(*s.Music)
	rl.SeekMusicStream(*s.Music, 0)
	if s.CancelTimedMusicStop != nil {
		(*s.CancelTimedMusicStop)()
	}
}

func (s *Song) Start() {
	if s.Music == nil {
		return
	}
	rl.SeekMusicStream(*s.Music, 0)
	rl.PlayMusicStream(*s.Music)
}

func (s *Song) PlayRandom(duration float32) {
	if s.Music == nil {
		return
	}
	rand.Seed(time.Now().UnixNano())
	max := s.Length - duration
	rl.SeekMusicStream(*s.Music, rand.Float32()*max)
	rl.PlayMusicStream(*s.Music)
	ctx, cancel := context.WithCancel(context.Background())
	s.CancelTimedMusicStop = &cancel
	go func(c context.Context) {
		time.Sleep(time.Duration(duration) * time.Second)
		select {
		case <-c.Done():
			return
		default:
			if s.IsPlaying() {
				s.Stop()
			}
		}
	}(ctx)
}
