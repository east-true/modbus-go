// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.
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
