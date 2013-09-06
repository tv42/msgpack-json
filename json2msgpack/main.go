package main

import (
	"encoding/json"
	"github.com/ugorji/go/codec"
	"io"
	"log"
	"os"
)

func main() {
	handle := &codec.MsgpackHandle{}
	dec := json.NewDecoder(os.Stdin)
	enc := codec.NewEncoder(os.Stdout, handle)
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
