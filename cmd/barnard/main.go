package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/erolg/lolli"
	"github.com/layeh/barnard/uiterm"
	"github.com/layeh/gumble/gumble"
	_ "github.com/layeh/gumble/opus"
)

func main() {
	// Command line flags
	server := flag.String("server", "localhost:64738", "the server to connect to")
	username := flag.String("username", "lolli", "the username of the client")
	insecure := flag.Bool("insecure", true, "skip server certificate verification")
	certificate := flag.String("certificate", "", "PEM encoded certificate and private key")

	flag.Parse()

	// Initialize
	b := barnard.Barnard{
		Config:  gumble.NewConfig(),
		Address: *server,
	}

	b.Config.Username = *username

	if *insecure {
		b.TLSConfig.InsecureSkipVerify = true
	}
	if *certificate != "" {
		cert, err := tls.LoadX509KeyPair(*certificate, *certificate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		b.TLSConfig.Certificates = append(b.TLSConfig.Certificates, cert)
	}
	mux := initializeApiMethods()
	b.Api = negroni.Classic()
	b.Api.UseHandler(mux)
	go b.Api.Run(":3000")
	b.Ui = uiterm.New(&b)
	b.Ui.Run()
}
