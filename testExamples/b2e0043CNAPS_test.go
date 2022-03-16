package testExamples

import (
	"encoding/xml"
	"fmt"
	"gitlab.koblitzdigital.com/lib/bocBankSDK"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/Modules"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/mock"
	"testing"
)

// b2e0043 查询联行号
func TestCnBankCode(t *testing.T) {
	head := mock.MockXMLHead("b2e0043")
	body := &Modules.TrnB2e0043Rq{
		B2e0043Rq: &Modules.B2e0043Rq{Tasktpy: "0"},
		// 0：CNAPS文件	   //1：中行机构号文件    //2：电子汇票行号文件
	}
	rsp, err := bocBankSDK.CNAPSQuery(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}
