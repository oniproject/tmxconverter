package tmx

import (
	"encoding/xml"
	"github.com/oniproject/geom"
	"io"
	"strconv"
	"strings"
)

type Data struct {
	Encoding    string  `xml:"encoding,attr"`
	Compression string  `xml:"compression,attr"`
	CharData    []byte  `xml:",chardata"`
	Tiles       []*Tile `xml:"tile" json:"-"`
}

func (data *Data) Data() (ii []int64, err error) {
	//b := bytes.NewReader(data.CharData)
	switch data.Encoding {
	case "csv":
		//log.Println(string(data.CharData))
		arr := strings.Split(string(data.CharData), ",")
		for _, el := range arr {
			v, err := strconv.ParseInt(strings.TrimSpace(el), 10, 64)
			if err != nil {

				//log.Println("!!!!!!!!!!!", err)
				return nil, err
			}
			ii = append(ii, v)
			//arr[i],err =
		}
		//reader := csv.NewReader(b)
		//return reader.ReadAll()
	case "base64":
		//decoder := base64.NewDecoder(base64.StdEncoding,b)
		panic("fail Encoding base64")
	default:
		panic("fail Encoding")
	}
	//return []int64{}, nil
	return
}

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

// XXX not return polyline
// TODO check rotation
func (layer *Layer) ObjectsByPoint(x, y float64) []*Object {
	ret := []*Object{}
	for _, object := range layer.Objects {
		if object.Rotation != 0 {
			panic("Rotation not support now")
		}
		switch {
		case object.Ellipse:
			dx, dy, w2, h2 :=
				x-(object.X+object.Width/2),
				y-(object.Y+object.Height/2),
				object.Width*object.Width,
				object.Height*object.Height

			if (dx*dx)/w2+(dy*dy)/h2 <= 1 {
				ret = append(ret, &object)
			}
		case len(object.Polygon) != 0:
		case len(object.Polyline) != 0:
			// pass
			poly := geom.Polygon{}
			for _, v := range object.Polyline {
				poly.AddVertex(geom.Coord{v.X, v.Y})
			}
			if poly.ContainsCoord(geom.Coord{x, y}) {
				ret = append(ret, &object)
			}
		case object.GID != 0:
		default: // rect
			if x >= object.X && x <= object.X+object.Width &&
				y >= object.Y && y <= object.Y+object.Height {
				ret = append(ret, &object)
			}
		}
	}
	return ret
}

func (layer *Layer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	//log.Println("ATTR", start)

	switch start.Name.Local {
	case "layer":
		layer.Type = "tilelayer"
	case "objectgroup", "imagelayer":
		layer.Type = start.Name.Local
	}

	layer.Visible = true
	layer.Opacicy = 1

	for _, elem := range start.Attr {
		//log.Println("\telem", elem)
		var i int64
		var f float64
		switch elem.Name.Local {
		case "name":
			layer.Name = elem.Value
		case "color":
			layer.Color = elem.Value
		case "x":
			i, err = strconv.ParseInt(elem.Value, 10, 0)
			layer.X = i
		case "y":
			i, err = strconv.ParseInt(elem.Value, 10, 0)
			layer.Y = i
		case "width":
			i, err = strconv.ParseInt(elem.Value, 10, 0)
			layer.Width = i
		case "height":
			i, err = strconv.ParseInt(elem.Value, 10, 0)
			layer.Height = i
		case "opacity":
			f, err = strconv.ParseFloat(elem.Value, 32)
			layer.Opacicy = f
		case "visible":
			i, err = strconv.ParseInt(elem.Value, 10, 0)
			layer.Visible = i != 0
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
			case "image":
				err = d.DecodeElement(&layer.Image, &token)
			case "data":
				err = d.DecodeElement(&layer.Data, &token)
			case "object":
				var o Object
				err = d.DecodeElement(&o, &token)
				//log.Println("OBJECT", o)
				layer.Objects = append(layer.Objects, o)
			case "properties":
				pro := &struct {
					Properties Properties `xml:"property" json:"properties,omitempty"`
				}{}
				err = d.DecodeElement(pro, &token)
				layer.Properties = pro.Properties
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
