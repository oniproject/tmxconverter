package tmx

import (
	//"bytes"
	//"encoding/csv"
	"encoding/xml"
	"io"
	"log"
	"strconv"
	"strings"
)

func (layer *Layer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	log.Println("ATTR", start)

	switch start.Name.Local {
	case "layer":
		layer.Type = "tilelayer"
	case "objectgroup", "imagelayer":
		layer.Type = start.Name.Local
	}

	layer.Visible = true
	layer.Opacicy = 1

	for _, elem := range start.Attr {
		log.Println("\telem", elem)
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
		log.Printf("TOKEN %T %v\n", t, t)
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
				log.Println("OBJECT", o)
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

				log.Println("!!!!!!!!!!!", err)
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

func (obj *Object) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	obj.Visible = true

	for _, elem := range start.Attr {
		log.Println("\telem", elem)
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
		log.Printf("TOKEN %T %v\n", t, t)
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
