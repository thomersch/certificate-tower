# Certificate Tower

Certificate Tower checks if your TLS certificates have expired and prints out warnings/errors. Use with [Bitbar](https://getbitbar.com).

## Instructions

**Prerequisites:** Go 1.8 or newer

Get and build with `go get -u github.com/thomersch/certificate-tower`. The binary will be in `$GOPATH/bin/certificate-tower`. If you haven't specified a `GOPATH` before, it will default to `$HOME/go`.

Place a text file in ~/.certificate-tower-hosts with one host name per line, e.g.

	example.com
	foobar.example.com

The path to the hosts file can be overridden by setting the `CERTTOWERHOSTS` environment variable.
