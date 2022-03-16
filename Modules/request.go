package Modules

import "encoding/xml"

// 各类 trans 定义
type Trans interface {
	XMLMarshal() (string, error)
}

//// 医疗机构专户付款接口 b2e0009
//const BOCB2E_OPTS_HOSPITAL_PAYMENT = "hospitalPayment"
type TrnB2e0009Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0009-rq"`
	Text      string     `xml:",chardata"`
	Ceitinfo  string     `xml:"ceitinfo"`   // 数字签名, 自动添加
	Transtype string     `xml:"transtype"`  // 交易类型
	B2e0009Rq *B2e0009Rq `xml:"b2e0009-rq"` // 转账请求内容(1~1000笔)
}

func (t *TrnB2e0009Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

// 单笔交易
type B2e0009Rq struct {
	Text    string      `xml:",chardata"`
	Insid   string      `xml:"insid"`   // 指令ID 一条转账指令在客户端的唯一标识，建议企业按时间顺序生成且不超过12位
	Obssid  string      `xml:"obssid"`  // 网银交易流水号	可空 交易类型为2时上送有用
	Fractn  *Fractn0009 `xml:"fractn"`  // 付款人账户
	Toactn  *Toactn0009 `xml:"toactn"`  // 收款人信息
	Trnamt  string      `xml:"trnamt"`  // 转账金额
	Trncur  string      `xml:"trncur"`  // 转账货币
	Priolv  string      `xml:"priolv"`  // 银行处理优先级
	Furinfo string      `xml:"furinfo"` // 用途
	Trfdate string      `xml:"trfdate"` // 要求的转账日期
	Trftime string      `xml:"trftime"` // 要求的转账时间
	Comacn  string      `xml:"comacn"`  // 手续费账号
}

// 付款人账户
type Fractn0009 struct {
	Text   string `xml:",chardata"`
	Fribkn string `xml:"fribkn"` // 联行号
	Actacn string `xml:"actacn"` // 付款账号
	Actnam string `xml:"actnam"` // 付款人名称
}

// 收款人账户
type Toactn0009 struct {
	Text   string `xml:",chardata"`
	Toibkn string `xml:"toibkn"` // 收款人联行号
	Actacn string `xml:"actacn"` // 收款账号
	Toname string `xml:"toname"` // 收款人名称
	Toaddr string `xml:"toaddr"` // 收款人地址
	Tobknm string `xml:"tobknm"` // 收款人开户行名称
}

//// 企业专户申请保理 b2e0477
//const BOCB2E_OPTS_COMPANY_FACTORING= "companyFactoring"
type TrnB2e0477Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0477-rq"`
	Text      string     `xml:",chardata"`
	Ceitinfo  string     `xml:"ceitinfo"`   // 数字签名, 自动添加
	B2e0477Rq *B2e0477Rq `xml:"b2e0477-rq"` // 报文体内容
}

