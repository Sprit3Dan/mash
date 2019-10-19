package transport

import (
	"github.com/golang/protobuf/proto"
	"mash.com/event"
	"net"
)

type UDPTransport struct {
	Port int
}

func (ut *UDPTransport) Listen(eventChan chan event.Event, errorChan chan error) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 1337,
		Zone: "",
	})

	if err != nil {
		errorChan <- err
		return
	}

	for {
		bArr := make([]byte, 1024)
		newEvt := &event.Event{}

		_, netErr := conn.Read(bArr)
		if netErr != nil {
			errorChan <- netErr
			continue
		}

		parseErr := proto.Unmarshal(bArr, newEvt)
		if parseErr != nil {
			errorChan <- parseErr
			continue
		}

		eventChan <- *newEvt
	}
}

func (ut *UDPTransport) Connected() bool {
	return false
}


