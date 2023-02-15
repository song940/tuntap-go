package cli

import (
	"log"

	"github.com/song940/tuntap-go/ethernet"
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
	// buf := make([]byte, 1500)
	// for {
	// 	n, err := ifce.Read(buf)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("packet: % x\n", buf[:n])
	// }
	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		n, err := ifce.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}
		frame = frame[:n]
		log.Printf("Ethertype: % x\n", frame.Ethertype())
		log.Printf("Src: %s\n", frame.Source())
		log.Printf("Dst: %s\n", frame.Destination())
		log.Printf("Payload: % x\n", frame.Payload())
	}
}
