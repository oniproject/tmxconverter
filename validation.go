package main

import "fmt"

func (m *Map) Validation() error {
	switch {
	case m.Version != "1.0":
		return fmt.Errorf("fail version `%s`", m.Version)
	case m.Orientation != "orthogonal":
		return fmt.Errorf("fail support only orthogonal orientation `%s`", m.Orientation)
	case m.RenderOrder != "left-up":
		return fmt.Errorf("fail support only left-down renderorder `$s`", m.RenderOrder)
	}
	return nil
}
