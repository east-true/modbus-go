// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.
package modbus_test

import (
	"testing"
	"time"

	"github.com/east-true/goutil/parser"
	. "github.com/east-true/modbus-go"
	"github.com/east-true/modbus-go/client"
	"github.com/east-true/modbus-go/memory"
)

func TestTCP(t *testing.T) {
	mem := memory.New(memory.FUNC_READ_HOLDING_REGISTERS, 0, parser.LITTLE_LOWER, parser.INT16ARR, 1)
	mb := NewTCP(&client.TCP{
		Address: "192.168.1.122:502",
		SlaveID: 1,
		Timeout: 60 * time.Second,
	}, mem)
	if err := mb.Connect(); err != nil {
		t.Error(err)
		return
	} else {
		defer mb.Close()
	}

	chunk := mb.Read()
	for i := range chunk {
		t.Logf("%+v", chunk[i])
	}
}
