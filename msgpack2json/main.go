package main

import (
	"encoding/json"
	"github.com/ugorji/go/codec"
	"io"
	"log"
	"os"
	"reflect"
)

func main() {
	handle := &codec.MsgpackHandle{}
	// json encode chokes on map[interface{}]interface{} without
	// this
	handle.DecodeOptions.MapType = reflect.TypeOf(map[string]interface{}(nil))
	dec := codec.NewDecoder(os.Stdin, handle)
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
