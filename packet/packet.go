package packet

import (
	"log"
	"net"

	"github.com/song940/tuntap-go/tuntap"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

type Packet []byte

func (p *Packet) Resize(length int) {
	if cap(*p) < length {
		old := *p
		*p = make(Packet, length)
		copy(*p, old)
	} else {
		*p = (*p)[:length]
	}
}

func (p Packet) Read(ifce *tuntap.Interface) (err error) {
	n, err := ifce.Read(p)
	if err != nil {
		return
	}
	p.Resize(n)
	log.Printf("packet: % x\n", p[:n])
	return
}

func (p Packet) Version() byte {
	return p[0] >> 4
}

type Header struct {
	Version uint8
	Src     net.IP
	Dst     net.IP
}

func (p Packet) ParseHeader() (header Header, err error) {
	header.Version = p.Version()
	switch header.Version {
	case 4:
		hdr, err := ipv4.ParseHeader(p)
		if err != nil {
			return header, err
		}
		header.Src = hdr.Src
		header.Dst = hdr.Dst
	case 6:
		hdr, err := ipv6.ParseHeader(p)
		if err != nil {
			return header, err
		}
		header.Src = hdr.Src
		header.Dst = hdr.Dst
	}
	return
}
