package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/oniproject/tmxconverter/tmx"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var src = flag.String("src", "", ".tmx file")
var wlk = flag.String("wlk", "", ".wlk.txt file")

var wlkDot = flag.String("D", ".", "wlk dot")
var wlkX = flag.String("X", "X", "wlk X")

//var dst = flag.String("dst", "", "json file")

func main() {
	flag.Parse()

	log.SetFlags(log.Lshortfile)

	data, err := ioutil.ReadFile(*src)
	if err != nil {
		log.Fatal("fail ReadFile", err)
	}

	var m tmx.Map
	if err := xml.Unmarshal(data, &m); err != nil {
		log.Fatal("fail Unmarshal", err)
	}
	if err := m.Validation(); err != nil {
		log.Fatal("fail Validation", err)
	}

	for _, t := range m.Tileset {
		t.Image = t.IImage
		t.TilesJSON = make(map[string]*tmx.Tile)
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

	/*for _, l := range m.Layers {
		log.Println(l)
	}*/

	switch {
	case *wlk != "":
		for _, layer := range m.Layers {
			if layer.Type == "tilelayer" && layer.Name == *wlk {
				arr, _ := layer.Data.Data()
				for k, v := range arr {
					if k != 0 && k%m.Width == 0 {
						fmt.Print("\n")
					}
					if v == 0 || v == -1 {
						fmt.Print(*wlkDot)
					} else {
						fmt.Print(*wlkX)
					}
				}
				return
			}
		}
	default:
		b, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		var out bytes.Buffer
		json.Indent(&out, b, "", "  ")
		out.WriteTo(os.Stdout)
	}

	/*data
	ioutil.WriteFile()
	*/
}
