package main

import "fmt"

func (m *Map) Validation() error {
	if m.Version != "1.0" {
		return fmt.Errorf("fail version `%s`", m.Version)
	}
	switch m.Orientation {
	case "orthogonal", "isometric", "staggered":
		//pass
	default:
		return fmt.Errorf("fail orientation `%s`", m.Orientation)
	}
	/*case m.Orientation != "orthogonal":
	case m.RenderOrder != "left-up":
		return fmt.Errorf("fail support only left-down renderorder `$s`", m.RenderOrder)
	*/
	return nil
}
