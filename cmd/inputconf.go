package cmd

import (
	"github.com/chzyer/readline"
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
	configRl := readline.Config{
		Prompt:            "Xploit >>> ",
		HistoryFile:       ".readline_history",
		AutoComplete:      completer,
		HistorySearchFold: true,
	}

	return configRl
}
