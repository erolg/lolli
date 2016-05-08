package barnard

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

func (b *Barnard) Run() {
	b.Initialize()

	for {
		if b.IsSwitchOn() {
			b.Pushed = true
			time.Sleep(time.Millisecond * 500)
		} else {
			if b.IsPushed() {
				b.Pushed = true
			} else {
				b.Pushed = false
			}
		}
		if b.LastState != b.Pushed {
			b.VoiceToggle()
		}
	}

	b.End()
}

func (b *Barnard) Initialize() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		fmt.Println("Error when open the gpio pins")
		os.Exit(1)
	}

	b.LedPin.High()
	b.PushPin.Input()
	b.SwitchPin.Input()

	b.LastState = false
	b.Pushed = false

	if err := b.start(); err != nil {
		fmt.Println("Error when starting client")
		os.Exit(1)
	} else {
		b.LedPin.High()
	}

}

func (b *Barnard) End() {
	b.Client.Disconnect()
	rpio.Close()
	fmt.Printf("bye")
}

func (b *Barnard) IsSwitchOn() bool {

	if b.SwitchPin.Read() == 1 {
		return true
	} else {
		return false
	}
}

func (b *Barnard) IsPushed() bool {

	if b.PushPin.Read() == 1 {
		return true
	} else {
		return false
	}
}
