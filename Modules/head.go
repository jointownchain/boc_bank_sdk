package Modules

import "encoding/xml"
// xml 请求头, resp 和 request 都有, 通用
type Head struct {
	XMLName xml.Name `xml:"head"`
	Text    string   `xml:",chardata"`
	Termid  string   `xml:"termid"`
	Trnid   string   `xml:"trnid"`
	Custid  string   `xml:"custid"`
	Cusopr  string   `xml:"cusopr"`
	Trncod  string   `xml:"trncod"`
	Token   string   `xml:"token"`
}

