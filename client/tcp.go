package handler

import (
	"time"

	"github.com/goburrow/modbus"
)

type TCP struct {
	Address string
	SlaveID byte
	Timeout time.Duration
	h       *modbus.TCPClientHandler
	c       modbus.Client
}

func NewTCP() *TCP {
	return &TCP{
		Address: "127.0.0.1:113",
		SlaveID: 1,
		Timeout: 5,
	}
}

func (tcp *TCP) SetHandler() {
	tcp.h = modbus.NewTCPClientHandler(tcp.Address)
	tcp.h.SlaveId = tcp.SlaveID
	tcp.h.Timeout = tcp.Timeout
}

func (tcp *TCP) Connect() error {
	if err := tcp.h.Connect(); err != nil {
		return err
	}

	tcp.c = modbus.NewClient(tcp.h)
	return nil
}

func (tcp *TCP) Close() error {
	return tcp.h.Close()
}

func (tcp *TCP) Client() modbus.Client {
	return tcp.c
}
