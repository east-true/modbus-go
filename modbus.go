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

func NewRTU(byteOrder string, dels []memory.MemDelegate) *Modbus {
	client := client.NewTCP()
	client.SetHandler()
	return &Modbus{
		c:   client,
		fns: dels,
	}
}

func NewTCP(byteOrder string, dels []memory.MemDelegate) *Modbus {
	client := client.NewTCP()
	client.SetHandler()
	return &Modbus{
		c:   client,
		fns: dels,
	}
}

// 리턴 되는 데이터를 바이트가 아닌 변환한 interface로 받아야 할 듯
// 여기서는 어떤 function이 호출되었는지 알 수 없음................................
// ..............................하
func (mb *Modbus) Read() {
	for i := range mb.fns {
		if data, err := mb.fns[i].Read(mb.c.GetClient()); err != nil {
			fmt.Println(err)
			continue
		} else {
			// TODO
			for i := range data {
				fmt.Println(data[i])
			}
		}
	}
}

// TODO : functions remove
// // quantity : coil cnt
// func (mem *Mem) readCoils(c modbus.Client) ([]interface{}, error) {
// 	b, err := c.ReadCoils(mem.Addr, mem.quantity)
// 	if err != nil {
// 		return nil, err
// 	}

// 	bits := mem.parser.ToBit(b)
// 	if len(bits) < 1 {
// 		return nil, errors.New("data is empty...!")
// 	}

// 	return bits[:len(mem.Tags)], nil
// }

// // quantity : input cnt
// func (mem *Mem) readDiscreteInputs(c modbus.Client) ([]interface{}, error) {
// 	b, err := c.ReadDiscreteInputs(mem.Addr, mem.quantity)
// 	if err != nil {
// 		return nil, err
// 	}

// 	bits := mem.parser.ToBit(b)
// 	if len(b) < 1 {
// 		return nil, errors.New("data is empty...!")
// 	}

// 	return bits[:len(mem.Tags)], nil
// }

// func (mem *Mem) readHoldingRegisters(c modbus.Client) ([]interface{}, error) {
// 	b, err := c.ReadHoldingRegisters(mem.Addr, mem.quantity)
// 	if err != nil {
// 		return nil, err
// 	}

// 	idata, err := mem.parser.ToData(b)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(idata) < 1 {
// 		return nil, errors.New("data is empty...!")
// 	}

// 	return idata, nil
// }

// func (mem *Mem) readInputRegisters(c modbus.Client) ([]interface{}, error) {
// 	b, err := c.ReadInputRegisters(mem.Addr, mem.quantity)
// 	if err != nil {
// 		return nil, err
// 	}

// 	idata, err := mem.parser.ToData(b)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(idata) < 1 {
// 		return nil, errors.New("data is empty...!")
// 	}

// 	return idata, nil
// }
