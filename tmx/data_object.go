package tmx

//import "encoding/xml"

type Point struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

// <object>
//
// Can contain: properties, ellipse (since 0.9.0), polygon, polyline, image
type Object struct {
	Name     string  `xml:"name,attr" json:"name"`                   // name: The name of the object. An arbitrary string.
	Type     string  `xml:"type,attr" json:"type,omitempty"`         // type: The type of the object. An arbitrary string.
	X        int64   `xml:"x,attr" json:"x"`                         // x: The x coordinate of the object in pixels.
	Y        int64   `xml:"y,attr" json:"y"`                         // y: The y coordinate of the object in pixels.
	Width    int64   `xml:"width,attr" json:"width,omitempty"`       // width: The width of the object in pixels (defaults to 0).
	Height   int64   `xml:"height,attr" json:"height,omitempty"`     // height: The height of the object in pixels (defaults to 0).
	Rotation float64 `xml:"rotation,attr" json:"rotation,omitempty"` // rotation: The rotation of the object in degrees clockwise (defaults to 0). (on git master)
	GID      int64   `xml:"gid,attr" json:"gid,omitempty"`           // gid: An reference to a tile (optional).
	Visible  bool    `xml:"visible,attr" json:"visible,omitempty"`   // TODO visible: Whether the object is shown (1) or hidden (0). Defaults to 1. (since 0.9.0)

	Properties Properties `xml:"properties>property" json:"properties,omitempty"`
	Ellipse    bool       `json:"ellipse,omitempty"`
	Polyline   []Point    `json:"polyline,omitempty"`
	Polygon    []Point    `json:"polygon,omitempty"`

	// TODO add some others
}
