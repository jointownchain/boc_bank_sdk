package testExamples

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gitlab.koblitzdigital.com/lib/bocBankSDK"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/Modules"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/mock"
	"testing"
	"time"
)

// 0358 对账服务-余额对账单查询(b2e0358)
func TestBalanceStatementsQuery(t *testing.T) {
	head := mock.MockFactoringXMLHead("b2e0358")
	head.Custid = "136170402"
	body := &Modules.TrnB2e0358Rq{
		B2e0358Rq: &Modules.B2e0358Rq{
			Checkmonth:  "202105",
			Overdueflag: "",
			Conagrno:    "",
			Begnum:      "",
			Recnum:      "",
		},
	}
	rsp, err := bocBankSDK.BalanceStatementsQuery(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	j, _ := json.Marshal(rsp)
	fmt.Println(string(j))
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

// 0359 对账服务-对账单下载(b2e0359)
func TestBalanceStatementsDownload(t *testing.T) {
	head := mock.MockXMLHead("b2e0359")
	head.Custid = "133832424"
	body := &Modules.TrnB2e0359Rq{
		B2e0359Rq: &Modules.B2e0359Rq{
			Insid:   time.Now().Format("20060102150205"),
			Stmtnum: "1",
		},
	}
	rsp, err := bocBankSDK.BalanceStatementsDownload(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	j, _ := json.Marshal(rsp)
	fmt.Println(string(j))
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

// 0360 对账服务-账单核对结果反馈(b2e0360)
func TestBalanceStatementsFeedback(t *testing.T) {
	head := mock.MockXMLHead("b2e0360")
	head.Custid = "133832424"
	body := &Modules.TrnB2e0360Rq{
		B2e0360Rq: &Modules.B2e0360Rq{
			Stmtnum:   "1",
			Isrestore: "1",
			Insid:     time.Now().Format("20060102150205"),
			Multidata: &[]Modules.Multidata0360{{
				Confirmdata: nil,
			}},
		},
	}
	rsp, err := bocBankSDK.BalanceStatementsFeedback(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	j, _ := json.Marshal(rsp)
	fmt.Println(string(j))
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}
