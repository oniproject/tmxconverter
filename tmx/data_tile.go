package tmx

import "encoding/xml"

// <tile>
//
// Can contain: properties, image (since 0.9.0)
type Tile struct {
	XMLName xml.Name `xml:"tile" json:"-"`

	Id          int64   `xml:"id,attr" json:"-"`      // id: The local tile ID within its tileset.
	Terrain     string  `xml:"terrain,attr" json:"-"` // terrain: Defines the terrain type of each corner of the tile, given as comma-separated indexes in the terrain types array in the order top-left, top-right, bottom-left, bottom-right. Leaving out a value means that corner has no terrain. (optional) (since 0.9.0)
	TerrainJSON []int64 `xml:"-" json:"terrain,omitempty"`
	Probability float32 `xml:"probability,attr" json:"probability,omitempty"` //probability: A percentage indicating the probability that this tile is chosen when it competes with others while editing with the terrain tool. (optional) (since 0.9.0)

	Properties Properties `xml:"properties>property" json:"properties,omitempty"`
	IImage     Image      `xml:"image" json:"-"`
	Image      `xml:"-"`
}
