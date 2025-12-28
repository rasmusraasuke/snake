package network

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type Client struct {
	ID   string
	Conn net.Conn
	Send chan []byte
}

type TCPNetwork struct {
	Addr     string
	clients  map[string]*Client
	mu       sync.Mutex
	incoming chan tcpMessage
}

type tcpMessage struct {
	clientID string
	payload  []byte
}

func NewTCPNetwork(addr string) *TCPNetwork {
	return &TCPNetwork{
		Addr:     addr,
		clients:  make(map[string]*Client),
		incoming: make(chan tcpMessage, 1024),
	}
}

func (t *TCPNetwork) Start() error {
	listener, err := net.Listen("tcp", t.Addr)
	if err != nil {
		return err
	}
	log.Println("TCP server listening on", t.Addr)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("Accept error:", err)
				continue
			}

			clientID := conn.RemoteAddr().String()
			client := &Client{
				ID:   clientID,
				Conn: conn,
				Send: make(chan []byte, 1024),
			}

			t.mu.Lock()
			t.clients[clientID] = client
			t.mu.Unlock()

			go t.handleClientRead(client)
			go t.handleClientWrite(client)

			log.Println("Client connected:", clientID)
		}
	}()

	return nil
}

func (t *TCPNetwork) handleClientRead(c *Client) {
	defer func() {
		c.Conn.Close()
		t.mu.Lock()
		delete(t.clients, c.ID)
		t.mu.Unlock()
		close(c.Send)
		log.Println("Client disconnected:", c.ID)
	}()

	for {
		var length uint16
		if err := binary.Read(c.Conn, binary.BigEndian, &length); err != nil {
			if err != io.EOF {
				log.Println("Read length error:", err)
			}
			return
		}

		buf := make([]byte, length)
		if _, err := io.ReadFull(c.Conn, buf); err != nil {
			log.Println("Read payload error:", err)
			return
		}

		t.incoming <- tcpMessage{clientID: c.ID, payload: buf}
	}
}

func (t *TCPNetwork) handleClientWrite(c *Client) {
	for msg := range c.Send {
		if _, err := c.Conn.Write(msg); err != nil {
			log.Println("Write error to", c.ID, err)
			return
		}
	}
}

func (t *TCPNetwork) Receive() (string, []byte, error) {
	msg := <-t.incoming
	return msg.clientID, msg.payload, nil
}

func (t *TCPNetwork) Send(clientID string, msg []byte) error {
	t.mu.Lock()
	client, ok := t.clients[clientID]
	t.mu.Unlock()
	if !ok {
		return fmt.Errorf("client not found: %s", clientID)
	}

	buf := make([]byte, 2+len(msg))
	binary.BigEndian.PutUint16(buf[:2], uint16(len(msg)))
	copy(buf[2:], msg)

	client.Send <- buf
	return nil
}

func (t *TCPNetwork) Broadcast(msg []byte) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	buf := make([]byte, 2+len(msg))
	binary.BigEndian.PutUint16(buf[:2], uint16(len(msg)))
	copy(buf[2:], msg)

	for _, client := range t.clients {
		client.Send <- buf
	}
	return nil
}
