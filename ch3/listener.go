package ch03

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()

	t.Logf("bound to %q", listener.Addr())
}

// the net.Listen function accepts a network type, a ip address and a port separed by a collon.

// it returns a net.Listener interface and a error interface, if is successfully the listener is bound to the ip.

// Binding means that the operating system has exclusively assigned the port on the given IP address to the listener.
