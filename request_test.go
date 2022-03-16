package bocBankSDK

import (
	"fmt"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/Modules"
	"testing"
)

// 构造完整 xml 请求
func TestTransStruct(t *testing.T) {

	tr := Modules.TrnB2e0009Rq{
		Transtype: "2",
		B2e0009Rq: &Modules.B2e0009Rq{
			Insid:  "123",
			Obssid: "123",
			Fractn: &Modules.Fractn0009{
				Fribkn: "123",
				Actacn: "123",
				Actnam: "123",
			},
			Toactn: &Modules.Toactn0009{
				Toibkn: "123",
				Actacn: "123",
				Toname: "123",
				Toaddr: "123",
				Tobknm: "123",
			},
			Trnamt:  "123",
			Trncur:  "123",
			Priolv:  "123",
			Furinfo: "123",
			Trfdate: "123",
			Trftime: "123",
			Comacn:  "123",
		},
	}

	// 构造请求头
	bh := NewXMLRequestHead("1", "2", "3", "4", "5", "6")
	// 插入XML请求体
	bb, err := NewXMLRequest(bh, tr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bb))
}
