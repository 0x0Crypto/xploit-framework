package cmd

import (
	"github.com/chzyer/readline"
	"github.com/fatih/color"
)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("dnsrecon"),
	readline.PcItem("bannergrab"),
	readline.PcItem("subrecon"),
	readline.PcItem("dirsearch"),
	readline.PcItem("whois"),
	readline.PcItem("robots"),
	readline.PcItem("geoip"),
	readline.PcItem("portscan"),
	readline.PcItem("banner"),
	readline.PcItem("clear"),
	readline.PcItem("exit"),
)

var rl *readline.Instance

func ConfigureInput() readline.Config {
    white := color.New(color.FgWhite, color.Bold).SprintfFunc()

	configRl := readline.Config{
		Prompt:            white("Xploit >>> "),
		HistoryFile:       ".readline_history",
		AutoComplete:      completer,
		HistorySearchFold: true,
	}

	return configRl
}
