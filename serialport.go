package serialport

import (
	"bufio"
	"log"

	"github.com/tarm/serial"
)

func Open(portName string, baudRate int, dataBits byte, stopBits byte) (ISerialPort, error) {
	port, err := serial.OpenPort(&serial.Config{
		Name:     portName,
		Baud:     baudRate,
		Size:     dataBits,
		StopBits: serial.StopBits(stopBits),
	})

	if err != nil {
		return nil, err
	}

	return &serialPort{port: port}, nil
}

type ISerialPort interface {
	Send(data []byte) error
	Received(delim byte, callback func(data []byte))
	Close() error
}

type serialPort struct {
	port *serial.Port
}

func (s *serialPort) Send(data []byte) error {
	_, err := s.port.Write(data)

	return err
}

func (s *serialPort) Received(delim byte, callback func(data []byte)) {
	go func() {
		reader := bufio.NewReader(s.port)
		for {
			data, err := reader.ReadBytes(delim)
			if err != nil {
				log.Fatalln("Error reading from port: ", err)
				continue
			}

			if callback != nil {
				callback(data)
			}
		}
	}()
}

func (s *serialPort) Close() error {
	return s.port.Close()
}
