// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.
package client

import (
	"time"

	"github.com/goburrow/modbus"
)

type Client interface {
	GetClient() modbus.Client
	Connect() error
	Close() error
}

type RTU struct {
	SlaveID  byte
	Address  string
	Timeout  time.Duration // nano sec
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
		Timeout:  3 * time.Second,
		BaudRate: 9600,
		DataBits: 8,
		Parity:   "N",
		StopBits: 1,
	}
}

func (rtu *RTU) GetClient() modbus.Client {
	return rtu.c
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

type TCP struct {
	Address string
	SlaveID byte
	Timeout time.Duration // nano sec
	h       *modbus.TCPClientHandler
	c       modbus.Client
}

func NewTCP() *TCP {
	return &TCP{
		Address: "127.0.0.1:113",
		SlaveID: 1,
		Timeout: 3 * time.Second,
	}
}

func (tcp *TCP) GetClient() modbus.Client {
	return tcp.c
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
