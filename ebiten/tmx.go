package main

import (
	"encoding/xml"
	"io"
	"strings"
	"unicode"
)

type TMX struct {
	XMLName xml.Name   `xml:"map"`
	Layers  []TMXLayer `xml:"layer"`
}

type TMXLayer struct {
	XMLName xml.Name `xml:"layer"`
	Data    string   `xml:"data"`

	cells []string
}

func (tmx *TMX) Decode(in io.Reader) error {
	reader := xml.NewDecoder(in)
	if err := reader.Decode(tmx); err != nil {
		return err
	}

	for index := range tmx.Layers {
		data := removeSpace(tmx.Layers[index].Data)
		tmx.Layers[index].cells = strings.Split(data, ",")
	}

	return nil
}

func removeSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
