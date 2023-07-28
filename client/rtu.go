package handler

import (
	"time"

	"github.com/goburrow/modbus"
)

type RTU struct {
	SlaveID  byte
	Address  string
	Timeout  time.Duration
	BaudRate int
	DataBits int
	Parity   string
	StopBits int
	h        *modbus.RTUClientHandler
	c        modbus.Client
}

func NewRTU() *RTU {
	return &RTU{
		SlaveID:  1,
		Address:  "/dev",
		Timeout:  5,
		BaudRate: 1,
		DataBits: 8,
		Parity:   "N",
		StopBits: 1,
	}
}

func (rtu *RTU) SetHandler() {
	rtu.h = modbus.NewRTUClientHandler(rtu.Address)
	rtu.h.BaudRate = rtu.BaudRate
	rtu.h.DataBits = rtu.DataBits
	rtu.h.Parity = rtu.Parity
	rtu.h.StopBits = rtu.StopBits
	rtu.h.SlaveId = rtu.SlaveID
	rtu.h.Timeout = rtu.Timeout
}

func (rtu *RTU) Connect() error {
	if err := rtu.h.Connect(); err != nil {
		return err
	}

	rtu.c = modbus.NewClient(rtu.h)
	return nil
}

func (rtu *RTU) Close() error {
	return rtu.h.Close()
}

func (rtu *RTU) Client() modbus.Client {
	return rtu.c
}