func (t *TrnB2e0477Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

type B2e0477Rq struct {
	Text       string          `xml:",chardata"`
	Acrcinvtrf *Acrcinvtrf0477 `xml:"acrcinvtrf"` // 通用消息
	Circlenum  string          `xml:"circlenum"`  // 笔明细数
	Invoice    *[]Invoice0477  `xml:"invoice"`    // 笔明细内容 (1~200笔)
}

type Acrcinvtrf0477 struct {
	Text             string `xml:",chardata"`
	Relatedref       string `xml:"relatedref"`       // 业务参号
	Channelno        string `xml:"channelno"`        // 供应链渠道代码
	Unitcode         string `xml:"unitcode"`         // 机构代码
	Finctype         string `xml:"finctype"`         // 协议产品类型
	Agreeref         string `xml:"agreeref"`         // 协议编号
	Applid           string `xml:"applid"`           // 申请人BANCS客户号
	Applorganid      string `xml:"applorganid"`      // 申请人组织机构代码
	Applerpid        string `xml:"applerpid"`        // 申请人ID(对手系统)
	Custremark       string `xml:"custremark"`       // 申请人备注
	Applyfincccy     string `xml:"applyfincccy"`     // 申请融资币别
	Applyfincpercent string `xml:"applyfincpercent"` // 申请融资比例
	Applyfincamt     string `xml:"applyfincamt"`     // 申请融资金额
	Guarcoid         string `xml:"guarcoid"`         // 保险公司ID
	Applfincacno     string `xml:"applfincacno"`     // 融资入账账号
}

type Invoice0477 struct {
	Text          string `xml:",chardata"`
	Sellerid      string `xml:"sellerid"`      // 卖方BANCS客户号
	Sellerorganid string `xml:"sellerorganid"` // 卖方组织机构代码
	Sellererpid   string `xml:"sellererpid"`   // 卖方ID(对手系统)
	Buyerid       string `xml:"buyerid"`       // 买方BANCS客户号
	Buyerorganid  string `xml:"buyerorganid"`  // 买方组织机构代码
	Buyererpid    string `xml:"buyererpid"`    // 买方ID(对手系统)
	Buyernm       string `xml:"buyernm"`       // 买方名称
	Contractno    string `xml:"contractno"`    // 合同号
	Invno         string `xml:"invno"`         // 发票号
	Invccy        string `xml:"invccy"`        // 发票币别
	Invamt        string `xml:"invamt"`        // 发票金额
	Invdate       string `xml:"invdate"`       // 发票日期
	Invvaldate    string `xml:"invvaldate"`    // 发票起算日
	Invduedate    string `xml:"invduedate"`    // 发票到期日
	M1            string `xml:"m1"`            // 备注1
	M2            string `xml:"m2"`            // 备注2
}

//// 企业专户申请保理结果查询 b2e0492
//const BOCB2E_OPTS_COMPANY_FACTORING_RESULT = "companyFactoringResult"
// TODO

//// 配送企业专户偿还保理 b2e0481
//const BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT= "companyFactoringRepayment"
type TrnB2e0481Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0481-rq"`
	Text      string     `xml:",chardata"`
	Ceitinfo  string     `xml:"ceitinfo"`
	B2e0481Rq *B2e0481Rq `xml:"b2e0481-rq"`
}

func (t *TrnB2e0481Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

type B2e0481Rq struct {
	Text       string          `xml:",chardata"`
	Acrcinvtrf *Acrcinvtrf0481 `xml:"acrcinvtrf"` // 通用消息
	Circlenum  string          `xml:"circlenum"`  // 笔明细数
	Invoice    *[]Invoice0481  `xml:"invoice"`    // 笔明内容
}

type Acrcinvtrf0481 struct {
	Text        string `xml:",chardata"`
	Relatedref  string `xml:"relatedref"`  // 业务参号
	Channelno   string `xml:"channelno"`   // 供应商渠道代码
	Unitcode    string `xml:"unitcode"`    // 机构代码
	Finctype    string `xml:"finctype"`    // 协议产品类型
	Applid      string `xml:"applid"`      // 申请人BANCS客户号
	Applorganid string `xml:"applorganid"` // 申请人组织机构代码
	Applerpid   string `xml:"applerpid"`   // 申请人ID(对手系统)
	Custremark  string `xml:"custremark"`  // 备注
	Repayacno   string `xml:"repayacno"`   // 归还本息账号
}

type Invoice0481 struct {
	Text          string `xml:",chardata"`
	Invref        string `xml:"invref"`        // 发票流水号
	Invno         string `xml:"invno"`         // 发票号
	Sellerid      string `xml:"sellerid"`      // 卖方BANCS客户号
	Sellerorganid string `xml:"sellerorganid"` // 卖方组织机构代码
	Sellererpid   string `xml:"sellererpid"`   // 卖方ID (对手系统)
}

//// 保理订单偿还进度通知 b2e0486
//const BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_NOTICE = "companyFactoringRepaymentNotice"
type TrnB2e0486Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0486-rq"`
	Text      string     `xml:",chardata"`
	B2e0486Rq *B2e0486Rq `xml:"b2e0486-rq"`
}

