package internal

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/likexian/whois"
)

func Whois(ctx context.Context, target string) {
	select {
	case <-ctx.Done():
		fmt.Println("Whois cancelled")
		return
	default:
		result, err := whois.Whois(target)
		if err != nil {
			color.Red("ERROR: ", err)
			return
		}

		color.Green(result)
	}
}
