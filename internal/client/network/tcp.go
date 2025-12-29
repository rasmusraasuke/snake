package network

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

type NetworkStatus int

const (
	Offline = iota
	Online
)

type TCPClient struct {
	Addr     string
	Status   NetworkStatus
	Conn     net.Conn
	Send     chan []byte
	Incoming chan []byte
	Done     chan struct{}
}

func NewTCPClient(addr string) *TCPClient {
	return &TCPClient{
		Addr:     addr,
		Status:   Offline,
		Send:     make(chan []byte, 1024),
		Incoming: make(chan []byte, 1024),
		Done:     make(chan struct{}),
	}
}

func (c *TCPClient) Connect() {
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		log.Println(err)
		return
	}
	c.Conn = conn
	c.Status = Online

	go c.readLoop()
	go c.writeLoop()
}

func (c *TCPClient) SendMessage(msg []byte) {
	buf := make([]byte, 2+len(msg))
	binary.BigEndian.PutUint16(buf[:2], uint16(len(msg)))
	copy(buf[2:], msg)
	c.Send <- buf
}

func (c *TCPClient) readLoop() {
	defer close(c.Incoming)
	for {
		var length uint16
		if err := binary.Read(c.Conn, binary.BigEndian, &length); err != nil {
			if err != io.EOF {
				log.Println("Read length error:", err)
			}
			close(c.Done)
			return
		}

		buf := make([]byte, length)
		if _, err := io.ReadFull(c.Conn, buf); err != nil {
			log.Println("Read payload error:", err)
			close(c.Done)
			return
		}

		c.Incoming <- buf
	}
}

func (c *TCPClient) writeLoop() {
	for msg := range c.Send {
		if _, err := c.Conn.Write(msg); err != nil {
			log.Println("Write error:", err)
			close(c.Done)
			return
		}
	}
}

func (c *TCPClient) Close() {
	c.Conn.Close()
	close(c.Send)
	c.Status = Offline
}
