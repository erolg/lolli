package barnard

import (
	"fmt"
//	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)


func (b *Barnard) Run() {
	b.Initialize()	
	
        for {
              time.Sleep(time.Millisecond * 1000)
        }

//	b.End()
}

func (b *Barnard) Initialize() {

	if err := b.start(); err != nil {
		fmt.Println("Error when starting client")
		os.Exit(1)
	} else {
		//b.LedPin.High()
		//b.Stream.StartSource()
		//b.LedPin.High()
		
	}

}

func (b *Barnard) End() {
	b.Client.Disconnect()
//	rpio.Close()
	fmt.Printf("bye")
}


