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

var newTimeStr = "14:00/15:00/16:00/17:00/18:00"
var govTimeStr = "16:30/17:00/17:30/18:30/19:30"

//10:00/11:00/14:00
var newAmount1 = "16000"
var newAmount2 = "16000"

var (
	YiBaoJuAcc       = "258508672821"
	SecondNormalAcc  = "259808672914"
	SecondSpecialAcc = "261108672684"
	ThirdNormalAcc   = "262408672802"
	ThirdSpecialAcc  = "250708672800"
	ForthNormalAcc   = "248108672922"
	ForthSpecialAcc  = "246808672716"
)

// b2e0231 查询归集历史记录
func TestNewSetFillUpsCon1(t *testing.T) {
	// 全部归集
	//fmt.Println("222::fillUp----------------------------------------------")
	//fillUp(SecondNormalAcc, SecondSpecialAcc, newTimeStr, newAmount1)
	//fmt.Println("333:fillUp----------------------------------------------")
	//fillUp(ThirdNormalAcc, ThirdSpecialAcc, newTimeStr, newAmount1)
	//fmt.Println("444:fillUp----------------------------------------------")
	fillUp(ForthNormalAcc, ForthSpecialAcc, newTimeStr, "12000")
	// 全部填平 - 医院
	//fmt.Println("222::GiveOut----------------------------------------------")
	//GiveOut(SecondNormalAcc, SecondSpecialAcc, newTimeStr, newAmount1)
	//fmt.Println("333::GiveOut----------------------------------------------")
	//GiveOut(ThirdNormalAcc, ThirdSpecialAcc, newTimeStr, newAmount1)
	//fmt.Println("444::GiveOut----------------------------------------------")
	//GiveOut(ForthNormalAcc, ForthSpecialAcc, newTimeStr, "18000")
	// 全部填平 - 医保局
	//fmt.Println("555::GiveOut----------------------------------------------")
	//GiveOut(YiBaoJuAcc, SecondSpecialAcc, govTimeStr, newAmount2)
	//fmt.Println("555::GiveOut----------------------------------------------")
	//GiveOut(YiBaoJuAcc, ThirdSpecialAcc, govTimeStr, newAmount2)
	//fmt.Println("555::GiveOut----------------------------------------------")
	//GiveOut(YiBaoJuAcc, ForthSpecialAcc, govTimeStr, "18000")
}

func TestNewGetBalancesCon1(t *testing.T) {
	fmt.Println("医保局::")
	GetBalance(YiBaoJuAcc)
	time.Sleep(5 * time.Second)
	fmt.Println("二医院 普户::")
	GetBalance(SecondNormalAcc)
	time.Sleep(5 * time.Second)
	fmt.Println("二医院 专户::")
	GetBalance(SecondSpecialAcc)
	time.Sleep(5 * time.Second)
	fmt.Println("三医院 普户::")
	GetBalance(ThirdNormalAcc)
	time.Sleep(5 * time.Second)
	fmt.Println("三医院 专户::")
	GetBalance(ThirdSpecialAcc)
	time.Sleep(5 * time.Second)
	fmt.Println("四医院 普户::")
	GetBalance(ForthNormalAcc)
	time.Sleep(5 * time.Second)
	fmt.Println("四医院 专户::")
	GetBalance(ForthSpecialAcc)
}

// 交易记录查询
func TestAccountTransactionsQuery(t *testing.T) {
	head := mock.MockXMLHead("b2e0035")
	body := &Modules.TrnB2e0035Rq{
		B2e0035Rq: &Modules.B2e0035Rq{
			Ibknum: "",
			Actacn: "249408672775",
			Type:   "2001",
			// 2001 查当天; 2002 查自定义; 2005查前一天
			Datescope: &Modules.Datescope0035{
				From: "20210322",
				To:   "20210322",
			},
			Amountscope: &Modules.Amountscope0035{
				From: "1",
				To:   "100000",
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

// 归集
// 专户 归集到 普户
// 普户, 专户, 时间段, 金额
// normalAcc  负责收钱的
// specialAcc 需要归集的账户
// timesStr 动作时间端
// amountStr目标金额 元
func fillUp(normalAcc, specialAcc, timesStr, amountStr string) {
	head := mock.MockXMLHead("b2e0232")
	body := &Modules.TrnB2e0232Rq{
		B2e0232Rq: &Modules.B2e0232Rq{
			Account: Modules.Account0232{
				Actacn: specialAcc,
			},
			Supactn: Modules.Supactn0232{
				Supbkno: "003",
				Supibkn: "00138",
				Actacn:  normalAcc,
			},
			Trntyp:       "00",
			Ruleid:       "001",
			Coltypset:    "100", // 归集类型设置: 差额部分按比例, 百分比
			Balance:      amountStr,
			Time:         timesStr,
			Targetamount: amountStr,
		},
	}
	rsp, err := bocBankSDK.FillUp(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

// 填平 下拨 - 补平
// 普户 给钱 给专户
// 普户, 专户, 时间段, 金额
// normalAcc 负责出钱填平
// specialAcc需要填平的
// timesStr 动作时间端
// amountStr目标金额 元
func GiveOut(normalAcc, specialAcc, timesStr, amountStr string) {
	head := mock.MockXMLHead("b2e0232")
	body := &Modules.TrnB2e0232Rq{
		B2e0232Rq: &Modules.B2e0232Rq{
			Account: Modules.Account0232{
				Actacn: specialAcc,
			},
			Supactn: Modules.Supactn0232{
				Supbkno: "003",
				Supibkn: "00138",
				Actacn:  normalAcc,
			},
			Trntyp:       "01",
			Ruleid:       "001",
			Coltypset:    amountStr, // 填平, 这里写填平金额
			Balance:      amountStr,
			Time:         timesStr,
			Targetamount: "",
		},
	}
	rsp, err := bocBankSDK.FillUp(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := xml.Marshal(rsp)
	fmt.Println(string(b))
}

func GetBalance(acc string) {
	head := mock.MockXMLHead("b2e0005")
	body := &Modules.TrnB2e0005Rq{
		B2e0005Rq: &Modules.B2e0005Rq{
			Account: &Modules.Account0005Rq{
				Ibknum: "",
				Actacn: acc,
			},
		},
	}
	rsp, err := bocBankSDK.HospitalAccountBalance(head, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("================")
	fmt.Println("acc 余额:", (*rsp.TrnB2e0005Rs.B2e0005Rs)[0].Balance.Bokbal)
	fmt.Println("================")
}
