package main

import (
	"encoding/json"
	"github.com/aaronland/go-ucd/v13"
	"log"
	"strings"
	"syscall/js"
)

var ucd_func js.Func

func main() {

	ucd_func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		str := args[0].String()
		chars := strings.Split(str, "")

		ucd_map := make(map[string]ucd.UCDName)

		for _, char := range chars {
			n := ucd.Name(char)
			ucd_map[char] = n
		}

		enc_map, err := json.Marshal(ucd_map)

		if err != nil {
			return nil
		}

		return string(enc_map)
	})

	defer ucd_func.Release()

	js.Global().Set("ucd", ucd_func)

	c := make(chan struct{}, 0)

	log.Println("UCD function initialized")
	<-c
}
