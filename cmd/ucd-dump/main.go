package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aaronland/go-ucd/v14/unicodedata"
	"github.com/aaronland/go-ucd/v14/unihan"
	"github.com/tidwall/pretty"
	"log"
	"os"
)

func main() {

	var include_unihan = flag.Bool("unihan", false, "Include Unihan data.")
	var make_pretty = flag.Bool("pretty", false, "Prettify JSON output.")

	flag.Parse()

	var dump map[string]string
	dump = unicodedata.UCD

	if *include_unihan {

		// sorting?

		for k, v := range unihan.UCDHan {
			dump[k] = v
		}
	}

	body, err := json.Marshal(dump)

	if err != nil {
		log.Fatal(err)
	}

	if *make_pretty {
		body = pretty.Pretty(body)
	}

	fmt.Printf("%s\n", body)
	os.Exit(0)
}
