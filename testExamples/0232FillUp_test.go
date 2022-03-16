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

var timeStr = "16:00/16:30/17:00/17:30/18:00"

//10:00/11:00/14:00
var amount1 = "4000"
var amount2 = "6000"

func TestSetFillUps(t *testing.T) {
	TestFillUp(t)  // 4000
	TestFillUp2(t) // 4000
	TestFillUp3(t) // 6000
	TestFillUp4(t) // 6000
}

func TestGetBalances(t *testing.T) {
	TestHospitalAccountBalance1(t) // 5000
	time.Sleep(5 * time.Second)
	TestHospitalAccountBalance2(t) // 5000
	time.Sleep(5 * time.Second)
	TestHospitalAccountBalance3(t) // 8000
	time.Sleep(5 * time.Second)
	TestHospitalAccountBalance4(t) // 8000
}

// b2e0232 归集
func TestFillUp(t *testing.T) {
	head := mock.MockXMLHead("b2e0232")
	body := &Modules.TrnB2e0232Rq{
		B2e0232Rq: &Modules.B2e0232Rq{
			Account: Modules.Account0232{
				Actacn: "249408672775",
			},
			Supactn: Modules.Supactn0232{
				Supbkno: "003",
				Supibkn: "00138",
				Actacn:  "248108672900",
			},
			Trntyp:       "00",
			Ruleid:       "001",
			Coltypset:    "100", // 归集类型设置: 差额部分按比例, 百分比
			Balance:      amount1,
			Time:         timeStr,
			Targetamount: amount1,
		},
	}
	rsp, err := bocBankSDK.FillUp(head, body)
	if err != nil {
		fmt.Println(err) // RaiseRequest B2E error: code : B999, msg: 账户未开通限额管理产品(GCMS.E065)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

// 下拨 - 补平
func TestFillUp2(t *testing.T) {
	head := mock.MockXMLHead("b2e0232")
	body := &Modules.TrnB2e0232Rq{
		B2e0232Rq: &Modules.B2e0232Rq{
			Account: Modules.Account0232{
				Actacn: "258508672832",
			},
			Supactn: Modules.Supactn0232{
				Supbkno: "003",
				Supibkn: "00138",
				Actacn:  "262408672788",
			},
			Trntyp:       "01",
			Ruleid:       "001",
			Coltypset:    amount1,
			Balance:      amount1,
			Time:         timeStr,
			Targetamount: "",
		},
	}
	rsp, err := bocBankSDK.FillUp(head, body)
	if err != nil {
		fmt.Println(err) // RaiseRequest B2E error: code : B999, msg: 账户未开通限额管理产品(GCMS.E065)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

// b2e0232 归集
func TestFillUp3(t *testing.T) {
	head := mock.MockXMLHead("b2e0232")
	body := &Modules.TrnB2e0232Rq{
		B2e0232Rq: &Modules.B2e0232Rq{
			Account: Modules.Account0232{
				Actacn: "258508672832",
			},
			Supactn: Modules.Supactn0232{
				Supbkno: "003",
				Supibkn: "00138",
				Actacn:  "257208672853",
			},
			Trntyp:       "00",
			Ruleid:       "001",
			Coltypset:    "100",
			Balance:      amount2,
			Time:         timeStr,
			Targetamount: amount2,
		},
	}
	rsp, err := bocBankSDK.FillUp(head, body)
	if err != nil {
		fmt.Println(err) // RaiseRequest B2E error: code : B999, msg: 账户未开通限额管理产品(GCMS.E065)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

// 补平
func TestFillUp4(t *testing.T) {
	head := mock.MockXMLHead("b2e0232")
	body := &Modules.TrnB2e0232Rq{
		B2e0232Rq: &Modules.B2e0232Rq{
			Account: Modules.Account0232{
				Actacn: "258508672832",
			},
			Supactn: Modules.Supactn0232{
				Supbkno: "003",
				Supibkn: "00138",
				Actacn:  "257208672853",
			},
			Trntyp:       "01",
			Ruleid:       "001",
			Coltypset:    amount2,
			Balance:      amount2,
			Time:         timeStr,
			Targetamount: "",
		},
	}
	rsp, err := bocBankSDK.FillUp(head, body)
	if err != nil {
		fmt.Println(err) // RaiseRequest B2E error: code : B999, msg: 账户未开通限额管理产品(GCMS.E065)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

func TestHospitalAccountBalance(t *testing.T) {
	head := mock.MockXMLHead("b2e0005")
	body := &Modules.TrnB2e0005Rq{
		B2e0005Rq: &Modules.B2e0005Rq{
			Account: &Modules.Account0005Rq{
				Ibknum: "",
				Actacn: "258508672832",
			},
		},
	}
	rsp, err := bocBankSDK.HospitalAccountBalance(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

// 2021年03月15日10:49:25  11点 归集 填平 100 测试
// 中医院一般户 248108672900    40080.00
func TestHospitalAccountBalance1(t *testing.T) {
	head := mock.MockXMLHead("b2e0005")
	body := &Modules.TrnB2e0005Rq{
		B2e0005Rq: &Modules.B2e0005Rq{
			Account: &Modules.Account0005Rq{
				Ibknum: "",
				Actacn: "248108672900",
			},
		},
	}
	rsp, err := bocBankSDK.HospitalAccountBalance(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("================")
	fmt.Println("中医院一般户 248108672900 余额:", (*rsp.TrnB2e0005Rs.B2e0005Rs)[0].Balance.Bokbal)
	fmt.Println("================")
}

// 中医院专户 249408672775    1.00
func TestHospitalAccountBalance2(t *testing.T) {
	head := mock.MockXMLHead("b2e0005")
	body := &Modules.TrnB2e0005Rq{
		B2e0005Rq: &Modules.B2e0005Rq{
			Account: &Modules.Account0005Rq{
				Ibknum: "",
				Actacn: "249408672775",
			},
		},
	}
	rsp, err := bocBankSDK.HospitalAccountBalance(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("================")
	fmt.Println("中医院 专户 249408672775 余额:", (*rsp.TrnB2e0005Rs.B2e0005Rs)[0].Balance.Bokbal)
	fmt.Println("================")
}

// 人民医院 专户 258508672832   10000.00
func TestHospitalAccountBalance4(t *testing.T) {
	head := mock.MockXMLHead("b2e0005")
	body := &Modules.TrnB2e0005Rq{
		B2e0005Rq: &Modules.B2e0005Rq{
			Account: &Modules.Account0005Rq{
				Ibknum: "",
				Actacn: "258508672832",
			},
		},
	}
	rsp, err := bocBankSDK.HospitalAccountBalance(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("================")
	fmt.Println("人民医院 专户 258508672832 余额:", (*rsp.TrnB2e0005Rs.B2e0005Rs)[0].Balance.Bokbal)
	fmt.Println("================")
}

// 人民医院 一般户 257208672853  19001.00
func TestHospitalAccountBalance3(t *testing.T) {
	head := mock.MockXMLHead("b2e0005")
	body := &Modules.TrnB2e0005Rq{
		B2e0005Rq: &Modules.B2e0005Rq{
			Account: &Modules.Account0005Rq{
				Ibknum: "",
				Actacn: "257208672853",
			},
		},
	}
	rsp, err := bocBankSDK.HospitalAccountBalance(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("================")
	fmt.Println("人民医院 一般户 257208672853 余额:", (*rsp.TrnB2e0005Rs.B2e0005Rs)[0].Balance.Bokbal)
	fmt.Println("================")
}
