package Modules

import "encoding/xml"

type Bocb2eResponseRoot struct {
	XMLName  xml.Name    `xml:"bocb2e"`
	Text     string      `xml:",chardata"`
	Version  string      `xml:"version,attr"`
	Security string      `xml:"security,attr"`
	Lang     string      `xml:"lang,attr"`
	Head     *Head       `xml:"head"`
	Trans    interface{} `xml:"trans"`
}

type Bocb2eRequestRoot struct {
	XMLName  xml.Name    `xml:"bocb2e"`
	Text     string      `xml:",chardata"`
	Version  string      `xml:"version,attr"`
	Security string      `xml:"security,attr"`
	Lang     string      `xml:"lang,attr"`
	Head     *Head       `xml:"head"`
	Trans    RequestTrans `xml:"trans"`
}

type RequestTrans struct {
	FillIn interface{}
}
