package main

import (
	"bufio"
	"encoding/json"
	"net/url"
	"os"
)

func main() {
	enc := json.NewEncoder(os.Stdout)

	var uri string
	if len(os.Args[1]) > 0 {
		uri = os.Args[1]
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			uri += scanner.Text()
		}
		orDie(scanner.Err(), enc)
	}
	u, err := url.Parse(uri)
	orDie(err, enc)
	enc.Encode(u)
}

func orDie(err error, enc *json.Encoder) {
	if err != nil {
		enc.Encode(struct{ Error string }{err.Error()})
		os.Exit(1)
	}
}
