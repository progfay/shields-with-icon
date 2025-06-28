package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// [Brands] type of [simple-icons/simple-icons]
//
// [simple-icons/simple-icons]: https://github.com/simple-icons/simple-icons
// [Brands]: https://github.com/simple-icons/simple-icons/blob/0ee651fd4dc50a48c35a96dc48a21984906a84e5/.jsonschema.json#L4
type Icon struct {
	Title  string `json:"title"`
	Hex    string `json:"hex"`
	Source string `json:"source"`
}

// ref. https://pkg.go.dev/encoding/json#example-Decoder.Decode-Stream
func decodeIcons(r io.Reader) ([]Icon, error) {
	// [simple-icons] type of [simple-icons/simple-icons]
	//
	// [simple-icons/simple-icons]: https://github.com/simple-icons/simple-icons
	// [simple-icons]:  https://github.com/simple-icons/simple-icons/blob/0ee651fd4dc50a48c35a96dc48a21984906a84e5/.jsonschema.json
	var icons []Icon
	dec := json.NewDecoder(r)

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		return nil, err
	}
	if t, ok := t.(json.Delim); !ok || t != '[' {
		return nil, fmt.Errorf("first token is not '[': %T: %v", t, t)
	}

	for dec.More() {
		var i Icon
		err := dec.Decode(&i)
		if err != nil {
			return nil, err
		}

		icons = append(icons, i)
	}

	t, err = dec.Token()
	if err != nil {
		return nil, err
	}
	if t, ok := t.(json.Delim); !ok || t != ']' {
		return nil, fmt.Errorf("last token is not ']': %T: %v", t, t)
	}

	return icons, nil
}

func getIcons() ([]Icon, error) {
	res, err := http.DefaultClient.Get("https://raw.githubusercontent.com/simple-icons/simple-icons/develop/data/simple-icons.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	defer io.Copy(io.Discard, res.Body)

	return decodeIcons(res.Body)
}
