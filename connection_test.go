package warden

import (
	"bufio"
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"errors"
	"fmt"
	"net"
	"testing"
	"time"
	. "launchpad.net/gocheck"
)

type WSuite struct {}

func Test(t *testing.T) { TestingT(t) }

func init() {
  Suite(&WSuite{})
}

func (w *WSuite) TestConnectionCreating(c *C) {
  conn := &fakeConn{
    ReadBuffer: messages(&CreateResponse{
      Handle: proto.String("foohandle"),
    }),

    WriteBuffer: bytes.NewBuffer([]byte{}),
  }

  connection := connectionWith(conn)

  resp, err := connection.Create()
  c.Assert(err, IsNil)

  c.Assert(
    string(conn.WriteBuffer.Bytes()),
    Equals,
    string(messages(&CreateRequest{}).Bytes()),
  )

  c.Assert(resp.GetHandle(), Equals, "foohandle")
}

func (w *WSuite) TestConnectionDestroying(c *C) {
  conn := &fakeConn{
    ReadBuffer: messages(&DestroyResponse{}),
    WriteBuffer: bytes.NewBuffer([]byte{}),
  }

  connection := connectionWith(conn)

  _, err := connection.Destroy("foo")
  c.Assert(err, IsNil)

  c.Assert(
    string(conn.WriteBuffer.Bytes()),
    Equals,
    string(messages(&DestroyRequest{Handle: proto.String("foo")}).Bytes()),
  )
}

func (w *WSuite) TestConnectionSpawn(c *C) {
  conn := &fakeConn{
    ReadBuffer: messages(&SpawnResponse{JobId: proto.Uint32(42)}),
    WriteBuffer: bytes.NewBuffer([]byte{}),
  }

  connection := connectionWith(conn)

  resp, err := connection.Spawn("foo-handle", "echo hi")
  c.Assert(err, IsNil)

  c.Assert(
    string(conn.WriteBuffer.Bytes()),
    Equals,
    string(messages(&SpawnRequest{
      Handle: proto.String("foo-handle"),
      Script: proto.String("echo hi"),
    }).Bytes()),
  )

  c.Assert(resp.GetJobId(), Equals, uint32(42))
}

func (w *WSuite) TestConnectionRun(c *C) {
  conn := &fakeConn{
    ReadBuffer: messages(&RunResponse{ExitStatus: proto.Uint32(137)}),
    WriteBuffer: bytes.NewBuffer([]byte{}),
  }

  connection := connectionWith(conn)

  resp, err := connection.Run("foo-handle", "echo hi")
  c.Assert(err, IsNil)

  c.Assert(
    string(conn.WriteBuffer.Bytes()),
    Equals,
    string(messages(&RunRequest{
      Handle: proto.String("foo-handle"),
      Script: proto.String("echo hi"),
    }).Bytes()),
  )

  c.Assert(resp.GetExitStatus(), Equals, uint32(137))
}

func (w *WSuite) TestConnectionStream(c *C) {
  conn := &fakeConn{
    ReadBuffer: messages(
      &StreamResponse{Name: proto.String("stdout"), Data: proto.String("1")},
      &StreamResponse{Name: proto.String("stderr"), Data: proto.String("2")},
      &StreamResponse{ExitStatus: proto.Uint32(3)},
    ),
    WriteBuffer: bytes.NewBuffer([]byte{}),
  }

  connection := connectionWith(conn)

  resp, err := connection.Stream("foo-handle", 42)
  c.Assert(err, IsNil)

  c.Assert(
    string(conn.WriteBuffer.Bytes()),
    Equals,
    string(messages(&StreamRequest{
      Handle: proto.String("foo-handle"),
      JobId: proto.Uint32(42),
    }).Bytes()),
  )

  res1 := <-resp
  c.Assert(res1.GetName(), Equals, "stdout")
  c.Assert(res1.GetData(), Equals, "1")

  res2 := <-resp
  c.Assert(res2.GetName(), Equals, "stderr")
  c.Assert(res2.GetData(), Equals, "2")

  res3, ok := <-resp
  c.Assert(res3.GetExitStatus(), Equals, uint32(3))
  c.Assert(ok, Equals, true)
}

func (w *WSuite) TestConnectionError(c *C) {
  conn := &fakeConn{
    ReadBuffer: messages(&ErrorResponse{Message: proto.String("boo")}),
    WriteBuffer: bytes.NewBuffer([]byte{}),
  }

  connection := connectionWith(conn)

  resp, err := connection.Run("foo-handle", "echo hi")
  c.Assert(resp, IsNil)
  c.Assert(err, Not(IsNil))

  c.Assert(err.Error(), Equals, "boo")
}

func connectionWith(conn net.Conn) *Connection {
  return &Connection{
    conn: conn,
    read: bufio.NewReader(conn),
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
      Type: Message_Type(message2type(msg)).Enum(),
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
