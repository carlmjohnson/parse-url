package main

import (
	"bufio"
	"encoding/json"
	"net/url"
	"os"
)

func main() {
	enc := json.NewEncoder(os.Stdout)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		u, err := url.Parse(scanner.Text())
		orDie(err, enc)
		enc.Encode(u)
	}
	orDie(scanner.Err(), enc)
}

func orDie(err error, enc *json.Encoder) {
	if err != nil {
		enc.Encode(struct{ Error string }{err.Error()})
		os.Exit(1)
	}
}
