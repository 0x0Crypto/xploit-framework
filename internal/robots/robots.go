package robots

import (
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func SearchRobots(targets []string) ([]string, error) {
	var resp []string

	for _, target := range targets {
		target = strings.TrimSuffix(target, "/")
		req, err := http.Get(target + "/robots.txt")
		if err != nil {
			return nil, err
		}
		defer req.Body.Close()

		if req.StatusCode != 404 {
			respByte, err := io.ReadAll(req.Body)
			if err != nil {
				return nil, err
			}

			resp = append(resp, string(respByte))
		} else {
			return nil, errors.Errorf("Robots.txt not found!")
		}
	}

	return resp, nil
}
