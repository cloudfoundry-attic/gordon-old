package warden

import (
	"bufio"
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"errors"
	"fmt"
	. "launchpad.net/gocheck"
	"net"
	"testing"
	"time"
)

type WSuite struct{}

func Test(t *testing.T) { TestingT(t) }

func init() {
	Suite(&WSuite{})
}

func connectionWith(conn net.Conn) *Connection {
	return &Connection{
		conn: conn,
		read: bufio.NewReader(conn),

		// buffer size of 1 so that read and write errors
		// can both send without blocking
		disconnected: make(chan bool, 1),
	}
}

func messages(msgs ...proto.Message) *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{})

	for _, msg := range msgs {
		payload, err := proto.Marshal(msg)
		if err != nil {
			panic(err.Error())
		}

		message := &Message{
			Type:    Message_Type(message2type(msg)).Enum(),
			Payload: payload,
		}

		messagePayload, err := proto.Marshal(message)
		if err != nil {
			panic("failed to marshal message")
		}

		buf.Write([]byte(fmt.Sprintf("%d\r\n%s\r\n", len(messagePayload), messagePayload)))
	}

	return buf
}

type fakeConn struct {
	ReadBuffer  *bytes.Buffer
	WriteBuffer *bytes.Buffer
	WriteChan   chan string
	Closed      bool
}

func (f *fakeConn) Read(b []byte) (n int, err error) {
	if f.Closed {
		return 0, errors.New("buffer closed")
	}

	return f.ReadBuffer.Read(b)
}

func (f *fakeConn) Write(b []byte) (n int, err error) {
	if f.Closed {
		return 0, errors.New("buffer closed")
	}

	if f.WriteChan != nil {
		f.WriteChan <- string(b)
	}

	return f.WriteBuffer.Write(b)
}

func (f *fakeConn) Close() error {
	f.Closed = true
	return nil
}

func (f *fakeConn) SetDeadline(time.Time) error {
	return nil
}

func (f *fakeConn) SetReadDeadline(time.Time) error {
	return nil
}

func (f *fakeConn) SetWriteDeadline(time.Time) error {
	return nil
}

func (f *fakeConn) LocalAddr() net.Addr {
	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:4222")
	return addr
}

func (f *fakeConn) RemoteAddr() net.Addr {
	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:65525")
	return addr
}
