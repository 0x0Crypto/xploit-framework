package subrecon

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/fatih/color"
)

func SubRecon(ctx context.Context, target, wordlistPath string) {
	var fullDomain string

	file, err := os.Open(wordlistPath)
	if err != nil {
		color.Red("ERROR: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	color.Blue("Searching...")

	var wg sync.WaitGroup
	var d1 []byte
	var foundOnline bool

	wg.Add(1)
	saveComplete := make(chan struct{})
	defer close(saveComplete)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				if foundOnline {
					if err := os.WriteFile(target+".txt", d1, 0644); err != nil {
						color.Red("ERROR: ", err)
					} else {
						fmt.Println("Saved to: ", target+".txt")
					}
				}
				return
			case <-saveComplete:
				return
			}
		}
	}()

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			fmt.Println("Subrecon cancelled")
			wg.Wait()
			return
		default:
			fmt.Printf("\r%s", scanner.Text())

			currentLine := scanner.Text()
			fullDomain = currentLine + "." + target

			ipAddresses, err := net.LookupHost(fullDomain)
			if err == nil {
				for _, ip := range ipAddresses {
					if !isRedirect(fullDomain, ip) {
						color.Green("\n\n[+] ONLINE: " + fullDomain + "\n\n")
						foundOnline = true
						d1 = append(d1, []byte(fullDomain+"\n")...)
					} else {
						color.Yellow("\n\n[!] IS FALSE POSITIVE: " + fullDomain + "\n\n")
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		color.Red("ERROR: ", err)
	}

	wg.Wait()
}

func isRedirect(domain, ip string) bool {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get("http://" + domain)
	if err != nil {
		return false
	}

	return resp.StatusCode >= 300 && resp.StatusCode < 400
}

// ex: subrecon nmap.org /usr/share/wordlists/dnsmap.txt
