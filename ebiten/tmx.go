package main

import (
	"encoding/xml"
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

	Cells []string
}

func (tmx *TMX) Open(level string) error {
	data, err := assets.ReadFile("asstes/maps/" + level + ".tmx")
	if err != nil {
		return err
	}

	if err := xml.Unmarshal(data, tmx); err != nil {
		return err
	}

	for index := range tmx.Layers {
		data := removeSpace(tmx.Layers[index].Data)
		tmx.Layers[index].Cells = strings.Split(data, ",")
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
