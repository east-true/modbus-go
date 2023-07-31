/*
BSD 3-Clause License

# Copyright (c) 2023, DongJinLee

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

 1. Redistributions of source code must retain the above copyright notice, this
    list of conditions and the following disclaimer.

 2. Redistributions in binary form must reproduce the above copyright notice,
    this list of conditions and the following disclaimer in the documentation
    and/or other materials provided with the distribution.

 3. Neither the name of the copyright holder nor the names of its
    contributors may be used to endorse or promote products derived from
    this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
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
