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
	username := flag.String("username", "lool", "the username of the client")
	insecure := flag.Bool("insecure", true, "skip server certificate verification")
	certificate := flag.String("certificate", "server.pem", "PEM encoded certificate and private key")
	key := flag.String("key", "server.key", "PEM encoded certificate and private key")

	flag.Parse()

        // Initialize
        b := barnard.Barnard{
                Config:  gumble.NewConfig(),
                Address: *server,
        }

    //use default pin naming
    PinMode(LED, OUTPUT)
    DigitalWrite(LED, HIGH)
    Delay(500)
    DigitalWrite(LED, LOW)

        //a goroutine to check button push event
        go func() {
                last_time := time.Now().UnixNano() / 1000000
                btn_pushed := 0
                for pin := range WiringPiISR(PUSHBUTTON, INT_EDGE_RISING) {
                        if pin > -1 {
                                n := time.Now().UnixNano() / 1000000
                                delta := n - last_time
                                if delta > 400 { //software debouncing
                                        if(DigitalRead(PUSHBUTTON) == HIGH){
                                                log.Println("button on")
						b.Stream.StartSource()
                                        }else{
                                                log.Println("button off")
						b.Stream.StopSource()
                                        }

                                        last_time = n
                                        btn_pushed++

                                }
                        }
                }
        }()

        //a goroutine to check switch event
        go func() {
                Slast_time := time.Now().UnixNano() / 1000000
                switch_on := 0
                for Switchpin := range WiringPiISR(SWITCH, INT_EDGE_RISING) {
                        if Switchpin > -1 {
                                Sn := time.Now().UnixNano() / 1000000
                                Sdelta := Sn - Slast_time
                                if Sdelta > 400 { //software debouncing
                                        if(DigitalRead(SWITCH) == HIGH){
						log.Println("switch on")
						b.Stream.StartSource()
					}else{
						log.Println("switch off")
						b.Stream.StopSource()
					}
                                        Slast_time = Sn
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
