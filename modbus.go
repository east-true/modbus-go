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

func NewRTU(mems ...memory.MemDelegate) *Modbus {
	client := client.NewTCP()
	client.SetHandler()
	return &Modbus{
		c:   client,
		fns: mems,
	}
}

func NewTCP(mems ...memory.MemDelegate) *Modbus {
	client := client.NewTCP()
	client.SetHandler()
	return &Modbus{
		c:   client,
		fns: mems,
	}
}

func (mb *Modbus) Connect() error {
	return mb.c.Connect()
}

func (mb *Modbus) Close() error {
	return mb.c.Close()
}

func (mb *Modbus) Read() {
	for i := range mb.fns {
		if data, err := mb.fns[i].Read(mb.c.GetClient()); err != nil {
			fmt.Println(err)
			continue
		} else {
			for i := range data {
				// TODO
				fmt.Println(data[i])
			}
		}
	}
}

func (mb *Modbus) Write() {
	for i := range mb.fns {
		if res, err := mb.fns[i].Write(mb.c.GetClient()); err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Println(string(res))
		}
	}
}
