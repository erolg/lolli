package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"

	"../../"
	"github.com/codegangsta/negroni"
	"github.com/layeh/gumble/gumble"
	_ "github.com/layeh/gumble/opus"
      . "github.com/hugozhu/rpi"
        "time"
        "log"

)

func main() {

    WiringPiSetup()

    const (
        LED = PIN_GPIO_3 //22
        PUSHBUTTON   = PIN_GPIO_0 //17
        SWITCH = PIN_GPIO_2 //27
    )



	// Command line flags
	server := flag.String("server", "192.168.0.201:64738", "the server to connect to")
	username := flag.String("username", "lolo", "the username of the client")
	insecure := flag.Bool("insecure", true, "skip server certificate verification")
	certificate := flag.String("certificate", "", "PEM encoded certificate and private key") //server.pem
	key := flag.String("key", "", "PEM encoded certificate and private key") //server.key

	flag.Parse()

        // Initialize
        b := barnard.Barnard{
                Config:  gumble.NewConfig(),
                Address: *server,
        }

    //use default pin naming
    PinMode(LED, OUTPUT)
    PinMode(PUSHBUTTON, INPUT)
    PinMode(SWITCH, INPUT)
    //DigitalWrite(PUSHBUTTON, LOW)
    //DigitalWrite(PUSHBUTTON, LOW)
//    pullUpDnControl(PUSHBUTTON,PUD_DOWN)
//    pullUpDnControl(SWITCH,PUD_DOWN)
 
    DigitalWrite(LED, HIGH)
    Delay(500)
    DigitalWrite(LED, LOW)
    Delay(500)
    DigitalWrite(LED, HIGH)		

        //a goroutine to check button push event
        go func() {
                last_time := time.Now().UnixNano() / 1000000
                btn_pushed := 0
                for pin := range WiringPiISR(PUSHBUTTON, INT_EDGE_BOTH) {
                        if pin > -1 {
				Delay(5)
                                n := time.Now().UnixNano() / 1000000
                                delta := n - last_time
                                if delta > 100 { //software debouncing
                                        if DigitalRead(SWITCH) == LOW { //PTT button mode
						if DigitalRead(PUSHBUTTON) == HIGH { 
                                                	log.Println("button on")				
							b.Client.Self.SetMuted(false)
						//	b.Stream.StartSource()
                                        	}else{
                                                	log.Println("button off")
							b.Client.Self.SetMuted(true)
						}
					}

                                        last_time = n
                                        btn_pushed++

                                }
                        }
                }
        }()


        //a goroutine to check switch event
        go func() {
                last_time := time.Now().UnixNano() / 1000000
                switch_on := 0
		
                for Switchpin := range WiringPiISR(SWITCH, INT_EDGE_BOTH) {
                        if Switchpin > -1 {
				Delay(5)
                                n := time.Now().UnixNano() / 1000000
                                delta := n - last_time
                                if delta > 100 { //software debouncing
                                        if(DigitalRead(SWITCH) == HIGH){  // VOX mode
						log.Println("switch ON")
						//b.Client.Self.SetDeafened(false)
						b.Client.Self.SetMuted(false)
                                                //b.Stream.StartSource()
					}else if (DigitalRead(SWITCH) == LOW){                           //disable VOX mode
						log.Println("switch OFF")
						b.Client.Self.SetMuted(true)
						
					}
                                        last_time = n
                                        switch_on++

                                }
                        }
                }
        }()

	// Initialize
//	b := barnard.Barnard{
//		Config:  gumble.NewConfig(),
//		Address: *server,
//	}

	b.Config.Username = *username

	if *insecure {
		b.TLSConfig.InsecureSkipVerify = true
	}
	if *certificate != "" {
		cert, err := tls.LoadX509KeyPair(*certificate, *key)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		b.TLSConfig.Certificates = append(b.TLSConfig.Certificates, cert)
	}

	mux := b.InitApi()
	b.Api = negroni.Classic()
	b.Api.UseHandler(mux)
	go b.Api.Run(":3000")
	b.Run()

}
