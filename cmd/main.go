package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	scraper "github.com/yu81/thispersondoesnotexist-scraper"
)

func get(u string) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(nameFromTime(), all, 0644)
}

func nameFromTime() string {
	t := time.Now()
	tFmt := t.Format("20060102150405")
	return "thispersondoesnotexist_" + tFmt + ".jpg"
}

func main() {
	c := flag.Int("c", 1, "number of images will be fetched")
	flag.Parse()
	if c != nil && *c <= 0 {
		*c = 1
	}
	for i := 0; i < *c; i++ {
		err := get(scraper.ImageURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	}
}
