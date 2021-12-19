package models


import (
	"encoding/xml"
)

type Herd_Xml struct {
	XMLName xml.Name `xml:"herd"`
	Labyak  []struct {
		Name string `xml:"name,attr"`
		Age  string `xml:"age,attr"`
		Sex  string `xml:"sex,attr"`
	} `xml:"labyak"`
}
