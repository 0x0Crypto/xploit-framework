package dirsearch

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

func DirSearch(ctx context.Context, target, wordlistPath string) {
	file, err := os.Open(wordlistPath)
	if err != nil {

		color.Red("ERROR: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	color.Blue("Searching... Please wait")

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			fmt.Println("\nDirsearch cancelled")
			return
		default:
			target = strings.TrimSuffix(target, "/")
			wordlist := strings.TrimSpace(scanner.Text())
			fmt.Printf("\r%s", wordlist)

			req, err := http.Get(target + "/" + wordlist)
			if err != nil {
				color.Red("\n\n" + "ERROR: " + err.Error() + "\n\n")
				continue
			}

			if req.StatusCode != 404 {
				color.Green("\n\n" + req.Status + " " + target + "/" + wordlist + "\n\n")
			}

			req.Body.Close()
		}
	}
}
