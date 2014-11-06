package tmx

import (
	"encoding/xml"
	"io/ioutil"
)

func LoadTMX(fname string) (m Map, err error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return
	}

	err = xml.Unmarshal(data, &m)

	if err != nil {
		return
	}

	return
}

// <map>
//
// The tilewidth and tileheight properties determine the general grid size of the map. The individual tiles may have different sizes. Larger tiles will extend at the top and right (anchored to the bottom left).
//
// Can contain: properties, tileset, layer, objectgroup, imagelayer
type Map struct {
	XMLName xml.Name `xml:"map"`

	Version     string `xml:"version,attr"`     // version: The TMX format version, generally 1.0.
	Orientation string `xml:"orientation,attr"` // orientation: Map orientation. Tiled supports "orthogonal", "isometric" and "staggered" (since 0.9) at the moment.
	Width       int    `xml:"width,attr"`       // width: The map width in tiles.
	Height      int    `xml:"height,attr"`      // height: The map height in tiles.
	TileWidth   int    `xml:"tilewidth,attr"`   // tilewidth: The width of a tile.
	TileHeight  int    `xml:"tileheight,attr"`  // tileheight: The height of a tile.
	// backgroundcolor: The background color of the map. (since 0.9, optional)
	BackgroundColor string `xml:"backgroundcolor,attr" json:"backgroundcolor,omitempty"`
	// renderorder: The order in which tiles on tile layers are rendered. Valid values are right-down (the default), right-up, left-down and left-up.
	// In all cases, the map is drawn row-by-row. (since 0.10, but only supported for orthogonal maps at the moment)
	RenderOrder string `xml:"renderorder,attr" json:"renderorder"`

	Properties Properties `xml:"properties>property" json:"properties,omitempty"`
	Tileset    []*Tileset `xml:"tileset" json:"tileset"`

	// layer, objectgroup, imagelayer
	Layers []Layer `xml:",any"`
}

func (m Map) LayerByName(name, t string) *Layer {
	for _, layer := range m.Layers {
		if layer.Name != name {
			continue
		}
		if t != "" && layer.Type != t {
			continue
		}
		return &layer
	}
	return nil
}
