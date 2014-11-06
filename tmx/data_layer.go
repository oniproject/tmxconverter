package tmx

import "encoding/xml"

// <layer>, <objectgroup>, <imagelayer>
//
// All <tileset> tags shall occur before the first <layer> tag so that parsers may rely on having the tilesets before needing to resolve tiles.
//
// Can contain: properties, data
type Layer struct {
	XMLName xml.Name `json:"-"`

	Type string
	//Name    string `xml:"name,attr"`
	//Width   int    `xml:"width,attr"`
	//Height  int    `xml:"height,attr"`

	// <layer>
	Name    string  `xml:"name,attr"`    // name: The name of the layer.
	X       int64   `xml:"x,attr"`       // x: The x coordinate of the layer in tiles. Defaults to 0 and can no longer be changed in Tiled Qt.
	Y       int64   `xml:"y,attr"`       // y: The y coordinate of the layer in tiles. Defaults to 0 and can no longer be changed in Tiled Qt.
	Width   int64   `xml:"width,attr"`   // width: The width of the layer in tiles. Traditionally required, but as of Tiled Qt always the same as the map width.
	Height  int64   `xml:"height,attr"`  // height: The height of the layer in tiles. Traditionally required, but as of Tiled Qt always the same as the map height.
	Opacicy float64 `xml:"opacity,attr"` // opacity: The opacity of the layer as a value from 0 to 1. Defaults to 1.
	Visible bool    `xml:"visible,attr"` // visible: Whether the layer is shown (1) or hidden (0). Defaults to 1.

	Color string

	Properties Properties `xml:"properties>property"`

	Data    Data
	Objects []Object

	Image Image `xml:"image"`
}
