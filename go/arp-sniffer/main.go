package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Nome da interface de rede (mude para "wlan0", "eth0", etc.)
	device := "eth0"

	// Abrir a interface para captura
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Criar o analisador de pacotes
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	fmt.Println("📡 Escutando pacotes ARP na rede... Pressione Ctrl+C para parar.\n")

	// Loop que recebe e trata pacotes
	for packet := range packetSource.Packets() {
		arpLayer := packet.Layer(layers.LayerTypeARP)
		if arpLayer != nil {
			arp := arpLayer.(*layers.ARP)

			fmt.Printf("🟢 %s: %s (%s) → %s (%s)\n",
				operationToString(arp.Operation),
				arp.SourceProtAddress,
				arp.SourceHwAddress,
				arp.DstProtAddress,
				arp.DstHwAddress,
			)
		}
	}
}

// Traduz o código da operação ARP para texto
func operationToString(op layers.ARPOperation) string {
	switch op {
	case layers.ARPRequest:
		return "Request"
	case layers.ARPReply:
		return "Reply"
	default:
		return "Unknown"
	}
}
