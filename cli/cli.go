package cli

import (
	"log"

	"github.com/song940/tuntap-go/packet"
	"github.com/song940/tuntap-go/tuntap"
)

func Run() {
	config := tuntap.Config{
		DeviceType: tuntap.TUN,
	}
	config.Name = "utun9"
	ifce, err := tuntap.New(config)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Interface Name: %s\n", ifce.Name())
	var packet packet.Packet
	packet.Resize(1500)
	for {
		n, err := ifce.Read(packet)
		if err != nil {
			log.Fatal(err)
		}
		packet.Resize(n)
		header, _ := packet.ParseHeader()
		log.Println(header.Version)
		log.Println(header.Src)
		log.Println(header.Dst)
	}
}
