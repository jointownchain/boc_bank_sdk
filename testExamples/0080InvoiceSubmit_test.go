package testExamples

import (
	"encoding/xml"
	"fmt"
	"gitlab.koblitzdigital.com/lib/bocBankSDK"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/Modules"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/mock"
	"testing"
)

func TestInvoiceSubmit(t *testing.T) {
	head := mock.MockXMLHead("b2e0080")
	body := &Modules.TrnB2e0080Rq{
		B2e0080Rq: &Modules.B2e0080Rq{
			Acrcinvtrf: &Modules.Acrcinvtrf0080{
				Relatedref:        "111",
				Finctype:          "18",
				Erpid:             "11",
				Factorcode:        "222",
				Guarcoid:          "1",
				Applyflg:          "1",
				Directpurchaseflg: "1",
				Directpurchaseccy: "CNY",
				Directpurchaseamt: "1",
				Branchid:          "41111",
				Applid:            "9999999",
				Authpoint:         "",
				Custremark:        "123abc",
			},
			Circlenum: "1",
			Invoice: &Modules.Invoice0080{
				Buyerid:     "12345",
				Buyererpid:  "555",
				Sellerid:    "",
				Sellererpid: "666",
				Contractno:  "1",
				Invno:       "fapiaohao111",
				Invccy:      "CNY",
				Invamt:      "123.33",
				Invdate:     "20111228",
				Invvaldate:  "20111229",
				Invduedate:  "20111230",
			},
		},
	}
	rsp, err := bocBankSDK.InvoiceSubmit(head, body)
	if err != nil {
		fmt.Println(err) // InvoiceSubmit TRN error: code : B077, msg: 操作员未开通此功能权限,请修改
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}
