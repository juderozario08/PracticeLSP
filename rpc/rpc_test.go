package rpc_test

import (
	"testing"

	"phri/lsp/rpc"
)

type EncodingExample struct {
	Testing bool
}

type DecodingExample struct {
	Method bool
}

func TestEncoding(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual, err := rpc.EncodeMessage(EncodingExample{Testing: true})
	if err != nil {
		t.Fatal(err.Error())
	} else if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecoding(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err.Error())
	} else if contentLength != 15 {
		t.Fatalf("Expected: %v, Actual: %v", 15, contentLength)
	} else if method != "hi" {
		t.Fatalf("Expected: hi, Actual: %v", method)
	}
}
