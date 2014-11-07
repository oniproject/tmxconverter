package tmx

import (
	"encoding/xml"
	"io"
	"strconv"
	"strings"
)

type Point struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

// <object>
//
// Can contain: properties, ellipse (since 0.9.0), polygon, polyline, image
type Object struct {
	XMLName xml.Name `xml:"object" json:"-"`

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

func (obj *Object) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	obj.Visible = true

	for _, elem := range start.Attr {
		//log.Println("\telem", elem)
		var i int64
		var f float64
		switch elem.Name.Local {
		case "name":
			obj.Name = elem.Value
		case "type":
			obj.Type = elem.Value
		case "x":
			i, err = strconv.ParseInt(elem.Value, 10, 0)
			obj.X = i
		case "y":
			i, err = strconv.ParseInt(elem.Value, 10, 0)
			obj.Y = i
		case "width":
			i, err = strconv.ParseInt(elem.Value, 10, 0)
			obj.Width = i
		case "height":
			i, err = strconv.ParseInt(elem.Value, 10, 0)
			obj.Height = i
		case "rotation":
			f, err = strconv.ParseFloat(elem.Value, 32)
			obj.Rotation = f
		case "gid":
			i, err = strconv.ParseInt(elem.Value, 10, 0)
			obj.GID = i
		case "visible":
			i, err = strconv.ParseInt(elem.Value, 10, 0)
			obj.Visible = i != 0
		}
		if err != nil {
			return err
		}
	}

	var t xml.Token
	for t, err = d.Token(); err == nil; t, err = d.Token() {
		//log.Printf("TOKEN %T %v\n", t, t)
		switch token := t.(type) {
		case xml.StartElement:
			switch token.Name.Local {
			// TODO image
			case "ellipse":
				err = d.Skip()
				obj.Ellipse = true
			case "polygon", "polyline":
				p := &struct {
					Points string `xml:"points,attr"`
				}{}
				err = d.DecodeElement(p, &token)
				if err != nil {
					return err
				}
				p1 := strings.Split(p.Points, " ")
				for _, pX := range p1 {
					p2 := strings.Split(pX, ",")
					x, _ := strconv.ParseInt(p2[0], 10, 64)
					y, _ := strconv.ParseInt(p2[1], 10, 64)

					if token.Name.Local == "polygon" {
						obj.Polygon = append(obj.Polygon, Point{x, y})
					} else {
						obj.Polyline = append(obj.Polyline, Point{x, y})
					}
				}
			case "properties":
				pro := &struct {
					Properties Properties `xml:"property" json:"properties,omitempty"`
				}{}
				err = d.DecodeElement(pro, &token)
				obj.Properties = pro.Properties
			}
			if err != nil {
				return
			}
		}
	}

	if err == io.EOF {
		return nil
	}

	return

	//return d.Skip()
}
