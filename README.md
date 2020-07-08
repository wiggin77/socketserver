# socketserver

Simple implementation of a socket server that accepts TCP connections, reads data, and outputs the stream to stdout. This is used for testing TCP services.

## build

Tested with Go 1.14.2, however any recent Go version should work fine.  Clone this repository then run

```bash
go build
```

## usage

```bash
❯ ./socketserver

Usage:
	socketserver <port number>
	- <port number> is the listening port
```

