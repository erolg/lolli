package barnard

import (
	"fmt"
	"github.com/layeh/gumble/gumble"
	"net/http"
	"strings"
)

func (b *Barnard) InitApi() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/add-channel/", b.AddChannel)

	mux.HandleFunc("/add-temp-channel/", b.AddTempChannel)

	mux.HandleFunc("/join-channel/", b.JoinChannel)

	mux.HandleFunc("/remove-channel/", b.RemoveChannel)

	mux.HandleFunc("/user-list/", b.UserList)

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

	root := b.Client.Channels[0]

	var targetChannel *gumble.Channel

	if channelName == "Root" && channelName == "root" {
		if root.IsRoot() {
			b.Client.Self.Move(root)
		} else {
			fmt.Fprint(w, root.IsRoot())
		}
	} else {
		channel := strings.Split(channelName, ",")
		if len(channel) == 1 {
			targetChannel = root.Find(channel[0])
		} else if len(channel) == 2 {
			targetChannel = root.Find(channel[0], channel[1])
		}

		b.Client.Self.Move(targetChannel)
		fmt.Fprint(w, channel)
	}

}

func (b *Barnard) RemoveChannel(w http.ResponseWriter, req *http.Request) {
	//	channelName := req.URL.Path[len("/remove-channel/"):]
	channel := b.Client.Self.Channel
	channel.Remove()
}

func (b *Barnard) UserList(w http.ResponseWriter, req *http.Request) {

	channelName := req.URL.Path[len("/user-list/"):]

	var list []string
	var channel *gumble.Channel

	root := b.Client.Channels[0]

	if channelName == "Root" && channelName == "root" {
		if root.IsRoot() {
			channel = root
		} else {
			fmt.Fprint(w, root.IsRoot())
		}
	} else {
		channels := strings.Split(channelName, ",")
		if len(channels) == 1 {
			channel = root.Find(channels[0])
		} else if len(channels) == 2 {
			channel = root.Find(channels[0], channels[1])
		}

	}

	for _, user := range channel.Users {
		list = append(list, user.Name)
	}

	fmt.Fprint(w, list)

}
