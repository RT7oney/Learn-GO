package main

import (
    "net/http"
    "github.com/gorilla/websocket"
)

http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
    conn, _ := websocket.Upgrade(r, w)
    ch := NewChannel(conn)
    //...
})

// Packet represents application level data.
type Packet struct {
	// ...
}

// Channel wraps user connection.
type Channel struct {
	conn net.Conn    // WebSocket connection.
	send chan Packet // Outgoing packets queue.
}

func NewChannel(conn net.Conn) *Channel {
	c := &Channel{
		conn: conn,
		send: make(chan Packet, N),
	}

	go c.reader()
	go c.writer()

	return c
}

func (c *Channel) reader() {
	// We make a buffered read to reduce read syscalls.
	buf := bufio.NewReader(c.conn)

	for {
		pkt, _ := readPacket(buf)
		c.handle(pkt)
	}
}


func (c *Channel) writer() {
    // We make buffered write to reduce write syscalls.
    buf := bufio.NewWriter(c.conn)

    for pkt := range c.send {
        _ := writePacket(buf, pkt)
        buf.Flush()
    }
}