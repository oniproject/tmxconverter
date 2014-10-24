package main

import "encoding/xml"

// <map>
//
// The tilewidth and tileheight properties determine the general grid size of the map. The individual tiles may have different sizes. Larger tiles will extend at the top and right (anchored to the bottom left).
//
// Can contain: properties, tileset, layer, objectgroup, imagelayer
type Map struct {
	XMLName         xml.Name `xml:"map"`
	Version         string   `xml:"version,attr"`         // version: The TMX format version, generally 1.0.
	Orientation     string   `xml:"orientation,attr"`     //orientation: Map orientation. Tiled supports "orthogonal", "isometric" and "staggered" (since 0.9) at the moment.
	Width           int      `xml:"width,attr"`           // width: The map width in tiles.
	Height          int      `xml:"height,attr"`          // height: The map height in tiles.
	TileWidth       int      `xml:"tilewidth,attr"`       // tilewidth: The width of a tile.
	TileHeight      int      `xml:"tileheight,attr"`      // tileheight: The height of a tile.
	BackgroundColor string   `xml:"backgroundcolor,attr"` // backgroundcolor: The background color of the map. (since 0.9, optional)
	RenderOrder     string   `xml:"renderorder,attr"`     // renderorder: The order in which tiles on tile layers are rendered. Valid values are right-down (the default), right-up, left-down and left-up. In all cases, the map is drawn row-by-row. (since 0.10, but only supported for orthogonal maps at the moment)

	Properties []Property `xml:"properties>property"`
	Tilesets   []Tileset  `xml:"tileset"`

	// layer, objectgroup, imagelayer
	Layers []Layer `xml:",any"`
}

// <tileset>
//
// Can contain: tileoffset (since 0.8.0), properties (since 0.8.0), image, terraintypes (since 0.9.0), tile
type Tileset struct {
	FirstGID string `xml:"firstgid,attr"` // firstgid: The first global tile ID of this tileset (this global ID maps to the first tile in this tileset).
	// source: If this tileset is stored in an external TSX (Tile Set XML) file, this attribute refers to that file. That TSX file has the same structure as the attribute as described here.
	// (There is the firstgid attribute missing and this source attribute is also not there. These two attributes are kept in the TMX map, since they are map specific.)
	Source       string     `xml:"source,attr"`
	Name         string     `xml:"name,attr"`       // name: The name of this tileset.
	TileWidth    int        `xml:"tilewidth,attr"`  // tilewidth: The (maximum) width of the tiles in this tileset.
	TileHeight   int        `xml:"tileheight,attr"` // tileheight: The (maximum) height of the tiles in this tileset.
	Spacing      int        `xml:"spacing,attr"`    // spacing: The spacing in pixels between the tiles in this tileset (applies to the tileset image).
	Margin       int        `xml:"margin,attr"`     // margin: The margin around the tiles in this tileset (applies to the tileset image).
	TileOffset   TileOffset `xml:"tileoffset"`
	Properties   []Property `xml:"properties>property"`
	Image        Image      `xml:"image"`
	TerrainTypes []Terrain  `xml:"terraintypes>terrain"`
	Tiles        []Tile     `xml:"tile"`
}

// <tileoffset>
//
// This element is used to specify an offset in pixels, to be applied when drawing a tile from the related tileset. When not present, no offset is applied.
type TileOffset struct {
	X int `xml:"x,attr"` // x: Horizontal offset in pixels
	Y int `xml:"y,attr"` // y: Vertical offset in pixels (positive is down)
}

// <image>
//
// As of the current version of Tiled Qt, each tileset has a single image associated with it, which is cut into smaller tiles based on the attributes defined on the tileset element.
// Later versions may add support for adding multiple images to a single tileset, as is possible in Tiled Java.
//
// Can contain: data (since 0.9.0)
type Image struct {
	Format           string `xml:"format,attr"` //format: Used for embedded images, in combination with a data child element. Valid values are file extensions like png, gif, jpg, bmp, etc. (since 0.9.0)
	Id               int    `xml:"id,attr"`     // id: Used by some versions of Tiled Java. Deprecated and unsupported by Tiled Qt.
	Source           string `xml:"source,attr"` // source: The reference to the tileset image file (Tiled supports most common image formats).
	TransparentColor string `xml:"trans,attr"`  // trans: Defines a specific color that is treated as transparent (example value: "#FF00FF" for magenta). Up until Tiled 0.10, this value is written out without a # but this is planned to change.
	Width            int    `xml:"width,attr"`  // width: The image width in pixels (optional, used for tile index correction when the image changes)
	Height           int    `xml:"height,attr"` // height: The image height in pixels (optional)
	// TODO data
}

// <terrain>
//
// Can contain: properties
type Terrain struct {
	Name       string     `xml:"name,attr"` // name: The name of the terrain type.
	Tile       int        `xml:"tile,attr"` // tile: The local tile-id of the tile that represents the terrain visually.
	Properties []Property `xml:"properties>property"`
}

// <tile>
//
// Can contain: properties, image (since 0.9.0)
type Tile struct {
	Id          int     `xml:"id,attr"`          // id: The local tile ID within its tileset.
	Terrain     string  `xml:"terrain,attr"`     // terrain: Defines the terrain type of each corner of the tile, given as comma-separated indexes in the terrain types array in the order top-left, top-right, bottom-left, bottom-right. Leaving out a value means that corner has no terrain. (optional) (since 0.9.0)
	Probability float32 `xml:"probability,attr"` //probability: A percentage indicating the probability that this tile is chosen when it competes with others while editing with the terrain tool. (optional) (since 0.9.0)

	Properties []Property `xml:"properties>property"`
	Image      Image      `xml:"image"`
}

// <layer>,
//
// All <tileset> tags shall occur before the first <layer> tag so that parsers may rely on having the tilesets before needing to resolve tiles.
//
// Can contain: properties, data
type Layer struct {
	XMLName xml.Name
	//Name    string `xml:"name,attr"`
	//Width   int    `xml:"width,attr"`
	//Height  int    `xml:"height,attr"`

	// <layer>
	Name    string  `xml:"name,attr"`    // name: The name of the layer.
	X       int     `xml:"x,attr"`       // x: The x coordinate of the layer in tiles. Defaults to 0 and can no longer be changed in Tiled Qt.
	Y       int     `xml:"y,attr"`       // y: The y coordinate of the layer in tiles. Defaults to 0 and can no longer be changed in Tiled Qt.
	Width   int     `xml:"width,attr"`   // width: The width of the layer in tiles. Traditionally required, but as of Tiled Qt always the same as the map width.
	Height  int     `xml:"height,attr"`  // height: The height of the layer in tiles. Traditionally required, but as of Tiled Qt always the same as the map height.
	Opacicy float32 `xml:"opacity,attr"` // opacity: The opacity of the layer as a value from 0 to 1. Defaults to 1.
	Visible int     `xml:"visible,attr"` // visible: Whether the layer is shown (1) or hidden (0). Defaults to 1.
}

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}
