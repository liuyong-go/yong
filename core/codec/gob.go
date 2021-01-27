package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

//GobCodec 结构体
type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

var _ Codec = (*GobCodec)(nil)

//NewGobCodec 实例化bob编码
func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{ //后面补充实现接口的方法
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

//ReadHeader 实现接口ReadHeader方法
func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

//ReadBody 实现接口ReadBody方法
func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

//Write 实现接口Write方法
func (c *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()
	if err := c.enc.Encode(h); err != nil {
		log.Println("rpc codec: gob error encoding header:", err)
	}
	if err := c.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob error encoding body:", err)
		return err
	}
	return nil
}

//Close 实现接口close方法
func (c *GobCodec) Close() error {
	return c.conn.Close()
}
