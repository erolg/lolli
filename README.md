# lolli

barnard is a terminal-based client for the [Mumble](http://mumble.info) voice
chat software.

lolli is a terminal based and menageable (via REST) client. Only support channels operation for now

## Requirements

- [gumble](https://github.com/layeh/gumble/tree/master/gumble)
- [gumbleopenal](https://github.com/layeh/gumble/tree/master/gumbleopenal)
- [termbox-go](https://github.com/nsf/termbox-go)
- [negroni](github.com/codegangsta/negroni)
- libopenal-dev, needed to instal via package manager

##before-using
```openssl genrsa -out server.key 1024
openssl req -new -x509 -key server.key -out server.pem ```

## License

MPL 2.0

## Author

Tim Cooper (<tim.cooper@layeh.com>)
Erol Guzoğlu
