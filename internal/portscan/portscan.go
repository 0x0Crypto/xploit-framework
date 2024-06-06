package portscan

import (
	"context"
	"fmt"
	"strconv"
	"time"

	portscanner "github.com/anvie/port-scanner"
	"github.com/fatih/color"
)

func Scan(ctx context.Context, targetHost string, maxPort int) {
	select {
	case <-ctx.Done():
		fmt.Println("Portscan cancelled")
		return
	default:
		cyan := color.New(color.FgCyan, color.Bold).PrintfFunc()

		color.Blue("Scanning: " + targetHost + " range " + string(strconv.Itoa(maxPort)) + "...")

		ps := portscanner.NewPortScanner(targetHost, 1*time.Second, 15)
		openedPorts := ps.GetOpenedPort(0, maxPort)

		for _, port := range openedPorts {
			select {
			case <-ctx.Done():
				fmt.Println("Portscan cancelled")
				return
			default:
				cyan("%d [open] %s\n", port, ps.DescribePort(port))
			}
		}
	}
}
