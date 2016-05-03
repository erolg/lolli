# lolli

barnard is a terminal-based client for the [Mumble](http://mumble.info) voice
chat software.

lolli is a terminal based and menageable (via REST) client. Only support channels operation for now

## Requirements

- [gumble](https://github.com/layeh/gumble/tree/master/gumble)
- libopenal-dev, needed to install via package manager (apt-get)
- libopus-dev, needed to install via package manager for gopus (apt-get)
- [gumbleopenal](https://github.com/layeh/gumble/tree/master/gumbleopenal)
- [termbox-go](https://github.com/nsf/termbox-go)
- [negroni](github.com/codegangsta/negroni)

to install package ```go get [github link (github.com/user/repo)]```

to update the package you can use ```-u``` parameter


##before-using
run these commands :

* ```openssl genrsa -out server.key 1024```
* ```openssl req -new -x509 -key server.key -out server.pem ```

##running
1. ```git clone ...```
2. ```cd lolli```
3. ```openssl genrsa -out server.key 1024```
4. ```openssl req -new -x509 -key server.key -out server.pem ```
5.  ```go run cmd/barnard/main.go```


## License

MPL 2.0

## Author

Tim Cooper (<tim.cooper@layeh.com>)
Erol Guzoğlu
