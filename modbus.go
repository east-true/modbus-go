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
package modbus

import (
	"fmt"

	"github.com/east-true/modbus-go/client"
	"github.com/east-true/modbus-go/memory"
)

type Modbus struct {
	c   client.Client
	fns []memory.MemDelegate
}

func NewRTU(rtu *client.RTU, mems ...memory.MemDelegate) *Modbus {
	if rtu == nil {
		rtu = client.NewRTU()
	}
	rtu.SetHandler()
	return &Modbus{
		c:   rtu,
		fns: mems,
	}
}

func NewTCP(tcp *client.TCP, mems ...memory.MemDelegate) *Modbus {
	if tcp == nil {
		tcp = client.NewTCP()
	}
	tcp.SetHandler()
	return &Modbus{
		c:   tcp,
		fns: mems,
	}
}

func (mb *Modbus) Connect() error {
	return mb.c.Connect()
}

func (mb *Modbus) Close() error {
	return mb.c.Close()
}

func (mb *Modbus) Read() []*memory.MemReadData {
	chunk := make([]*memory.MemReadData, len(mb.fns))
	for i := range mb.fns {
		if data, err := mb.fns[i].Read(mb.c.GetClient()); err != nil {
			fmt.Println(err)
			continue
		} else {
			chunk[i] = data
		}
	}

	return chunk
}

func (mb *Modbus) Write(value ...uint16) {
	for i := range mb.fns {
		if res, err := mb.fns[i].Write(mb.c.GetClient(), value[i]); err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Println(string(res))
		}
	}
}
