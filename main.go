package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var src = flag.String("src", "", "tmx file")

//var dst = flag.String("dst", "", "json file")

func main() {
	flag.Parse()

	data, err := ioutil.ReadFile(*src)
	if err != nil {
		log.Fatal("fail ReadFile", err)
	}

	var m Map
	if err := xml.Unmarshal(data, &m); err != nil {
		log.Fatal("fail Unmarshal", err)
	}
	if err := m.Validation(); err != nil {
		log.Fatal("fail Validation", err)
	}

	for _, t := range m.Tilesets {
		t.Image = t.IImage
		t.TilesJSON = make(map[string]*Tile)
		for _, tile := range t.Tiles {
			t.TilesJSON[strconv.FormatInt(tile.Id, 10)] = tile
			tile.Image = tile.Image
			arr := strings.Split(tile.Terrain, ",")
			for _, s := range arr {
				i, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					i = -1
				}
				tile.TerrainJSON = append(tile.TerrainJSON, i)
			}
			if len(tile.TerrainJSON) != 4 {
				tile.TerrainJSON = nil
			}
		}
		log.Println(t.Image, t.IImage)
	}

	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")
	out.WriteTo(os.Stdout)

	/*data
	ioutil.WriteFile()
	*/
}
