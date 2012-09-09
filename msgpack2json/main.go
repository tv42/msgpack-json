package main

import (
	"github.com/ugorji/go-msgpack"
	"log"
	"os"
	"io"
	"encoding/json"
	"reflect"
)

func main() {
	opts := msgpack.DefaultDecoderContainerResolver
	// json encode chokes on map[interface{}]interface{} without
	// this
	opts.MapType =  reflect.TypeOf(map[string]interface{}(nil))
	dec := msgpack.NewDecoder(os.Stdin, &opts)
	enc := json.NewEncoder(os.Stdout)
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
