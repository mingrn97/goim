package transport_test

import (
	"bufio"
	"bytes"
	"testing"

	"itumate.com/im/transport"
)

func TestEncode(t *testing.T) {

	var message = &transport.Message{
		Type: transport.PlainType,
		Data: "hello, world",
	}

	_, err := transport.Encode(message)
	if err != nil {
		t.Errorf("消息编码失败: %s\n", err.Error())
	}
}

func TestDecode(t *testing.T) {

	var message = &transport.Message{
		Type: transport.PlainType,
		Data: "hello, world",
	}

	encode, err := transport.Encode(message)
	if err != nil {
		t.Errorf("消息编码失败: %s\n", err.Error())
	}

	r := bufio.NewReader(bytes.NewBuffer(encode))
	_, err = transport.Decode(r)
	if err != nil {
		t.Errorf("消息解码失败: %s\n", err.Error())
	}
}
