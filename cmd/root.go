package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	// Internal Recon Libs
	"xploit/internal/bannergrab"
	"xploit/internal/dirsearch"
	"xploit/internal/dnsrecon"
	"xploit/internal/geoip"
	"xploit/internal/portscan"
	"xploit/internal/robots"
	"xploit/internal/subrecon"

	// External

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/likexian/whois"
)

type helpOption struct {
	Command     string
	Description string
}

var helpOptions = []helpOption{
	{"dnsrecon", "DNS Recognition"},
	{"bannergrab", "Host Banner Grab"},
	{"subrecon", "Subdomain scan"},
	{"dirsearch", "Directory scan"},
	{"whois", "Whois scan"},
	{"robots", "Detect robots.txt"},
	{"geoip", "Geolocate IP Address"},
	{"portscan", "Scan open ports from target host"},
	{"banner", "Show the Xploit banner"},
	{"clear", "Clear screen"},
	{"exit", "Exit from Xploit"},
}

// colors
var (
	red      = color.New(color.FgRed, color.Italic).PrintlnFunc()
	boldBlue = color.New(color.FgBlue, color.Bold).PrintlnFunc()
)

func InteractiveShell(ctx context.Context) {
	configRl := ConfigureInput()

	ShowAscii()

	for {
		rl, err := readline.NewEx(&configRl)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		select {
		case <-ctx.Done():
			return
		default:
			line, err := rl.Readline()
			if err != nil {
				if strings.Contains(err.Error(), "Interrupt") {
					fmt.Println("Type 'exit' to close.")
				} else {
					fmt.Println(err.Error())
				}
				continue
			}

			commandLine := line
			commandLine = strings.TrimSpace(commandLine)

			if commandLine == "" {
				continue
			}

			commandSplit := strings.Fields(commandLine)
			if len(commandSplit) == 0 {
				color.Red("Command not found!")
				continue
			}

			command := commandSplit[0]

			switch command {
			case "help":
				fmt.Println("COMMANDS:")
				for _, option := range helpOptions {
					red(option.Command + " => " + option.Description)
				}
			case "dnsrecon":
				if len(commandSplit[1:]) == 0 {
					color.Red("Use: dnsrecon WWW.URL1.COM WWW.URL2.COM .........")
					continue
				}

				hosts := commandSplit[1:]
				results, err := dnsrecon.Dnsrecon(hosts)
				if err != nil {
					color.Red("ERROR: " + err.Error())
					continue
				}

				for domain, addrs := range results {
					fmt.Println(domain)
					for _, addr := range addrs {
						fmt.Println("-", addr)
					}
				}
			case "bannergrab":
				if len(commandSplit) < 3 {
					color.Red("Invalid URL format. Use: bannergrab WWW.URL.COM PORT")
					continue
				}

				url := commandSplit[1]
				portStr := commandSplit[2]

				port, err := strconv.Atoi(portStr)
				if err != nil {
					color.Red("ERROR: " + err.Error())
					continue
				}

				results, err := bannergrab.BannerGrab(url, port)
				if err != nil {
					color.Red("ERROR: " + err.Error())
					continue
				}

				color.Blue("===\n" + url + ":" + strconv.Itoa(port))
				color.Cyan(results)
			case "robots":
				if len(commandSplit[1:]) == 0 {
					color.Red("Use: robots httpX://WWW.URL1.COM httpX://WWW.URL2.COM .........")
					continue
				}

				urls := commandSplit[1:]
				resps, err := robots.SearchRobots(urls)
				if err != nil {
					color.Red("ERROR: " + err.Error())
					continue
				}

				for _, url := range urls {
					for _, resp := range resps {
						color.Blue("===\n" + url)
						color.Cyan(resp)
					}
				}
			case "whois":
				if len(commandSplit) < 2 {
					color.Red("Invalid URL format. Use: whois WWW.URL.COM or ASN")
					continue
				}

				url := commandSplit[1]

				color.Cyan("Whois for " + url + "...")
				result, err := whois.Whois(url)
				if err != nil {
					color.Red("ERROR: " + err.Error())
					continue
				}
				color.Green(result)
			case "dirsearch":
				if len(commandSplit) < 3 {
					color.Red("Invalid URL format. Use: dirsearch httpX://WWW.URL.COM WORLDIST_PATH")
					continue
				}

				url := commandSplit[1]
				wordlistPath := commandSplit[2]

				dirsearch.DirSearch(ctx, url, wordlistPath)
			case "banner":
				ShowAscii()
			case "subrecon":
				if len(commandSplit) < 2 {
					color.Red("Invalid format. Use: subrecon URL.COM WORLDIST_PATH")
					continue
				}

				url := commandSplit[1]
				wordlistPath := commandSplit[2]

				subrecon.SubRecon(ctx, url, wordlistPath)
			case "geoip":
				if len(commandSplit) < 2 {
					color.Red("Use: geoip HOSTNAME")
					continue
				}

				addr := commandSplit[1]
				geoip.Locate(addr)
			case "portscan":
				if len(commandSplit) < 3 {
					color.Red("Invalid format. Use: portscan HOST MAX_PORT (EX: 1024)")
					continue
				}

				targetHost := commandSplit[1]
				maxPort, _ := strconv.Atoi(commandSplit[2])

				portscan.Scan(ctx, targetHost, maxPort)
			case "clear":
				if runtime.GOOS == "windows" {
					cmd := exec.Command("cls")
					out, _ := cmd.Output()
					fmt.Println(string(out))
				} else {
					cmd := exec.Command("clear")
					out, _ := cmd.Output()
					fmt.Println(string(out))
				}
			case "exit":
				color.Cyan("Bye (:")
				os.Exit(0)
			default:
				boldBlue("Executing command system...")
				cmd := exec.Command(command, commandSplit[1:]...)
				out, err := cmd.Output()
				if err != nil {
					red("ERROR: ", err.Error())
					continue
				}

				fmt.Println(string(out))
			}
		}
	}
}
