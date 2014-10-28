package tmx

import (
	"encoding/xml"
	"io/ioutil"
)

func LoadTSX(fname string) (ts Tileset, err error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return
	}

	err = xml.Unmarshal(data, &ts)

	if err != nil {
		return
	}

	return
}

// <tileoffset>
//
// This element is used to specify an offset in pixels, to be applied when drawing a tile from the related tileset. When not present, no offset is applied.
type TileOffset struct {
	XMLName xml.Name `xml:"tileoffset" json:"-"`

	X int `xml:"x,attr" json:"x"` // x: Horizontal offset in pixels
	Y int `xml:"y,attr" json:"y"` // y: Vertical offset in pixels (positive is down)
}

// <terrain>
//
// Can contain: properties
type Terrain struct {
	XMLName xml.Name `xml:"terrain" json:"-"`

	Name       string     `xml:"name,attr" json:"name"` // name: The name of the terrain type.
	Tile       int        `xml:"tile,attr" json:"tile"` // tile: The local tile-id of the tile that represents the terrain visually.
	Properties Properties `xml:"properties>property" json:"properties,omitempty"`
}

// <tileset>
//
// Can contain: tileoffset (since 0.8.0), properties (since 0.8.0), image, terraintypes (since 0.9.0), tile
type Tileset struct {
	XMLName xml.Name `xml:"tileset" json:"-"`

	FirstGID int `xml:"firstgid,attr" json:"firstgid"` // firstgid: The first global tile ID of this tileset (this global ID maps to the first tile in this tileset).
	// source: If this tileset is stored in an external TSX (Tile Set XML) file, this attribute refers to that file. That TSX file has the same structure as the attribute as described here.
	// (There is the firstgid attribute missing and this source attribute is also not there. These two attributes are kept in the TMX map, since they are map specific.)
	Source       string     `xml:"source,attr" json:"source,omitempty"`
	Name         string     `xml:"name,attr" json:"name"`             // name: The name of this tileset.
	TileWidth    int        `xml:"tilewidth,attr" json:"tilewidth"`   // tilewidth: The (maximum) width of the tiles in this tileset.
	TileHeight   int        `xml:"tileheight,attr" json:"tileheight"` // tileheight: The (maximum) height of the tiles in this tileset.
	Spacing      int        `xml:"spacing,attr" json:"spacing"`       // spacing: The spacing in pixels between the tiles in this tileset (applies to the tileset image).
	Margin       int        `xml:"margin,attr" json:"margin"`         // margin: The margin around the tiles in this tileset (applies to the tileset image).
	TileOffset   TileOffset `xml:"tileoffset" json:"tileoffset"`
	Properties   Properties `xml:"properties>property" json:"properties,omitempty"`
	IImage       Image      `xml:"image" json:"-"`
	Image        `xml:"-"`
	TerrainTypes []*Terrain       `xml:"terraintypes>terrain" json:"terrains,omitempty"`
	Tiles        []*Tile          `xml:"tile" json:"-"`
	TilesJSON    map[string]*Tile `xml:"-" json:"tiles"`
	// TODO imageheight, imagewidth, tileproperties
}
