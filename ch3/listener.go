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
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go func(c net.Conn) {
			defer c.Close()
		}(conn)
	}
}

/*
the net.Listen function accepts a network type, a ip address and a port separed by a collon.

it returns a net.Listener interface and a error interface, if is successfully the listener is bound to the ip.

Binding means that the operating system has exclusively assigned the port on the given IP address to the listener.

The operating system allows no other processes to listen for incoming traffic on bound ports.

You should always be diligent about closing your listener gracefully by calling its Close method.


Failure to close the listener may lead to memory leaks or deadlocks in your code, because calls to the listener’s Accept method may block
indefinitely. Closing the listener immediately unblocks calls to the Accept method.
*/
