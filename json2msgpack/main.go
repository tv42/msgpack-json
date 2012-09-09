package main

import (
	"encoding/json"
	"github.com/ugorji/go-msgpack"
	"io"
	"log"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := msgpack.NewEncoder(os.Stdout)
	for {
		var data interface{}
		err := dec.Decode(&data)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		err = enc.Encode(data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
