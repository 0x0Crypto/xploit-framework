package bannergrab

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/fatih/color"
)

func BannerGrab(target string, port int) (string, error) {
	var conn net.Conn
	var err error

	if port == 443 {
		color.Green("Establishing TLS connection...")
		conn, err = tls.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}, "tcp", fmt.Sprintf("%s:%d", target, port), &tls.Config{})
	} else {
		color.Green("Establishing normal connection...")
		conn, err = net.DialTimeout("tcp", fmt.Sprintf("%s:%d", target, port), 5*time.Second)
	}
	if err != nil {
		return "", err
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %s\r\n\r\n", target)

	buffer := make([]byte, 2048)
	n, _ := conn.Read(buffer)

	response := string(buffer[:n])

	return response, nil
}
