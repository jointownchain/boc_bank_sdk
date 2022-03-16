package testExamples

import (
	"encoding/xml"
	"fmt"
	"gitlab.koblitzdigital.com/lib/bocBankSDK"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/Modules"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/mock"
	"testing"
	"time"
)

// 2021年03月15日10:45:56 396002108_20210315_7802_10016.zip
// todo b2e0577
// todo b2e0580 定期存款的利息账户设置 ？ 不需要
// b2e0500 回单下载
func TestReceiptDownload(t *testing.T) {
	head := mock.MockXMLHead("b2e0500")
	//head := mock.MockFactoringXMLHead("b2e0500")
	body := &Modules.TrnB2e0500Rq{
		B2e0500Rq: &Modules.B2e0500Rq{
			Tratype: "A",
			Amtscope: &Modules.Amtscope0500{
				Minamt: "10",
				Maxamt: "80000",
			},
			Datescope: &Modules.Datescope0500{
				From: "20210625",
				To:   "20210625",
			},
			Domaccount: &Modules.Domaccount0500{
				Transact:  "254608672703",
				Transname: "",
			},
			Insid: time.Now().Format("20060102150405"),
			Multidata: &Modules.Multidata0500{
				Confirmdata: &[]Modules.Confirmdata0500{{
					Actacn: "255908672987",
					Ibknum: "",
				}},
			},
		},
	}
	rsp, err := bocBankSDK.ReceiptDownload(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}
