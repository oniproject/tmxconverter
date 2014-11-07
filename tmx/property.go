package tmx

import "encoding/xml"

type Properties []Property

type Property struct {
	XMLName xml.Name `xml:"property" json:"-"`

	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}
