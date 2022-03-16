package bocBankSDK

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/Modules"
	"io/ioutil"
	"net/http"
	"time"
)

// 向银行发送请求
// xmls marshal 过的 xml 文件
// url 银行 url 传参
// opts 选项, 规定是哪种类型的请求
func doRequest(xmls []byte, opts string) (interface{}, error) {

	method := "POST"

	client := new(http.Client)
	client.Timeout = time.Second * 25

	bankUrl, err := GetBankUrl()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, bankUrl, bytes.NewReader(xmls))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/xml")

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http error")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(body))

	return parseResponseXML(body, opts)
}

// NewXMLRequestHead 构造 统一 request XML bytes 头 <head></head>
func NewXMLRequestHead(termid, trnid, custid, cusopr, trncod, token string) (bh *Modules.Head) {
	return &Modules.Head{
		Termid: termid,
		Trnid:  trnid,
		Custid: custid,
		Cusopr: cusopr,
		Trncod: trncod,
		Token:  token,
	}
}

//<head>
//<termid>E127000000001</termid>
//<trnid>20060704001</trnid>
//<custid>12345678</custid>
//<cusopr>BOC</cusopr>
//<trncod>b2e0003</trncod>
//<token>9TTQALYGH1</token>
//</head>

// 构造请求 xml request
func NewXMLRequest(bh *Modules.Head, trans interface{}) ([]byte, error) {
	if bh.Trnid == "" {
		return nil, errors.New("Modules.Head is nil")
	}
	if trans == nil {
		return nil, errors.New("trans is nil")
	}
	reqXML := &Modules.Bocb2eRequestRoot{
		Version:  "120",
		Security: "true",
		Lang:     "zh_CN",
		Head:     bh,
		Trans:    Modules.RequestTrans{FillIn: trans},
	}
	XMLBytes, err := xml.Marshal(reqXML)
	if err != nil {
		return nil, errors.Wrap(err, "NewXMLRequest xml.Marshal")
	}
	fmt.Println("Request:\n", string(XMLBytes), "\nRequest done")
	XMLBytes = []byte(xml.Header + string(XMLBytes))
	return XMLBytes, nil
}
