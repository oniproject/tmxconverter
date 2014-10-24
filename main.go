package main

import (
	"encoding/xml"
	"flag"
	"io/ioutil"
	"log"
)

var src = flag.String("src", "", "tmx file")

//var dst = flag.String("dst", "", "json file")

func main() {
	flag.Parse()

	data, err := ioutil.ReadFile(*src)
	if err != nil {
		log.Fatal("fail ReadFile", err)
	}

	var x Map
	if err := xml.Unmarshal(data, &x); err != nil {
		log.Fatal("fail Unmarshal", err)
	}
	if err := x.Validation(); err != nil {
		log.Fatal("fail Validation", err)
	}

	log.Println(x)

	/*data
	ioutil.WriteFile()
	*/
}
