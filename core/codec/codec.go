package codec

import "io"

//Header 头部包含信息
type Header struct {
	ServiceMethod string
	Seq           uint64
	Error         string
}

//Codec 抽象接口
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

//NewCodecFunc 抽象codec 构造函数
type NewCodecFunc func(io.ReadWriteCloser) Codec

//Type 定义类型
type Type string

const (
	//GobType 类型
	GobType Type = "application/gob"
	//JSONType json类型
	JSONType Type = "application/json"
)

//NewCodecFuncMap codec结构体集合
var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
