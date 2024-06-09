package geoip

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
    "strings"

	"github.com/fatih/color"
)

type GeoIP struct {
    Status     string `json:"status"`
    Target     string `json:"query"`
    Country    string `json:"country"`
    Region     string `json:"region"`
    RegionName string `json:"regionName"`
    City       string `json:"city"`
    Zip        string `json:"zip"`
    Lat        string `json:"lat"`
    Lon        string `json:"lon"`
    Isp        string `json:"isp"`
}

func Locate(target string) {
    var geoip GeoIP

	apiUrl := fmt.Sprintf("http://ip-api.com/json/%s", target)
	req, err := http.Get(apiUrl)
	if err != nil {
	    color.Red("ERROR: ", err.Error())
    }
    defer req.Body.Close()

	resp, err := io.ReadAll(req.Body)
	if err != nil {
       color.Red("ERROR: ", err.Error())
    }

    json.Unmarshal(resp, &geoip)

    if strings.ToLower(geoip.Status) != "fail" {
        fmt.Printf("TARGET: %s\nCOUNTRY: %s\nREGION: %s\nREGION NAME: %s\nCITY: %s\nZIP: %s\nLAT %s LON %s\nISP: %s\n",
                geoip.Target,
                geoip.Country,
                geoip.Region,
                geoip.RegionName,
                geoip.City,
                geoip.Zip,
                geoip.Lat,
                geoip.Lon,
                geoip.Isp,
        )
    } else {
        color.Red("ERROR: hostname invalid!")
    }
}
