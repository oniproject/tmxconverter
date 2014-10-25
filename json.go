package main

import "encoding/json"

func (list Properties) MarshalJSON() ([]byte, error) {
	data := make(map[string]string)

	for _, p := range list {
		data[p.Name] = p.Value
	}
	return json.Marshal(data)
}
