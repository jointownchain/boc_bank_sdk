package testExamples

import (
	"encoding/xml"
	"fmt"
	"testing"

	bocBankSDK "github.com/jointownchain/boc_bank_sdk"
	"github.com/jointownchain/boc_bank_sdk/Modules"
	"github.com/jointownchain/boc_bank_sdk/mock"
)

// 提额
func TestRaiseRequest(t *testing.T) {
	head := mock.MockXMLHead("b2e0057")
	body := &Modules.TrnB2e0057Rq{
		B2e0057Rq: &Modules.B2e0057Rq{
			Account: &Modules.Account0057{
				Ibknum: "46243",
				Actacn: "255908672987",
			},
			Allquotas: &Modules.Allquotas0057{
				Cycletype: "01",
				Quotalist: &Modules.Quotalist0057{
					Chltyp:    "03",
					Txntyp:    "03",
					Quttyp:    "01",
					Cycletype: "Y",
					Quota:     "0",
					Accutype:  "N",
					Resettype: "Y",
					Tempquota: "100",
				},
			},
		},
	}
	rsp, err := bocBankSDK.RaiseRequest(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

// 限额查询接口
func TestTransferLimitQuery(t *testing.T) {
	head := mock.MockXMLHead("b2e0058")
	body := &Modules.TrnB2e0058Rq{
		B2e0058Rq: &Modules.B2e0058Rq{
			Account0058: &Modules.Account0058{
				Ibknum: "",
				Actacn: "255908672987",
			},
		},
	}
	rsp, err := bocBankSDK.TransferLimitQuery(head, body)
	if err != nil {
		fmt.Println(err) // RaiseRequest B2E error: code : B999, msg: 账户未开通限额管理产品(GCMS.E065)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))

}