func (t *TrnB2e0486Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

type B2e0486Rq struct {
	Text       string `xml:",chardata"`
	Obssid     string `xml:"obssid"`
	Relatedref string `xml:"relatedref"`
}

//// 保理订单偿还进度查询 b2e0487
//const BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_RESULT = "companyFactoringRepaymentResult"
type TrnB2e0487Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0487-rq"`
	Text      string     `xml:",chardata"`
	B2e0487Rq *B2e0487Rq `xml:"b2e0487-rq"`
}

func (t *TrnB2e0487Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

type B2e0487Rq struct {
	Text       string `xml:",chardata"`
	Obssid     string `xml:"obssid"`
	Relatedref string `xml:"relatedref"`
}

//// 专户交易流水查询接口 b2e0035
//const BOCB2E_OPTS_ACCOUNT_TRANSACTIONS_QUERY = "accountTransactionsQuery"
type TrnB2e0035Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0035-rq"`
	Text      string     `xml:",chardata"`
	B2e0035Rq *B2e0035Rq `xml:"b2e0035-rq"`
}

func (t *TrnB2e0035Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

type B2e0035Rq struct {
	Text        string           `xml:",chardata"`
	Ibknum      string           `xml:"ibknum"`
	Actacn      string           `xml:"actacn"`
	Type        string           `xml:"type"`
	Datescope   *Datescope0035   `xml:"datescope"`
	Amountscope *Amountscope0035 `xml:"amountscope"`
	Begnum      string           `xml:"begnum"`
	Recnum      string           `xml:"recnum"`
	Direction   string           `xml:"direction"`
}

type Datescope0035 struct {
	Text string `xml:",chardata"`
	From string `xml:"from"`
	To   string `xml:"to"`
}

type Amountscope0035 struct {
	Text string `xml:",chardata"`
	From string `xml:"from"`
	To   string `xml:"to"`
}

//// 医疗机构专户余额接口 b2e0005
//const BOCB2E_OPTS_HOSPITAL_ACCOUNT_BALANCE = "hospitalAccountBalance"
type TrnB2e0005Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0005-rq"`
	Text      string     `xml:",chardata"`
	B2e0005Rq *B2e0005Rq `xml:"b2e0005-rq"`
}

func (t *TrnB2e0005Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

type B2e0005Rq struct {
	Text    string         `xml:",chardata"`
	Account *Account0005Rq `xml:"account"`
}

type Account0005Rq struct {
	Text   string `xml:",chardata"`
	Ibknum string `xml:"ibknum"`
	Actacn string `xml:"actacn"`
}

//// 发票信息提交 b2e0080
//const BOCB2E_OPTS_INVOICE_SUBMIT = "invoiceSubmit"
type TrnB2e0080Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0080-rq"`
	Text      string     `xml:",chardata"`
	Ceitinfo  string     `xml:"ceitinfo"`
	B2e0080Rq *B2e0080Rq `xml:"b2e0080-rq"`
}

