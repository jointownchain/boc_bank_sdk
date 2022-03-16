package testExamples

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	bocBankSDK "github.com/jointownchain/boc_bank_sdk"
	"github.com/jointownchain/boc_bank_sdk/Modules"
	"github.com/jointownchain/boc_bank_sdk/mock"
)

// 1. 构造 xml 请求头
// NewXMLRequestHead
// 2. 构造 xml 请求体
// new(TrnB2e0009Rq)/或者别的
// 3. 调用 HospitalPayment/或者别的 方法
func TestHospitalPayment(t *testing.T) {
	head := mock.MockXMLHead("b2e0009")
	body := &Modules.TrnB2e0009Rq{
		Ceitinfo:  "",
		Transtype: "", // 1、 委托待授权 2、 授权退回修改该项可空，表示普通转账汇划交易；非空时只能为1或2
		B2e0009Rq: &Modules.B2e0009Rq{
			Insid:  "1234561120", // 每次都应该不一样
			Obssid: "",
			Fractn: &Modules.Fractn0009{
				Fribkn: "",
				Actacn: "249408672775",
				Actnam: "濮阳市中医院",
			},
			Toactn: &Modules.Toactn0009{
				Toibkn: "44951",
				Actacn: "255908672987",
				Toname: "濮阳海王医药有限公司",
				Toaddr: "",
				Tobknm: "",
			},
			Trnamt:  "10000",
			Trncur:  "001",
			Priolv:  "1",
			Furinfo: "第0002号合同货款，请查收",
			Trfdate: "20210615",
			Trftime: "",
			Comacn:  "249408672775",
		},
	}
	mock.MockSetBankURL()
	rsp, err := bocBankSDK.HospitalPayment(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

func TestHospitalPayment1(t *testing.T) {
	head := mock.MockXMLHead("b2e0009")
	body := &Modules.TrnB2e0009Rq{
		Ceitinfo:  "",
		Transtype: "", // 1、 委托待授权 2、 授权退回修改该项可空，表示普通转账汇划交易；非空时只能为1或2
		B2e0009Rq: &Modules.B2e0009Rq{
			Insid:  "1234561191", // 每次都应该不一样
			Obssid: "",           // 16005010586
			Fractn: &Modules.Fractn0009{
				Fribkn: "",
				Actacn: "255908672987",
				Actnam: "濮阳海王医药有限公司",
			},
			Toactn: &Modules.Toactn0009{
				Toibkn: "44951",
				Actacn: "249408672775",
				Toname: "濮阳市中医院",
				Toaddr: "",
				Tobknm: "",
			},
			Trnamt:  "1",
			Trncur:  "001",
			Priolv:  "1",
			Furinfo: "企业往医药转",
			Trfdate: "20210312",
			Trftime: "",
			Comacn:  "255908672987",
		},
	}
	mock.MockSetBankURL()
	rsp, err := bocBankSDK.HospitalPayment(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

func TestHospitalPayment3(t *testing.T) {
	head := mock.MockXMLHead("b2e0009")
	body := &Modules.TrnB2e0009Rq{
		Ceitinfo:  "",
		Transtype: "", // 1、 委托待授权 2、 授权退回修改该项可空，表示普通转账汇划交易；非空时只能为1或2
		B2e0009Rq: &Modules.B2e0009Rq{
			Insid:  "1234561191", // 每次都应该不一样
			Obssid: "",           // 16005010586
			Fractn: &Modules.Fractn0009{
				Fribkn: "",
				Actacn: "258508672832",
				Actnam: "濮阳市人民医院",
			},
			Toactn: &Modules.Toactn0009{
				Toibkn: "44951",
				Actacn: "255908672987",
				Toname: "濮阳海王医药有限公司",
				Toaddr: "",
				Tobknm: "",
			},
			Trnamt:  "1",
			Trncur:  "001",
			Priolv:  "1",
			Furinfo: "医药转企业",
			Trfdate: "20210312",
			Trftime: "",
			Comacn:  "255908672987",
		},
	}
	mock.MockSetBankURL()
	rsp, err := bocBankSDK.HospitalPayment(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

// 中医院一般户 248108672900    40080.00
func TestHospitalAccountBalanceII(t *testing.T) {
	// head := mock.MockXMLHead("b2e0005")
	head := mock.MockFactoring0312XMLHead("b2e0005")
	body := &Modules.TrnB2e0005Rq{
		B2e0005Rq: &Modules.B2e0005Rq{
			Account: &Modules.Account0005Rq{
				Ibknum: "46243",
				//Actacn: "249408672775",
				Actacn: "246808672749",
			},
		},
	}
	rsp, err := bocBankSDK.HospitalAccountBalance(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.MarshalIndent(rsp, "\t", "\t")
	fmt.Println(string(b))
}

// TODO 测试外行转账

func TestPaymentsQuery(t *testing.T) {
	head := mock.MockXMLHead("b2e0007")
	body := &Modules.TrnB2e0007Rq{
		B2e0007Rq: &[]Modules.B2e0007Rq{{
			Insid:  "1234561191",
			Obssid: "16005010586", // 银行返回 ID
		},
		},
	}
	rsp, err := bocBankSDK.PaymentsQuery(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	j, _ := json.Marshal(rsp)
	fmt.Println(string(j))
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

func query(insid, obssid string) {
	head := mock.MockXMLHead("b2e0007")
	body := &Modules.TrnB2e0007Rq{
		B2e0007Rq: &[]Modules.B2e0007Rq{{
			Insid:  insid,
			Obssid: obssid, // 银行返回 ID
		},
		},
	}
	rsp, err := bocBankSDK.PaymentsQuery(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	j, _ := json.Marshal(rsp)
	fmt.Println(string(j))
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

func TestManQuery(t *testing.T) {
	for {
		query("210624135847", "16005610339")
		time.Sleep(time.Second * 5)
	}
}

// 转账后测试是否成功
func TestPaymentsAndQuery(t *testing.T) {

	insid := time.Now().Format("20060102150405")

	head := mock.MockXMLHead("b2e0009")
	body := &Modules.TrnB2e0009Rq{
		Ceitinfo:  "",
		Transtype: "", // 1、 委托待授权 2、 授权退回修改该项可空，表示普通转账汇划交易；非空时只能为1或2
		B2e0009Rq: &Modules.B2e0009Rq{
			Insid:  insid, // 每次都应该不一样
			Obssid: "",
			Fractn: &Modules.Fractn0009{
				Fribkn: "",
				Actacn: "254608672703",
				//Actnam: "濮阳市医保局",
			},
			//Fractn: &Modules.Fractn0009{
			//	Fribkn: "",
			//	Actacn: "248108672900",
			//	Actnam: "濮阳市中医院",
			//},
			Toactn: &Modules.Toactn0009{
				Toibkn: "44951",
				Actacn: "249408672775",
				Toname: "濮阳市中医院",
				Toaddr: "",
				Tobknm: "",
			},
			//Toactn: &Modules.Toactn0009{
			//	Toibkn: "44951",
			//	Actacn: "254608672703",
			//	Toname: "濮阳海王医药有限公司",
			//	Toaddr: "",
			//	Tobknm: "",
			//},
			Trnamt:  "5000",
			Trncur:  "001",
			Priolv:  "0",
			Furinfo: "医保局往医院转",
			Trfdate: time.Now().Format("20060102"),
			//Trfdate: "20210331",
			Trftime: "",
			Comacn:  "248108672900",
		},
	}
	mock.MockSetBankURL()
	rsp, err := bocBankSDK.HospitalPayment(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	obssid := (*rsp.TrnB2e0009Rs.B2e0009Rs)[0].Obssid
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
	fmt.Println()
	for i := 0; i < 100; i++ {
		query(insid, obssid)
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

// todo b2e0061 公对私转账会话
// 交易记录查询
func TestAccountTransactionsQueryII(t *testing.T) {
	head := mock.MockXMLHead("b2e0035")
	body := &Modules.TrnB2e0035Rq{
		B2e0035Rq: &Modules.B2e0035Rq{
			Ibknum: "",
			Actacn: "249408672775",
			Type:   "2002",
			// 2001 查当天; 2002 查自定义; 2005查前一天
			Datescope: &Modules.Datescope0035{
				From: "20210624",
				To:   "20210624",
			},
			Amountscope: &Modules.Amountscope0035{
				From: "0.1",
				To:   "10000",
			},
			Begnum:    "1",
			Recnum:    "10",
			Direction: "2",
		},
	}
	rsp, err := bocBankSDK.AccountTransactionsQuery(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	j, _ := json.Marshal(rsp)
	fmt.Println(string(j))
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

func TestPaymentRequest(t *testing.T) {
	head := mock.MockXMLHead("b2e0081")
	body := &Modules.TrnB2e0081Rq{
		B2e0081Rq: &Modules.B2e0081Rq{
			Acrcpayment: &Modules.Acrcpayment0081{
				Relatedref: "1111",
				Applyflg:   "1",
				Erpid:      "222",
				Branchid:   "333",
				Applid:     "444",
				Pmtamt:     "555",
				Pmtacno:    "100430218453",
				Fribkn:     "40142",
				Pmttype:    "2",
				Pmtflg:     "2",
				Custremark: "aaaaaaaa",
			},
			Circlenum: "1",
			Payment: &Modules.Payment0081{
				Invref:      "99",
				Sellerid:    "88",
				Sellererpid: "77",
				Invno:       "66",
				Pmtamt:      "55",
				Settacno:    "44",
			},
		},
	}
	rsp, err := bocBankSDK.PaymentRequest(head, body)
	if err != nil {
		fmt.Println(err) //PaymentRequest TRN error: code : B077, msg: 操作员未开通此功能权限,请修改
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}
