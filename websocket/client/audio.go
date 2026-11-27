package client

import (
	"bytes"
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

var (
	otoCtx     *oto.Context
	otoOnce    sync.Once
	otoInitErr error
)

func PlaySound(ctx context.Context, path string) {
	fB, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	mp3, err := mp3.NewDecoder(bytes.NewReader(fB))
	if err != nil {
		panic(err)
	}

	otoOnce.Do(func() {
		op := &oto.NewContextOptions{
			SampleRate:   mp3.SampleRate(),
			ChannelCount: 2,
			Format:       oto.FormatSignedInt16LE,
		}
		var readyChan chan struct{}
		otoCtx, readyChan, otoInitErr = oto.NewContext(op)
		if otoInitErr == nil {
			<-readyChan
		}
	})

	if otoInitErr != nil || otoCtx == nil {
		log.Println("oto contexte error :", otoInitErr)
		return
	}

	player := otoCtx.NewPlayer(mp3)
	player.Play()

	tick := time.NewTicker(50 * time.Millisecond)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			player.Pause()
			player.Close()
			return
		case <-tick.C:
			if !player.IsPlaying() {
				player.Close()
				return
			}
		}
	}
}
