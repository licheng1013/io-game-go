package message

type Message interface {
	GetMerge() int64
	GetBody() []byte
	GetHeartbeat() bool
	GetCode() int64
	GetMessage() string
	GetBytesResult() []byte
	SetBody([]byte)
}
