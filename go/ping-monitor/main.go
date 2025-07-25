package main

import (
	"log"
	"net"
	"time"

	"github.com/go-ping/ping"
)

var (
	host     = []string{"google.com", "8.8.8.8", "github.com"}
	mode     = "tcp"
	port     = "80"
	timeout  = 2 * time.Second
	interval = 10 * time.Second
)

func main() {
	log.Println("Iniciando monitoramento...")
	for {
		for _, host := range host {
			switch mode {
			case "icmp":
				icmpCheck(host)
			case "tcp":
				tcpCheck(host, port)
			default:
				log.Fatalf("Modo inválido: %s", mode)
			}
		}
		log.Printf("Aguardando %v antes do próximo ciclo...\n", interval)
		time.Sleep(interval)
	}
}

// ICMP Ping
func icmpCheck(host string) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		log.Printf("[ICMP][%s] ERRO ao criar pinger: %v\n", host, err)
		return
	}
	pinger.Count = 1
	pinger.Timeout = timeout

	start := time.Now()
	err = pinger.Run()
	latency := time.Since(start)

	if err != nil {
		log.Printf("[ICMP][%s] OK - Lantência: %v\n", host, err)
		return
	}

	stats := pinger.Statistics()
	if stats.PacketsRecv > 0 {
		log.Printf("[ICMP][%s] OK - Latência: %v\n", host, latency)
	} else {
		log.Printf("[ICMP][%s] SEM RESPOSTA\n", host)
	}
}

func tcpCheck(host, port string) {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	latency := time.Since(start)

	if err != nil {
		log.Printf("[TCP][%s:%s] FALHA: %v\n", host, port, err)
		return
	}
	defer conn.Close()

	log.Printf("[TCP][%s:%s] OK - Lantência: %v\n", host, port, latency)
}
