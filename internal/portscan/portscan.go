package portscan

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"net"
	"strconv"
	"time"
)

func Scan(ctx context.Context, targetHost string, maxPort int) {
    portStr := strconv.Itoa(maxPort)

	color.Blue("Scanning: " + targetHost + " range " + portStr + "...")

	for i := 0; i <= maxPort; i++ {
		select {
		case <-ctx.Done():
                fmt.Println("PortScan canceled")
                return
		default:
			ip := targetHost + ":" + strconv.Itoa(i)

			conn, err := net.DialTimeout("tcp", ip, time.Duration(300)*time.Millisecond)
			if err == nil {
				defer conn.Close()
				color.Cyan("OPEN %v tcp\n", i)
			}
		}
	}
}