func (t *TrnB2e0080Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

type B2e0080Rq struct {
	Text       string          `xml:",chardata"`
	Acrcinvtrf *Acrcinvtrf0080 `xml:"acrcinvtrf"`
	Circlenum  string          `xml:"circlenum"`
	Invoice    *Invoice0080    `xml:"invoice"`
}

type Acrcinvtrf0080 struct {
	Text              string `xml:",chardata"`
	Relatedref        string `xml:"relatedref"`
	Finctype          string `xml:"finctype"`
	Erpid             string `xml:"erpid"`
	Factorcode        string `xml:"factorcode"`
	Guarcoid          string `xml:"guarcoid"`
	Applyflg          string `xml:"applyflg"`
	Directpurchaseflg string `xml:"directpurchaseflg"`
	Directpurchaseccy string `xml:"directpurchaseccy"`
	Directpurchaseamt string `xml:"directpurchaseamt"`
	Branchid          string `xml:"branchid"`
	Applid            string `xml:"applid"`
	Authpoint         string `xml:"authpoint"`
	Custremark        string `xml:"custremark"`
}

type Invoice0080 struct {
	Text        string `xml:",chardata"`
	Buyerid     string `xml:"buyerid"`
	Buyererpid  string `xml:"buyererpid"`
	Sellerid    string `xml:"sellerid"`
	Sellererpid string `xml:"sellererpid"`
	Contractno  string `xml:"contractno"`
	Invno       string `xml:"invno"`
	Invccy      string `xml:"invccy"`
	Invamt      string `xml:"invamt"`
	Invdate     string `xml:"invdate"`
	Invvaldate  string `xml:"invvaldate"`
	Invduedate  string `xml:"invduedate"`
}

//// 付款申请 b2e0081
//const BOCB2E_OPTS_PAYMENT_REQUEST = "paymentRequest"

type TrnB2e0081Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0081-rq"`
	Text      string     `xml:",chardata"`
	Ceitinfo  string     `xml:"ceitinfo"`
	B2e0081Rq *B2e0081Rq `xml:"b2e0081-rq"`
}

func (t *TrnB2e0081Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

type B2e0081Rq struct {
	Text        string           `xml:",chardata"`
	Acrcpayment *Acrcpayment0081 `xml:"acrcpayment"`
	Circlenum   string           `xml:"circlenum"`
	Payment     *Payment0081     `xml:"payment"`
}

type Payment0081 struct {
	Text        string `xml:",chardata"`
	Invref      string `xml:"invref"`
	Sellerid    string `xml:"sellerid"`
	Sellererpid string `xml:"sellererpid"`
	Invno       string `xml:"invno"`
	Pmtamt      string `xml:"pmtamt"`
	Settacno    string `xml:"settacno"`
}

type Acrcpayment0081 struct {
	Text       string `xml:",chardata"`
	Relatedref string `xml:"relatedref"`
	Applyflg   string `xml:"applyflg"`
	Erpid      string `xml:"erpid"`
	Branchid   string `xml:"branchid"`
	Applid     string `xml:"applid"`
	Pmtamt     string `xml:"pmtamt"`
	Pmtacno    string `xml:"pmtacno"`
	Fribkn     string `xml:"fribkn"`
	Pmttype    string `xml:"pmttype"`
	Pmtflg     string `xml:"pmtflg"`
	Custremark string `xml:"custremark"`
}

//// 主动推送 应收账款 转让结果 b2e0082
//const BOCB2E_OPTS_TRANSFER_RESULT_OF_ACCOUNTS_RECEIVABLE = "transferResultOfAccountsReceivable"
//
//// 主动推送 融资结果通知 b2e0083
//const BOCB2E_OPTS_FINANCING_RESULT_NOTICE = "financingResultNotice"

// b2e0057
// 转账前提额度
type TrnB2e0057Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0057-rq"`
	Text      string     `xml:",chardata"`
	B2e0057Rq *B2e0057Rq `xml:"b2e0057-rq"`
}

