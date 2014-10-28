package tmx

import "encoding/xml"

// <image>
//
// As of the current version of Tiled Qt, each tileset has a single image associated with it, which is cut into smaller tiles based on the attributes defined on the tileset element.
// Later versions may add support for adding multiple images to a single tileset, as is possible in Tiled Java.
//
// Can contain: data (since 0.9.0)
type Image struct {
	XMLName xml.Name `xml:"image" json:"-"`

	ImageFormat      string `xml:"format,attr" json:"imageformat,omitempty"`     //format: Used for embedded images, in combination with a data child element. Valid values are file extensions like png, gif, jpg, bmp, etc. (since 0.9.0)
	Id               int    `xml:"id,attr" json:"-"`                             // id: Used by some versions of Tiled Java. Deprecated and unsupported by Tiled Qt.
	ImageSource      string `xml:"source,attr" json:"image,omitempty"`           // source: The reference to the tileset image file (Tiled supports most common image formats).
	TransparentColor string `xml:"trans,attr" json:"transparentcolor,omitempty"` // trans: Defines a specific color that is treated as transparent (example value: "#FF00FF" for magenta). Up until Tiled 0.10, this value is written out without a # but this is planned to change.
	Width            int    `xml:"width,attr" json:"imagewidth,omitempty"`       // width: The image width in pixels (optional, used for tile index correction when the image changes)
	Height           int    `xml:"height,attr" json:"imageheight,omitempty"`     // height: The image height in pixels (optional)
	// TODO data
}
