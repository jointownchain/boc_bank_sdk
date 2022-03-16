package Modules

import "encoding/xml"

type Error struct {
	XMLName  xml.Name   `xml:"bocb2e"`
	Text     string     `xml:",chardata"`
	Version  string     `xml:"version,attr"`
	Security string     `xml:"security,attr"`
	Lang     string     `xml:"lang,attr"`
	Head     ErrorHead  `xml:"head"`
	Trans    ErrorTrans `xml:"trans"`
}
type ErrorTrans struct {
	Text          string        `xml:",chardata"`
	TrnB2eerrorRs TrnB2eerrorRs `xml:"trn-b2eerror-rs"`
}

type TrnB2eerrorRs struct {
	Text   string      `xml:",chardata"`
	Status ErrorStatus `xml:"status"`
}

type ErrorStatus struct {
	Text   string `xml:",chardata"`
	Rspcod string `xml:"rspcod"`
	Rspmsg string `xml:"rspmsg"`
}

type ErrorHead struct {
	Text  string `xml:",chardata"`
	Trnid string `xml:"trnid"`
}