func (t *TrnB2e0057Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

type B2e0057Rq struct {
	Text      string         `xml:",chardata"`
	Account   *Account0057   `xml:"account"`
	Allquotas *Allquotas0057 `xml:"allquotas"`
}

type Allquotas0057 struct {
	Text      string         `xml:",chardata"`
	Cycletype string         `xml:"cycletype"`
	Period    string         `xml:"period"`
	Interval  string         `xml:"interval"`
	Startdt   string         `xml:"startdt"`
	Enddt     string         `xml:"enddt"`
	Quotalist *Quotalist0057 `xml:"quotalist"`
}

type Account0057 struct {
	Text   string `xml:",chardata"`
	Ibknum string `xml:"ibknum"`
	Actacn string `xml:"actacn"`
}

type Quotalist0057 struct {
	Text      string `xml:",chardata"`
	Chltyp    string `xml:"chltyp"`
	Txntyp    string `xml:"txntyp"`
	Quttyp    string `xml:"quttyp"`
	Cycletype string `xml:"cycletype"`
	Quota     string `xml:"quota"`
	Accutype  string `xml:"accutype"`
	Resettype string `xml:"resettype"`
	Tempquota string `xml:"tempquota"`
}

// b2e0007
// 交易状态查询接口
type TrnB2e0007Rq struct {
	XMLName   xml.Name     `xml:"trn-b2e0007-rq"`
	Text      string       `xml:",chardata"`
	B2e0007Rq *[]B2e0007Rq `xml:"b2e0007-rq"`
}

func (t *TrnB2e0007Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

type B2e0007Rq struct {
	Text   string `xml:",chardata"`
	Insid  string `xml:"insid"`
	Obssid string `xml:"obssid"`
}

// b2e0058
// 限额参数管理 - 信息查询
type TrnB2e0058Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0058-rq"`
	Text      string     `xml:",chardata"`
	B2e0058Rq *B2e0058Rq `xml:"b2e0058-rq"`
}

func (t *TrnB2e0058Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

type B2e0058Rq struct {
	Text        string       `xml:",chardata"`
	Account0058 *Account0058 `xml:"account"`
}

type Account0058 struct {
	Text   string `xml:",chardata"`
	Ibknum string `xml:"ibknum"`
	Actacn string `xml:"actacn"`
}

// b2e0500 电子回单查询
type TrnB2e0500Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0500-rq"`
	Text      string     `xml:",chardata"`
	B2e0500Rq *B2e0500Rq `xml:"b2e0500-rq"`
}

func (t *TrnB2e0500Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

type B2e0500Rq struct {
	Text       string          `xml:",chardata"`
	Tratype    string          `xml:"tratype"`
	Amtscope   *Amtscope0500   `xml:"amtscope"`
	Datescope  *Datescope0500  `xml:"datescope"`
	Domaccount *Domaccount0500 `xml:"domaccount"`
	Insid      string          `xml:"insid"`
	Multidata  *Multidata0500  `xml:"multidata"`
}

type Amtscope0500 struct {
	Text   string `xml:",chardata"`
	Minamt string `xml:"minamt"`
	Maxamt string `xml:"maxamt"`
}

type Datescope0500 struct {
	Text string `xml:",chardata"`
	From string `xml:"from"`
	To   string `xml:"to"`
}

type Domaccount0500 struct {
	Text      string `xml:",chardata"`
	Transact  string `xml:"transact"`
	Transname string `xml:"transname"`
}

type Multidata0500 struct {
	Text        string             `xml:",chardata"`
	Confirmdata *[]Confirmdata0500 `xml:"confirmdata"`
}

type Confirmdata0500 struct {
	Text   string `xml:",chardata"`
	Actacn string `xml:"actacn"`
	Ibknum string `xml:"ibknum"`
}

// b2e0483 主动查询 保理结果
type TrnB2e0483Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0483-rq"`
	Text      string     `xml:",chardata"`
	B2e0483Rq *B2e0483Rq `xml:"b2e0483-rq"`
}

type B2e0483Rq struct {
	Text       string `xml:",chardata"`
	Relatedref string `xml:"relatedref"`
	Obssid     string `xml:"obssid"`
}

func (t *TrnB2e0483Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

// b2e0232 补平用
type TrnB2e0232Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0232-rq"`
	Text      string     `xml:",chardata"`
	B2e0232Rq *B2e0232Rq `xml:"b2e0232-rq"`
}

type B2e0232Rq struct {
	Text           string      `xml:",chardata"`
	Account        Account0232 `xml:"account"`
	Supactn        Supactn0232 `xml:"supactn"`
	Trntyp         string      `xml:"trntyp"`
	Ruleid         string      `xml:"ruleid"`
	Coltypset      string      `xml:"coltypset"`
	Balance        string      `xml:"balance"`
	Time           string      `xml:"time"`
	Restoreperc    string      `xml:"restoreperc"`
	Restoredate    string      `xml:"restoredate"`
	Statcycleunit  string      `xml:"statcycleunit"`
	Statcycledigit string      `xml:"statcycledigit"`
	Targetamount   string      `xml:"targetamount"`
}

type Supactn0232 struct {
	Text    string `xml:",chardata"`
	Supbkno string `xml:"supbkno"`
	Supibkn string `xml:"supibkn"`
	Actacn  string `xml:"actacn"`
}

type Account0232 struct {
	Text   string `xml:",chardata"`
	Ibknum string `xml:"ibknum"`
	Actacn string `xml:"actacn"`
}

func (t *TrnB2e0232Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

// b2e0358 对账服务-余额对账单查询
type TrnB2e0358Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0358-rq"`
	Text      string     `xml:",chardata"`
	B2e0358Rq *B2e0358Rq `xml:"b2e0358-rq"`
}

type B2e0358Rq struct {
	Text        string `xml:",chardata"`
	Checkmonth  string `xml:"checkmonth"`
	Overdueflag string `xml:"overdueflag"`
	Conagrno    string `xml:"conagrno"`
	Begnum      string `xml:"begnum"`
	Recnum      string `xml:"recnum"`
}

func (t *TrnB2e0358Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

// b2e0359 对账单下载
type TrnB2e0359Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0359-rq"`
	Text      string     `xml:",chardata"`
	B2e0359Rq *B2e0359Rq `xml:"b2e0359-rq"`
}

type B2e0359Rq struct {
	Text    string `xml:",chardata"`
	Insid   string `xml:"insid"`
	Stmtnum string `xml:"stmtnum"`
}

func (t *TrnB2e0359Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

// b2e0360 账单核对结果反馈
type TrnB2e0360Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0360-rq"`
	Text      string     `xml:",chardata"`
	B2e0360Rq *B2e0360Rq `xml:"b2e0360-rq"`
}

type B2e0360Rq struct {
	Text      string           `xml:",chardata"`
	Stmtnum   string           `xml:"stmtnum"`
	Isrestore string           `xml:"isrestore"`
	Insid     string           `xml:"insid"`
	Multidata *[]Multidata0360 `xml:"multidata"`
}

type Multidata0360 struct {
	Text        string             `xml:",chardata"`
	Confirmdata *[]Confirmdata0360 `xml:"confirmdata"`
}

type Confirmdata0360 struct {
	Text     string `xml:",chardata"`
	Actacn   string `xml:"actacn"`
	Bookno   string `xml:"bookno"`
	Seqno    string `xml:"seqno"`
	Bokbal   string `xml:"bokbal"`
	Vchnum   string `xml:"vchnum"`
	Amount   string `xml:"amount"`
	Moddate  string `xml:"moddate"`
	Memo     string `xml:"memo"`
	Opreason string `xml:"opreason"`
}

func (t *TrnB2e0360Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}

// b2e0043 联行号查询结果
type TrnB2e0043Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0043-rq"`
	Text      string     `xml:",chardata"`
	B2e0043Rq *B2e0043Rq `xml:"b2e0043-rq"`
}

type B2e0043Rq struct {
	Text    string `xml:",chardata"`
	Tasktpy string `xml:"tasktpy"`
}

func (t *TrnB2e0043Rq) XMLMarshal() (string, error) {
	res, err := xml.Marshal(t)
	if err == nil {
		return string(res), err
	} else {
		return "", err
	}
}
