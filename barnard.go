package barnard

import (
	"crypto/tls"

	"github.com/codegangsta/negroni"
	"github.com/layeh/barnard/uiterm"
	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleopenal"
	"github.com/stianeikeland/go-rpio"
)

type Barnard struct {
	Config *gumble.Config
	Client *gumble.Client

	Address   string
	TLSConfig tls.Config

	Stream *gumbleopenal.Stream

	Api *negroni.Negroni

	Pushed    bool
	LastState bool
	PushPin   rpio.Pin
	LedPin    rpio.Pin
	SwitchPin rpio.Pin
}
