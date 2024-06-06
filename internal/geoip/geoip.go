package geoip

import (
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/pretty"
)

func Locate(targets []string) ([]string, error) {
	var results []string

	for _, target := range targets {
		apiUrl := fmt.Sprintf("http://ip-api.com/json/%s", target)
		req, err := http.Get(apiUrl)
		if err != nil {
			return []string{}, err
		}
		defer req.Body.Close()

		resp, err := io.ReadAll(req.Body)
		if err != nil {
			return []string{}, err
		}
		finalResult := pretty.Pretty(resp)

		results = append(results, string(finalResult))
	}

	return results, nil
}
