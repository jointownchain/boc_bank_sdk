package testExamples

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	bocBankSDK "github.com/jointownchain/boc_bank_sdk"
	"github.com/jointownchain/boc_bank_sdk/Modules"
	"github.com/jointownchain/boc_bank_sdk/mock"
)

// 发起保理 <- 网银客户号是医保局的
// 16005011195
// 查询保理是否成功
//nolint:funlen
func TestCompanyFactoring(t *testing.T) {
	head := mock.MockFactoring0312XMLHead("b2e0477")
	body := &Modules.TrnB2e0477Rq{
		B2e0477Rq: &Modules.B2e0477Rq{
			Acrcinvtrf: &Modules.Acrcinvtrf0477{
				Relatedref: "1003",
				//Channelno:        "1", // 供应链渠道代码需要为空
				Unitcode: "00001",                // 00001 测试机构代码； 11611 固定死的， 生产环境的
				Finctype: "10",                   // 产品协议类型， 固定10
				Agreeref: "ACSF000012100000047A", // 协议编号， 固定为 ACSF000012100000047A
				//Agreeref:    "ACSF000012100000047A", // 协议编号， 固定为 2
				//Applid:      "396131933", // 核心客户号 - 银行才有， 使用组织机构代码， 社会统一信用代码 第九位 往后数8位 即为 组织机构代码
				Applorganid: "91410902",
				//Applerpid:        "396002125", // 申请人id为空
				Custremark:   "东森申请保理",
				Applyfincccy: "CNY",
				//Applyfincpercent: "",
				Applyfincamt: "6.3", // 金额
				//Guarcoid:         "234", // 保险需要为空
				//Applfincacno:     "255908672987", //
				//Applfincacno:     "263708672973", // 一般户
				Applfincacno: "246808672749", // 专户
			},
			Circlenum: "1",
			Invoice: &[]Modules.Invoice0477{{
				//Sellerid: "396131933", // 物流 卖方还款 核心客户号
				Sellerorganid: "91410902", // 组织机构代码， 需要银行进行维护
				//Sellererpid:    "濮阳海王医药有限公司",
				//Buyerid: "396002113", // 中医院
				Buyerorganid: "91410001", // 组织机构代码， 需要银行进行维护
				//Buyererpid:     "濮阳市中医院",
				//Buyernm:        "濮阳市中医院",
				Contractno: "123",
				Invno:      "0410019001052",
				Invccy:     "CNY",
				Invamt:     "2000.13",
				Invdate:    "20210701",
				Invvaldate: "20210723",
				Invduedate: "20210830",
				M1:         time.Now().Format("04:05"),
				M2:         "m2",
			}},
		},
	}
	rsp, err := bocBankSDK.CompanyFactoring(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("=============================")
	fmt.Println((*rsp.TrnB2e0477Rs.B2e0477Rs)[0].Obssid)
	fmt.Println("=============================")
	b, _ := xml.MarshalIndent(rsp, "\t", "\t")
	fmt.Println(string(b))
	// 查询保理结果 b2e0483
	// 查询保理是否成功
	head = mock.MockFactoring0312XMLHead("b2e0483")
	resultBody := &Modules.TrnB2e0483Rq{
		B2e0483Rq: &Modules.B2e0483Rq{
			Relatedref: "1003",
			Obssid:     (*rsp.TrnB2e0477Rs.B2e0477Rs)[0].Obssid,
		},
	}
	result, err := bocBankSDK.FactoringResult(head, resultBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	ab, _ := xml.MarshalIndent(result, "\t", "\t")
	fmt.Println(string(ab))
}

// 查看是否保理成功
func TestFactoringStatus(t *testing.T) {
	// 查询保理结果 b2e0483
	// 查询保理是否成功
	head := mock.MockFactoring0312XMLHead("b2e0483")
	query := func(obssid string) {
		resultBody := &Modules.TrnB2e0483Rq{
			B2e0483Rq: &Modules.B2e0483Rq{
				Relatedref: "1003",
				//Obssid:     "16005645806",
				//Obssid:     "16005645716",
				// 16005645837
				Obssid: obssid,
			},
		}
		result, err := bocBankSDK.FactoringResult(head, resultBody)
		if err != nil {
			fmt.Println(err)
			return
		}
		ab, _ := xml.MarshalIndent(result, "\t", "\t")
		fmt.Println(string(ab))
	}
	// for i := 16005645716;i<16005645899;i++{
	query("16005649484")
	// }
}

// TestCompanyFactoringRepayment 保理还款
func TestCompanyFactoringRepayment(t *testing.T) {
	head := mock.MockFactoring0312XMLHead("b2e0481")
	body := &Modules.TrnB2e0481Rq{
		Ceitinfo: "",
		B2e0481Rq: &Modules.B2e0481Rq{
			Acrcinvtrf: &Modules.Acrcinvtrf0481{
				Relatedref: "1003",
				//Channelno:   "1",
				Unitcode: "00001",
				Finctype: "10",
				//Applid:      "299957633",
				Applorganid: "91410902",
				//Applerpid:   "2345678",
				Custremark: "东森偿还保理",
				Repayacno:  "246808672749", // 账号
			},
			Circlenum: "1",
			Invoice: &[]Modules.Invoice0481{{
				//Invref:         "acc123",
				Invno: "0410019001052",
				//Sellerid:       "Delivery001",
				Sellerorganid: "91410902",
				//Sellererpid:    "物流",
			}},
		},
	}
	rsp, err := bocBankSDK.CompanyFactoringRepayment(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.MarshalIndent(rsp, "\t", "\t")
	fmt.Println(string(b))

}

// 保理还款通知
func TestCompanyFactoringRepaymentNotice(t *testing.T) {
	head := mock.MockFactoring0312XMLHead("b2e0486")
	body := &Modules.TrnB2e0486Rq{
		B2e0486Rq: &Modules.B2e0486Rq{
			Obssid:     "16005649996",
			Relatedref: "1003",
		},
	}
	rsp, err := bocBankSDK.CompanyFactoringRepaymentNotice(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.MarshalIndent(rsp, "\t", "\t")
	fmt.Println(string(b))
}

// 保理还款结果查询
func TestCompanyFactoringRepaymentResult(t *testing.T) {
	head := mock.MockFactoring0312XMLHead("b2e0487")
	body := &Modules.TrnB2e0487Rq{
		B2e0487Rq: &Modules.B2e0487Rq{
			Obssid:     "16005649991",
			Relatedref: "1003",
		},
	}
	rsp, err := bocBankSDK.CompanyFactoringRepaymentResult(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.MarshalIndent(rsp, "\t", "\t")
	fmt.Println(string(b))
}
