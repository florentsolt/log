package log

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

type TcpInterceptor struct {
	connections map[uint64]net.Conn
	mux         sync.Mutex
	backlog     [][]byte
	channel     chan []byte
	inc         uint64
}

const BacklogMaxLen = 100

func (ti *TcpInterceptor) ListenAndAccept(address string) error {
	l, err := net.Listen("tcp4", address)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer l.Close()
	ti.connections = map[uint64]net.Conn{}
	ti.backlog = [][]byte{}
	ti.channel = make(chan []byte)
	go ti.waitForLog()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return err
		}
		if ti.writeBacklog(c) { // if backlog was sent, add it to the pool of connection
			ti.mux.Lock()
			ti.inc++
			ti.connections[ti.inc] = c
			ti.mux.Unlock()
		}
	}
}

func (ti *TcpInterceptor) waitForLog() {
	for p := range ti.channel {
		ti.mux.Lock()
		if len(ti.backlog) == BacklogMaxLen {
			ti.backlog = append(ti.backlog[1:], p)
		} else {
			ti.backlog = append(ti.backlog, p)
		}
		ti.mux.Unlock()
		for k, c := range ti.connections {
			_, err := c.Write(p)
			if err != nil {
				c.Close()
				ti.mux.Lock()
				delete(ti.connections, k)
				ti.mux.Unlock()
			}
		}

	}
}

func (ti *TcpInterceptor) writeBacklog(c net.Conn) bool {
	buf := new(bytes.Buffer)
	ti.mux.Lock()
	for _, p := range ti.backlog {
		buf.Write(p)
		buf.WriteString("\n")
	}
	ti.mux.Unlock()
	_, err := buf.WriteTo(c)
	if err != nil {
		c.Close()
		return false
	}
	return true
}

func (ti *TcpInterceptor) Write(p []byte) (n int, err error) {
	cpy := make([]byte, len(p)+1)
	copy(cpy, p)
	cpy[len(p)] = byte('\n')
	ti.channel <- cpy
	return len(p), nil
}
