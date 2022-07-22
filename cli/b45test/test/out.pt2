package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pschlump/gintools/qr_svr2/base45"
)

var encode = flag.String("encode", "", "file to encode")
var decode = flag.String("decode", "", "file to encode")
var output = flag.String("output", "", "file to encode")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "qr_gen_server: Usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse() // Parse CLI arguments to this, --cfg <name>.json

	fns := flag.Args()
	if len(fns) != 0 {
		fmt.Printf("Extra arguments are not supported [%s]\n", fns)
		os.Exit(1)
	}

	if *encode != "" {
		buf, err := ioutil.ReadFile(*encode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s on %s\n", err, *encode)
			os.Exit(1)
		}
		s := base45.Base45Encode([]byte(buf))
		if *output == "" {
			fmt.Printf("%s\n", s)
		} else {
			err := ioutil.WriteFile(*output, []byte(s), 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s on %s\n", err, *encode)
				os.Exit(1)
			}
		}
	} else if *decode != "" {
		buf, err := ioutil.ReadFile(*decode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s on %s\n", err, *decode)
			os.Exit(1)
		}
		bb := base45.Base45Decode(string(buf))
		if *output == "" {
			fmt.Printf("%s\n", bb)
		} else {
			err := ioutil.WriteFile(*output, bb, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s on %s\n", err, *decode)
				os.Exit(1)
			}
		}
	}
}
