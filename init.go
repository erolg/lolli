package barnard

import (
	"fmt"
//	"github.com/stianeikeland/go-rpio"
	"os"
//	"time"
	"sync"
)


func (b *Barnard) Run() {
	b.Initialize()	
	b.Stream.StartSource()

	var wg sync.WaitGroup // wait for the routines
	wg.Add(1)

	wg.Wait()
	
//        for {
  //            time.Sleep(time.Millisecond * 1000)
    //    }

//	b.End()
}

func (b *Barnard) Initialize() {

	if err := b.start(); err != nil {
		fmt.Println("Error when starting client")
		os.Exit(1)
	} else {
		
		
	}

}

func (b *Barnard) End() {
	b.Client.Disconnect()
//	rpio.Close()
	fmt.Printf("bye")
}


