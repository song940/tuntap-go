package cli

import (
	"log"

	"github.com/song940/tuntap-go/tuntap"
)

func Run() {
	config := tuntap.Config{
		DeviceType: tuntap.TUN,
	}
	ifce, err := tuntap.New(config)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Interface Name: %s\n", ifce.Name())
	buf := make([]byte, 1500)
	for {
		n, err := ifce.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Received packet from TUN device: %v", buf[:n])
	}
}
