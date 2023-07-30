package modbus_test

import (
	"testing"

	"github.com/east-true/goutil/parser"
	. "github.com/east-true/modbus-go"
	"github.com/east-true/modbus-go/memory"
)

func TestTCP(t *testing.T) {
	mem := memory.New(memory.FUNC_READ_HOLDING_REGISTERS, 0, parser.LITTLE_LOWER, parser.INT32ARR, 1)
	tcp := NewTCP(nil, mem)
	if err := tcp.Connect(); err != nil {
		t.Error(err)
	} else {
		defer tcp.Close()
	}

	tcp.Read()
}
