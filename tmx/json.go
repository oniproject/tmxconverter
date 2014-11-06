package tmx

import "encoding/json"
import "log"
import "strconv"
import "path"

/*func (tile Tile) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if tile.Probability != 0 {
		data["probability"] = tile.Probability
	}
	if tile.Terrain != "" {
		data["terrain"] = tile.Probability

		t.TilesJSON[strconv.FormatInt(tile.Id, 10)] = tile
		tile.Image = tile.Image
		arr := strings.Split(tile.Terrain, ",")
		for _, s := range arr {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				i = -1
			}
			tile.TerrainJSON = append(tile.TerrainJSON, i)
		}
		if len(tile.TerrainJSON) != 4 {
			tile.TerrainJSON = nil
		}
	}

	/*XMLName xml.Name `xml:"tile" json:"-"`

	Id          int64   `xml:"id,attr" json:"-"`      // id: The local tile ID within its tileset.
	Terrain     string  `xml:"terrain,attr" json:"-"` // terrain: Defines the terrain type of each corner of the tile, given as comma-separated indexes in the terrain types array in the order top-left, top-right, bottom-left, bottom-right. Leaving out a value means that corner has no terrain. (optional) (since 0.9.0)
	TerrainJSON []int64 `xml:"-" json:"terrain,omitempty"`
	Probability float32 `xml:"probability,attr" json:"probability,omitempty"` //probability: A percentage indicating the probability that this tile is chosen when it competes with others while editing with the terrain tool. (optional) (since 0.9.0)

	Properties Properties `xml:"properties>property" json:"properties,omitempty"`
	Image     Image      `xml:"image" json:"-"`
	* /
}*/

func (list Properties) MarshalJSON() ([]byte, error) {
	data := make(map[string]string)
	for _, p := range list {
		data[p.Name] = p.Value
	}
	return json.Marshal(data)
}

func (set Tileset) MarshalJSON() ([]byte, error) {
	dir := ""
	if set.Source != "" {
		ts, err := LoadTSX(set.Source)
		if err != nil {
			return nil, err
		}
		log.Println("LoadTSX:", ts)
		dir = path.Dir(set.Source)
		ts.FirstGID = set.FirstGID
		set = *ts
	}
	data := map[string]interface{}{
		"firstgid":    set.FirstGID,
		"name":        set.Name,
		"tilewidth":   set.TileWidth,
		"tileheight":  set.TileHeight,
		"spacing":     set.Spacing,
		"margin":      set.Margin,
		"tileoffset":  set.TileOffset,
		"properties":  set.Properties,
		"terrains":    set.TerrainTypes,
		"image":       path.Join(dir, set.IImage.ImageSource),
		"imagewidth":  set.IImage.Width,
		"imageheight": set.IImage.Height,
	}

	if data["terrains"] == nil {
		data["terrains"] = make(map[string]interface{})
	}

	tiles := make(map[string]*Tile)
	for _, tile := range set.Tiles {
		name := strconv.FormatInt(tile.Id, 10)
		tiles[name] = tile
	}
	data["tiles"] = tiles
	// TODO image

	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	return b, err
}

func (m Map) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"version":     m.Version,
		"orientation": m.Orientation,
		"width":       m.Width,
		"height":      m.Height,
		"tilewidth":   m.TileWidth,
		"tileheight":  m.TileHeight,
		"renderorder": m.RenderOrder,
		"properties":  m.Properties,
		"tilesets":    m.Tileset,
	}

	if m.BackgroundColor != "" {
		data["backgroundcolor"] = m.BackgroundColor
	}

	layers := []interface{}{}
	for _, layer := range m.Layers {
		ll := map[string]interface{}{
			"type":       layer.Type,
			"name":       layer.Name,
			"x":          layer.X,
			"y":          layer.Y,
			"width":      layer.Width,
			"height":     layer.Height,
			"opacity":    layer.Opacicy,
			"visible":    layer.Visible,
			"properties": layer.Properties,
		}
		switch layer.Type {
		case "tilelayer":
			d, err := layer.Data.Data()
			if err != nil {
				return nil, err
			}
			ll["data"] = d
		case "imagelayer":
			ll["image"] = layer.Image.ImageSource
			ll["imagewidth"] = layer.Image.Width
			ll["imageheight"] = layer.Image.Height
		case "objectgroup":
			ll["objects"] = layer.Objects
		default:
			log.Println("!!!!!!!!!!!", layer.Type)
		}
		layers = append(layers, ll)
	}
	if len(layers) != 0 {
		data["layers"] = layers
	}

	b, err := json.Marshal(data)
	log.Println(err)
	return b, err
}
