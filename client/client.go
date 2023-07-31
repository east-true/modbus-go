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
		Address: "127.0.0.1:502",
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
