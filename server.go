package barnard

import (
	_ "github.com/layeh/gumble/gumble"
	"net/http"
)

func (b *Barnard) InitApi() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/add-channel/", b.AddChannel)

	mux.HandleFunc("/add-temp-channel/", b.AddTempChannel)

	mux.HandleFunc("/join-channel/", b.JoinChannel)

	mux.HandleFunc("/remove-channel/", b.RemoveChannel)

	return mux
}

func (b *Barnard) AddTempChannel(w http.ResponseWriter, req *http.Request) {
	channelName := req.URL.Path[len("/add-temp-channel/"):]
	if !b.Client.Self.IsRegistered() {
		b.Client.Self.Register()
	}
	channel := b.Client.Self.Channel
	channel.Add(channelName, true)
}

func (b *Barnard) AddChannel(w http.ResponseWriter, req *http.Request) {
	channelName := req.URL.Path[len("/add-channel/"):]
	if !b.Client.Self.IsRegistered() {
		b.Client.Self.Register()
	}
	channel := b.Client.Self.Channel
	channel.Add(channelName, false)
}
func (b *Barnard) JoinChannel(w http.ResponseWriter, req *http.Request) {
	channelName := req.URL.Path[len("/join-channel/"):]

	channel := b.Client.Self.Channel
	found := channel.Find(channelName)

	if found != nil && found != channel {
		b.Client.Self.Move(found)
	}
}

func (b *Barnard) RemoveChannel(w http.ResponseWriter, req *http.Request) {
	//	channelName := req.URL.Path[len("/remove-channel/"):]
	channel := b.Client.Self.Channel
	channel.Remove()
}
