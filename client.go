package barnard

import (
	"fmt"
	"net"
	"os"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleopenal"
	"github.com/layeh/gumble/gumbleutil"
)

func (b *Barnard) start() error {
	b.Config.Attach(gumbleutil.AutoBitrate)
	b.Config.Attach(b)

	var err error
	_, err = gumble.DialWithDialer(new(net.Dialer), b.Address, b.Config, &b.TLSConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err
	}

	// Audio
	if os.Getenv("ALSOFT_LOGLEVEL") == "" {
		os.Setenv("ALSOFT_LOGLEVEL", "0")
	}
	if stream, err := gumbleopenal.New(b.Client); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err
	} else {
		b.Stream = stream
		return nil
	}
}

func (b *Barnard) OnConnect(e *gumble.ConnectEvent) {
	b.Client = e.Client

	fmt.Printf("To: %s", e.Client.Self.Channel.Name)
}

func (b *Barnard) OnDisconnect(e *gumble.DisconnectEvent) {
	var reason string
	switch e.Type {
	case gumble.DisconnectError:
		reason = "connection error"
	}
	if reason == "" {
		fmt.Printf("Disconnected")
		b.End()
	} else {
		fmt.Printf("Disconnected: " + reason)
		b.End()
	}
}

func (b *Barnard) OnTextMessage(e *gumble.TextMessageEvent) {
}

func (b *Barnard) OnUserChange(e *gumble.UserChangeEvent) {
	if e.Type.Has(gumble.UserChangeChannel) && e.User == b.Client.Self {
		fmt.Printf("To: %s", e.User.Channel.Name)
	}
}

func (b *Barnard) OnChannelChange(e *gumble.ChannelChangeEvent) {
}

func (b *Barnard) OnPermissionDenied(e *gumble.PermissionDeniedEvent) {
	var info string
	switch e.Type {
	case gumble.PermissionDeniedOther:
		info = e.String
	case gumble.PermissionDeniedPermission:
		info = "insufficient permissions"
	case gumble.PermissionDeniedSuperUser:
		info = "cannot modify SuperUser"
	case gumble.PermissionDeniedInvalidChannelName:
		info = "invalid channel name"
	case gumble.PermissionDeniedTextTooLong:
		info = "text too long"
	case gumble.PermissionDeniedTemporaryChannel:
		info = "temporary channel"
	case gumble.PermissionDeniedMissingCertificate:
		info = "missing certificate"
	case gumble.PermissionDeniedInvalidUserName:
		info = "invalid user name"
	case gumble.PermissionDeniedChannelFull:
		info = "channel full"
	case gumble.PermissionDeniedNestingLimit:
		info = "nesting limit"
	}
	fmt.Printf("Permission denied: %s", info)
}

func (b *Barnard) VoiceToggle() {
//	if b.Pushed == false {
		//b.Client.Self.SetSelfMuted(true) //
		b.Stream.StopSource()
		fmt.Println("Idle")
	//	b.LastState = false
//	} else {
		//b.Client.Self.SetSelfMuted(false)
		b.Stream.StartSource()
		fmt.Println("Tx")
	//	b.LastState = true
//	}
}
func (b *Barnard) OnUserList(e *gumble.UserListEvent) {
}

func (b *Barnard) OnACL(e *gumble.ACLEvent) {
}

func (b *Barnard) OnBanList(e *gumble.BanListEvent) {
}

func (b *Barnard) OnContextActionChange(e *gumble.ContextActionChangeEvent) {
}

func (b *Barnard) OnServerConfig(e *gumble.ServerConfigEvent) {
}
