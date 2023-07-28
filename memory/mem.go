package memory

import (
	"strings"

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

type CallFunc interface {
	Read(c modbus.Client) ([]byte, error)
	Write(c modbus.Client) ([]byte, error)
}

type Mem struct {
	address  uint16
	quantity uint16
	read     func(c modbus.Client) ([]byte, error)
}

func New(address uint16, dataType string, cnt uint16) *Mem {
	var quantity uint16
	if strings.Contains(dataType, "32") {
		quantity = 2 * cnt
	} else if strings.Contains(dataType, "64") {
		quantity = 4 * cnt
	} else {
		quantity = 1 * cnt
	}

	if quantity < 1 && quantity > 125 {
		return nil
	}
	return &Mem{
		address:  address,
		quantity: uint16(quantity),
	}
}

func (mem *Mem) Read(c modbus.Client) ([]byte, error) {
	return mem.read(c)
}

func (mem *Mem) Write(c modbus.Client) ([]byte, error) {
	// TODO
	return nil, nil
}

func (mem *Mem) readCoils(c modbus.Client) ([]byte, error) {
	return c.ReadCoils(mem.address, mem.quantity)
}

func (mem *Mem) readDiscreteInputs(c modbus.Client) ([]byte, error) {
	return c.ReadDiscreteInputs(mem.address, mem.quantity)
}

func (mem *Mem) readHoldingRegisters(c modbus.Client) ([]byte, error) {
	return c.ReadHoldingRegisters(mem.address, mem.quantity)
}

func (mem *Mem) readInputRegisters(c modbus.Client) ([]byte, error) {
	return c.ReadInputRegisters(mem.address, mem.quantity)
}

// TODO : 파라미터 통합.
// type WriteReuqest struct {

// }
func (mem *Mem) writeCoil(c modbus.Client) ([]byte, error) {
	return c.WriteSingleCoil(mem.address, 0)
}

func (mem *Mem) writeCoils(c modbus.Client) ([]byte, error) {
	return c.WriteMultipleCoils(mem.address, mem.quantity, []byte{byte(0), byte(0)})
}

func (mem *Mem) writeRegister(c modbus.Client) ([]byte, error) {
	return c.WriteSingleRegister(mem.address, 0)
}

func (mem *Mem) writeRegisters(c modbus.Client) ([]byte, error) {
	return c.WriteMultipleRegisters(mem.address, mem.quantity, []byte{byte(0), byte(0)})
}
