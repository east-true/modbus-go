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
package memory

import (
	"errors"
	"time"

	"github.com/east-true/goutil/parser"
	"github.com/goburrow/modbus"
)

const (
	_ uint8 = iota
	FUNC_READ_COILS
	FUNC_READ_DISCRETE_INPUTS
	FUNC_READ_HOLDING_REGISTERS
	FUNC_READ_INPUT_REGISTERS
	FUNC_WRITE_COIL
	FUNC_WRITE_COILS
	FUNC_WRITE_REGISTER
	FUNC_WRITE_REGISTERS
)

type MemDelegate interface {
	Read(c modbus.Client) (*MemReadData, error)
	Write(c modbus.Client, value uint16) ([]byte, error)
}

type MemReadData struct {
	Addr  uint16
	Time  time.Time
	Value []interface{}
}

type Mem struct {
	address  uint16
	quantity uint16
	dataType uint8
	p        parser.Parser
	read     func(c modbus.Client) ([]interface{}, error)
	write    func(c modbus.Client, value uint16) ([]byte, error)
}

// dataType : parser.XXXXARR
func New(funcCode uint8, address uint16, order string, dataType uint8, cnt uint16) MemDelegate {
	mem := &Mem{
		address:  address,
		dataType: dataType,
		p:        parser.New(order),
	}

	switch dataType {
	case parser.INT32ARR, parser.UINT32ARR, parser.FLOAT32ARR:
		mem.quantity = 2 * cnt
	case parser.INT64ARR, parser.UINT64ARR, parser.FLOAT64ARR:
		mem.quantity = 4 * cnt
	default:
		mem.quantity = 1 * cnt
	}

	switch funcCode {
	case FUNC_READ_COILS:
		mem.read = mem.readCoils
	case FUNC_READ_DISCRETE_INPUTS:
		mem.read = mem.readDiscreteInputs
	case FUNC_READ_HOLDING_REGISTERS:
		mem.read = mem.readHoldingRegisters
	case FUNC_READ_INPUT_REGISTERS:
		mem.read = mem.readInputRegisters
	case FUNC_WRITE_COIL:
		mem.write = mem.writeCoil
	case FUNC_WRITE_REGISTER:
		mem.write = mem.writeRegister
	}

	return mem
}

func (mem *Mem) Read(c modbus.Client) (*MemReadData, error) {
	value, err := mem.read(c)
	return &MemReadData{
		Addr:  mem.address,
		Time:  time.Now(),
		Value: value,
	}, err
}

func (mem *Mem) Write(c modbus.Client, value uint16) ([]byte, error) {
	return mem.write(c, value)
}

func (mem *Mem) readCoils(c modbus.Client) ([]interface{}, error) {
	if b, err := c.ReadCoils(mem.address, mem.quantity); err != nil {
		return nil, err
	} else {
		return mem.ToBit(b), nil
	}
}

func (mem *Mem) readDiscreteInputs(c modbus.Client) ([]interface{}, error) {
	if b, err := c.ReadDiscreteInputs(mem.address, mem.quantity); err != nil {
		return nil, err
	} else {
		return mem.ToBit(b), nil
	}
}

func (mem *Mem) readHoldingRegisters(c modbus.Client) ([]interface{}, error) {
	if b, err := c.ReadHoldingRegisters(mem.address, mem.quantity); err != nil {
		return nil, err
	} else {
		return mem.ToData(b)
	}
}

func (mem *Mem) readInputRegisters(c modbus.Client) ([]interface{}, error) {
	if b, err := c.ReadInputRegisters(mem.address, mem.quantity); err != nil {
		return nil, err
	} else {
		return mem.ToData(b)
	}
}

func (mem *Mem) ToBit(b []byte) []interface{} {
	bits := mem.p.ToBitArr(b)
	wrapper := make([]interface{}, len(bits))
	for i := range wrapper {
		wrapper[i] = bits[i]
	}

	return wrapper
}

func (mem *Mem) ToData(b []byte) ([]interface{}, error) {
	switch mem.dataType {
	case parser.INT16ARR:
		data, err := mem.p.ToInt16Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case parser.UINT16ARR:
		data, err := mem.p.ToUint16Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case parser.INT32ARR:
		data, err := mem.p.ToInt32Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case parser.UINT32ARR:
		data, err := mem.p.ToUint32Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case parser.INT64ARR:
		data, err := mem.p.ToInt64Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case parser.UINT64ARR:
		data, err := mem.p.ToUint64Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case parser.FLOAT32ARR:
		data, err := mem.p.ToFloat32Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case parser.FLOAT64ARR:
		data, err := mem.p.ToFloat64Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	default:
		return nil, errors.New("Not Supported Data Type")
	}
}

func (mem *Mem) writeCoil(c modbus.Client, value uint16) ([]byte, error) {
	return c.WriteSingleCoil(mem.address, value)
}

func (mem *Mem) writeRegister(c modbus.Client, value uint16) ([]byte, error) {
	return c.WriteSingleRegister(mem.address, value)
}

// TODO
// func (mem *Mem) writeCoils(c modbus.Client) ([]byte, error) {
// 	return c.WriteMultipleCoils(mem.address, mem.quantity, []byte{byte(0), byte(0)})
// }

// func (mem *Mem) writeRegisters(c modbus.Client) ([]byte, error) {
// 	return c.WriteMultipleRegisters(mem.address, mem.quantity, []byte{byte(0), byte(0)})
// }
