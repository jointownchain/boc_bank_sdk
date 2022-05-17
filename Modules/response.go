package Modules

import "encoding/xml"

type Status struct {
	Text   string `xml:",chardata"`
	Rspcod string `xml:"rspcod"`
	Rspmsg string `xml:"rspmsg"`
}

// 各类 trans 定义
// >>>>>  Text / XMLName 类型的没有意义, 不要修改 <<<<<<

//// 医疗机构专户付款接口 b2e0009
//const BOCB2E_OPTS_HOSPITAL_PAYMENT = "hospitalPayment"
type HospitalPaymentRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0009Rs *TrnB2e0009Rs `xml:"trn-b2e0009-rs"`
}

type TrnB2e0009Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0009-rs"`
	Text      string       `xml:",chardata"`
	Status    *Status      `xml:"status"`
	B2e0009Rs *[]B2e0009Rs `xml:"b2e0009-rs"`
}

type B2e0009Rs struct {
	Text   string  `xml:",chardata"`
	Status *Status `xml:"status"`
	Insid  string  `xml:"insid"`
	Obssid string  `xml:"obssid"`
}

//// 企业专户申请保理 b2e0477
//const BOCB2E_OPTS_COMPANY_FACTORING= "companyFactoring"
type CompanyFactoringRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0477Rs *TrnB2e0477Rs `xml:"trn-b2e0477-rs"`
}

type TrnB2e0477Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0477-rs"`
	Text      string       `xml:",chardata"`
	Status    *Status      `xml:"status"`
	B2e0477Rs *[]B2e0477Rs `xml:"b2e0477-rs"`
}

type B2e0477Rs struct {
	Text   string  `xml:",chardata"`
	Status *Status `xml:"status"`
	Obssid string  `xml:"obssid"` // 网银流水号(申请号)
}

//// 企业专户申请保理结果查询 b2e0492 //TODO <= 主动推送?
//const BOCB2E_OPTS_COMPANY_FACTORING_RESULT = "companyFactoringResult"
type TrnB2e0492Rq struct {
	XMLName   xml.Name   `xml:"trn-b2e0492-rq"`
	Text      string     `xml:",chardata"`
	B2e0492Rq *B2e0492Rq `xml:"b2e0492-rq"`
}

type B2e0492Rq struct {
	Text string    `xml:",chardata"`
	Scfp *Scfp0492 `xml:"scfp"`
}

type Scfp0492 struct {
	Text            string               `xml:",chardata"`
	GeneralInfo     *GeneralInfo0492     `xml:"generalInfo"`
	TransReply      *TransReply0492      `xml:"transReply"`
	TransReturnData *TransReturnData0492 `xml:"transReturnData"`
}

type GeneralInfo0492 struct {
	Text        string        `xml:",chardata"`
	MainRef     string        `xml:"mainRef"`
	MsgType     string        `xml:"msgType"`
	MsgSendTime string        `xml:"msgSendTime"`
	TransCode   string        `xml:"transCode"`
	PageInfo    *PageInfo0492 `xml:"pageInfo"`
}

type PageInfo0492 struct {
	Text     string `xml:",chardata"`
	TotalRec string `xml:"totalRec"`
	PageSize string `xml:"pageSize"`
	StartRec string `xml:"startRec"`
}

type TransReply0492 struct {
	Text     string `xml:",chardata"`
	RtnCode  string `xml:"rtnCode"`
	ErrorMsg string `xml:"errorMsg"`
}

type TransReturnData0492 struct {
	Text               string                  `xml:",chardata"`
	AcsfInvFincApplyRs *AcsfInvFincApplyRs0492 `xml:"acsfInvFincApplyRs"`
}

type AcsfInvFincApplyRs0492 struct {
	Text            string       `xml:",chardata"`
	ApplyRef        string       `xml:"applyRef"`
	RelatedRef      string       `xml:"relatedRef"`
	ChannelFlg      string       `xml:"channelFlg"`
	ChannelCustID   string       `xml:"channelCustID"`
	ProcessDate     string       `xml:"processDate"`
	ApplyState      string       `xml:"applyState"`
	ProcessFeedback string       `xml:"processFeedback"`
	FincRef         string       `xml:"fincRef"`
	FincCcy         string       `xml:"fincCcy"`
	FincAmt         string       `xml:"fincAmt"`
	FincRate        string       `xml:"fincRate"`
	IntMthd         string       `xml:"intMthd"`
	IntAmt          string       `xml:"intAmt"`
	FloatPeriod     string       `xml:"floatPeriod"`
	IntPeriod       string       `xml:"intPeriod"`
	FincAcNo        string       `xml:"fincAcNo"`
	FincAcCcy       string       `xml:"fincAcCcy"`
	FincAcAmt       string       `xml:"fincAcAmt"`
	PayList         *PayList0492 `xml:"payList"`
	InvList         *InvList0492 `xml:"invList"`
}

type PayList0492 struct {
	Text      string   `xml:",chardata"`
	CircleNum string   `xml:"circleNum"`
	Pay       *Pay0492 `xml:"pay"`
}

type Pay0492 struct {
	Text     string `xml:",chardata"`
	CustNm   string `xml:"custNm"`
	PayAcNo  string `xml:"payAcNo"`
	PayAcCcy string `xml:"payAcCcy"`
	PayAcAmt string `xml:"payAcAmt"`
}

type InvList0492 struct {
	Text      string       `xml:",chardata"`
	CircleNum string       `xml:"circleNum"`
	Invoice   *Invoice0492 `xml:"invoice"`
}

type Invoice0492 struct {
	Text            string `xml:",chardata"`
	ApplyDtalRef    string `xml:"applyDtalRef"`
	ApplyState      string `xml:"applyState"`
	ProcessFeedback string `xml:"processFeedback"`
	InvRef          string `xml:"invRef"`
	InvNo           string `xml:"invNo"`
	SellerID        string `xml:"sellerID"`
	SellerOrganID   string `xml:"sellerOrganID"`
	SellerErpID     string `xml:"sellerErpID"`
	SellerNm        string `xml:"sellerNm"`
	BuyerID         string `xml:"buyerID"`
	BuyerOrganID    string `xml:"buyerOrganID"`
	BuyerErpID      string `xml:"buyerErpID"`
	BuyerNm         string `xml:"buyerNm"`
	InvCcy          string `xml:"invCcy"`
	InvAmt          string `xml:"invAmt"`
	InvFincCcy      string `xml:"invFincCcy"`
	InvFincAmt      string `xml:"invFincAmt"`
	FincStartDate   string `xml:"fincStartDate"`
	FincDueDate     string `xml:"fincDueDate"`
	FincDays        string `xml:"fincDays"`
	FincRate        string `xml:"fincRate"`
	InvIntBy        string `xml:"invIntBy"`
	IntCustID       string `xml:"intCustID"`
	InvIntAmt       string `xml:"invIntAmt"`
	ChgBy           string `xml:"chgBy"`
	ChgCustID       string `xml:"chgCustID"`
	CommAmtRecv     string `xml:"commAmtRecv"`
	CommAmtAccr     string `xml:"commAmtAccr"`
	HandingCcy      string `xml:"handingCcy"`
	HandingAmtRecv  string `xml:"handingAmtRecv"`
	HandingAmtAccr  string `xml:"handingAmtAccr"`
}

//// 配送企业专户偿还保理 b2e0481
//const BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT= "companyFactoringRepayment"
type CompanyFactoringRepaymentRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0481Rs *TrnB2e0481Rs `xml:"trn-b2e0481-rs"`
}

type TrnB2e0481Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0481-rs"`
	Text      string       `xml:",chardata"`
	Status    *Status      `xml:"status"`
	B2e0481Rs *[]B2e0481Rs `xml:"b2e0481-rs"`
}

type B2e0481Rs struct {
	Text   string  `xml:",chardata"`
	Status *Status `xml:"status"`
	Obssid string  `xml:"obssid"`
}

//// 保理订单偿还进度通知 b2e0486
//const BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_NOTICE = "companyFactoringRepaymentNotice"
type CompanyFactoringRepaymentNoticeRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0486Rs *TrnB2e0486Rs `xml:"trn-b2e0486-rs"`
}

type TrnB2e0486Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0486-rs"`
	Text      string       `xml:",chardata"`
	Status    *Status      `xml:"status"`
	B2e0486Rs *[]B2e0486Rs `xml:"b2e0486-rs"`
}

type B2e0486Rs struct {
	Text            string      `xml:",chardata"`
	Status          *Status     `xml:"status"`
	Obssid          string      `xml:"obssid"`
	Relatedref      string      `xml:"relatedref"`
	Channelflg      string      `xml:"channelflg"`
	Channelcustid   string      `xml:"channelcustid"`
	Processdate     string      `xml:"processdate"`
	Applystate      string      `xml:"applystate"`
	Processfeedback string      `xml:"processfeedback"`
	Settclrclenum   string      `xml:"settclrclenum"`
	Sett            *Sett0486   `xml:"sett"`
	Circlenum       string      `xml:"circlenum"`
	Invoice         Invoice0486 `xml:"invoice"`
}

type Sett0486 struct {
	Text          string `xml:",chardata"`
	Sellerid      string `xml:"sellerid"`
	Sellerorganid string `xml:"sellerorganid"`
	Settacno      string `xml:"settacno"`
	Settacccy     string `xml:"settacccy"`
	Settacamt     string `xml:"settacamt"`
	Payacno       string `xml:"payacno"`
	Payacccy      string `xml:"payacccy"`
	Payacamt      string `xml:"payacamt"`
}

type Invoice0486 struct {
	Text            string `xml:",chardata"`
	Applydtalref    string `xml:"applydtalref"`
	Applystate      string `xml:"applystate"`
	Processfeedback string `xml:"processfeedback"`
	Invref          string `xml:"invref"`
	Invno           string `xml:"invno"`
	Sellerid        string `xml:"sellerid"`
	Sellerorganid   string `xml:"sellerorganid"`
	Sellererpid     string `xml:"sellererpid"`
	Buyerid         string `xml:"buyerid"`
	Buyererpid      string `xml:"buyererpid"`
	Invccy          string `xml:"invccy"`
	Invamt          string `xml:"invamt"`
	Pmtamt          string `xml:"pmtamt"`
	Invbal          string `xml:"invbal"`
	Invfincccy      string `xml:"invfincccy"`
	Invfincamt      string `xml:"invfincamt"`
	Invrepayamt     string `xml:"invrepayamt"`
	Invfincbal      string `xml:"invfincbal"`
	Invintby        string `xml:"invintby"`
	Intcustid       string `xml:"intcustid"`
	Invintamt       string `xml:"invintamt"`
}

//// 保理订单偿还进度查询 b2e0487
//const BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_RESULT = "companyFactoringRepaymentResult"
type CompanyFactoringRepaymentResultRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0487Rs *TrnB2e0487Rs `xml:"trn-b2e0487-rs"`
}

type TrnB2e0487Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0487-rs"`
	Text      string       `xml:",chardata"`
	Status    *Status      `xml:"status"`
	B2e0487Rs *[]B2e0487Rs `xml:"b2e0487-rs"`
}

type B2e0487Rs struct {
	Text            string       `xml:",chardata"`
	Status          *Status      `xml:"status"`
	Obssid          string       `xml:"obssid"`
	Relatedref      string       `xml:"relatedref"`
	Channelflg      string       `xml:"channelflg"`
	Channelcustid   string       `xml:"channelcustid"`
	Processdate     string       `xml:"processdate"`
	Applystate      string       `xml:"applystate"`
	Processfeedback string       `xml:"processfeedback"`
	Fincccy         string       `xml:"fincccy"`
	Fincrepayamt    string       `xml:"fincrepayamt"`
	Fincintamt      string       `xml:"fincintamt"`
	Repayacno       string       `xml:"repayacno"`
	Repayacccy      string       `xml:"repayacccy"`
	Repayacamt      string       `xml:"repayacamt"`
	Paycirclenum    string       `xml:"paycirclenum"`
	Pay             *Pay0487     `xml:"pay"`
	Circlenum       string       `xml:"circlenum"`
	Invoice         *[]Invoice0487 `xml:"invoice"`
}

type Pay0487 struct {
	Text     string `xml:",chardata"`
	Custid   string `xml:"custid"`
	Custnm   string `xml:"custnm"`
	Payacno  string `xml:"payacno"`
	Payacccy string `xml:"payacccy"`
	Payacamt string `xml:"payacamt"`
}

type Invoice0487 struct {
	Text            string `xml:",chardata"`
	Applydtalref    string `xml:"applydtalref"`
	Applystate      string `xml:"applystate"`
	Processfeedback string `xml:"processfeedback"`
	Invref          string `xml:"invref"`
	Invno           string `xml:"invno"`
	Sellerid        string `xml:"sellerid"`
	Sellerorganid   string `xml:"sellerorganid"`
	Sellererpid     string `xml:"sellererpid"`
	Sellernm        string `xml:"sellernm"`
	Buyerid         string `xml:"buyerid"`
	Buyernm         string `xml:"buyernm"`
	Invccy          string `xml:"invccy"`
	Invamt          string `xml:"invamt"`
	Invfincccy      string `xml:"invfincccy"`
	Invfincamt      string `xml:"invfincamt"`
	Invrepayamt     string `xml:"invrepayamt"`
	Invfincbal      string `xml:"invfincbal"`
	Invintby        string `xml:"invintby"`
	Intcustid       string `xml:"intcustid"`
	Invintamt       string `xml:"invintamt"`
}

//// 专户交易流水查询接口 b2e0035
//const BOCB2E_OPTS_ACCOUNT_TRANSACTIONS_QUERY = "accountTransactionsQuery"
type AccountTransactionsQueryRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0035Rs *TrnB2e0035Rs `xml:"trn-b2e0035-rs"`
}

type TrnB2e0035Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0035-rs"`
	Text      string       `xml:",chardata"`
	Status    *Status      `xml:"status"`
	Totalnum  string       `xml:"totalnum"`
	Notenum   string       `xml:"notenum"`
	B2e0035Rs *[]B2e0035Rs `xml:"b2e0035-rs"`
}

type B2e0035Rs struct {
	Text         string      `xml:",chardata"`
	Status       *Status     `xml:"status"`
	Fractn       *Fractn0035 `xml:"fractn"`
	Toactn       *Toactn0035 `xml:"toactn"`
	Mactibkn     string      `xml:"mactibkn"`
	Mactacn      string      `xml:"mactacn"`
	Mactname     string      `xml:"mactname"`
	Mactbank     string      `xml:"mactbank"`
	Vchnum       string      `xml:"vchnum"`
	Transid      string      `xml:"transid"`
	Insid        string      `xml:"insid"`
	Txndate      string      `xml:"txndate"`
	Txntime      string      `xml:"txntime"`
	Txnamt       string      `xml:"txnamt"`
	Acctbal      string      `xml:"acctbal"`
	Avlnal       string      `xml:"avlnal"`
	Frzamt       string      `xml:"frzamt"`
	Overdramt    string      `xml:"overdramt"`
	Avloverdramt string      `xml:"avloverdramt"`
	Useinfo      string      `xml:"useinfo"`
	Furinfo      string      `xml:"furinfo"`
	Transtype    string      `xml:"transtype"`
	Bustype      string      `xml:"bustype"`
	Trncur       string      `xml:"trncur"`
	Direction    string      `xml:"direction"`
	Feeact       string      `xml:"feeact"`
	Feeamt       string      `xml:"feeamt"`
	Feecur       string      `xml:"feecur"`
	Valdat       string      `xml:"valdat"`
	Vouchtp      string      `xml:"vouchtp"`
	Vouchnum     string      `xml:"vouchnum"`
	Fxrate       string      `xml:"fxrate"`
	Interinfo    string      `xml:"interinfo"`
	Channelflg   string      `xml:"channelflg"`
	Commnum      string      `xml:"commnum"`
	Reserve1     string      `xml:"reserve1"`
	Reserve2     string      `xml:"reserve2"`
	Reserve3     string      `xml:"reserve3"`
}

type Fractn0035 struct {
	Text     string `xml:",chardata"`
	Ibknum   string `xml:"ibknum"`
	Actacn   string `xml:"actacn"`
	Acntname string `xml:"acntname"`
	Ibkname  string `xml:"ibkname"`
	Outref   string `xml:"outref"`
}

type Toactn0035 struct {
	Text   string `xml:",chardata"`
	Toibkn string `xml:"toibkn"`
	Actacn string `xml:"actacn"`
	Toname string `xml:"toname"`
	Tobank string `xml:"tobank"`
	Tobref string `xml:"tobref"`
}

//// 医疗机构专户余额接口 b2e0005
//const BOCB2E_OPTS_HOSPITAL_ACCOUNT_BALANCE = "hospitalAccountBalance"
type HospitalAccountBalanceRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0005Rs *TrnB2e0005Rs `xml:"trn-b2e0005-rs"`
}

type TrnB2e0005Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0005-rs"`
	Text      string       `xml:",chardata"`
	Status    Status       `xml:"status"`
	B2e0005Rs *[]B2e0005Rs `xml:"b2e0005-rs"`
}

type B2e0005Rs struct {
	Text    string       `xml:",chardata"`
	Status  Status       `xml:"status"`
	Account *Account0005 `xml:"account"`
	Balance *Balance0005 `xml:"balance"`
	Baldat  string       `xml:"baldat"`
}

type Balance0005 struct {
	Text   string `xml:",chardata"`
	Bokbal string `xml:"bokbal"`
	Avabal string `xml:"avabal"`
	Stpamt string `xml:"stpamt"`
	Ovramt string `xml:"ovramt"`
}

type Account0005 struct {
	Text    string `xml:",chardata"`
	Ibknum  string `xml:"ibknum"`
	Actacn  string `xml:"actacn"`
	Curcde  string `xml:"curcde"`
	Actname string `xml:"actname"`
}

//// 发票信息提交 b2e0080
//const BOCB2E_OPTS_INVOICE_SUBMIT = "invoiceSubmit"
type InvoiceSubmitRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0080Rs *TrnB2e0080Rs `xml:"trn-b2e0080-rs"`
}

type TrnB2e0080Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0080-rs"`
	Text      string       `xml:",chardata"`
	Status    *Status      `xml:"status"`
	B2e0080Rs *[]B2e0080Rs `xml:"b2e0080-rs"`
}

type B2e0080Rs struct {
	Text   string  `xml:",chardata"`
	Status *Status `xml:"status"`
	Obssid string  `xml:"obssid"`
}

//// 付款申请 b2e0081
//const BOCB2E_OPTS_PAYMENT_REQUEST = "paymentRequest"

type PaymentRequestRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0081Rs *TrnB2e0081Rs `xml:"trn-b2e0081-rs"`
}

type TrnB2e0081Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0081-rs"`
	Text      string       `xml:",chardata"`
	Status    *Status      `xml:"status"`
	B2e0081Rs *[]B2e0081Rs `xml:"b2e0081-rs"`
}

type B2e0081Rs struct {
	Text   string  `xml:",chardata"`
	Status *Status `xml:"status"`
	Obssid string  `xml:"obssid"`
}

//// 主动推送 应收账款 转让结果 b2e0082
//const BOCB2E_OPTS_TRANSFER_RESULT_OF_ACCOUNTS_RECEIVABLE = "transferResultOfAccountsReceivable"
type TrnB2e0082Rq struct {
	XMLName   xml.Name     `xml:"trn-b2e0082-rq"`
	Text      string       `xml:",chardata"`
	B2e0082Rq *[]B2e0082Rq `xml:"b2e0082-rq"`
}

type B2e0082Rq struct {
	Text string    `xml:",chardata"`
	Scfp *Scfp0082 `xml:"scfp"`
}

type Scfp0082 struct {
	Text            string               `xml:",chardata"`
	GeneralInfo     *GeneralInfo0082     `xml:"generalInfo"`
	TransReply      *TransReply0082      `xml:"transReply"`
	TransReturnData *TransReturnData0082 `xml:"transReturnData"`
}

type GeneralInfo0082 struct {
	Text        string        `xml:",chardata"`
	MainRef     string        `xml:"mainRef"`
	MsgType     string        `xml:"msgType"`
	MsgSendTime string        `xml:"msgSendTime"`
	TransCode   string        `xml:"transCode"`
	PageInfo    *PageInfo0082 `xml:"pageInfo"`
}

type PageInfo0082 struct {
	Text     string `xml:",chardata"`
	TotalRec string `xml:"totalRec"`
	PageSize string `xml:"pageSize"`
	StartRec string `xml:"startRec"`
}

type TransReply0082 struct {
	Text     string `xml:",chardata"`
	RtnCode  string `xml:"rtnCode"`
	ErrorMsg string `xml:"errorMsg"`
}

type TransReturnData0082 struct {
	Text         string            `xml:",chardata"`
	AcrcInvTrfRs *AcrcInvTrfRs0082 `xml:"acrcInvTrfRs"`
}

type AcrcInvTrfRs0082 struct {
	Text             string                `xml:",chardata"`
	MainRef          string                `xml:"mainRef"`
	RelatedRef       string                `xml:"relatedRef"`
	TrxDate          string                `xml:"trxDate"`
	ProcessDate      string                `xml:"processDate"`
	ApplyState       string                `xml:"applyState"`
	ChannelCustID    string                `xml:"channelCustID"`
	ProcessFeedback  string                `xml:"processFeedback"`
	InvoiceApprvData *InvoiceApprvData0082 `xml:"invoiceApprvData"`
}

type InvoiceApprvData0082 struct {
	Text         string            `xml:",chardata"`
	CircleNum    string            `xml:"circleNum"`
	InvoiceApprv *InvoiceApprv0082 `xml:"invoiceApprv"`
}

type InvoiceApprv0082 struct {
	Text             string `xml:",chardata"`
	TrfApplyRef      string `xml:"trfApplyRef"`
	BocnetRef        string `xml:"bocnetRef"`
	InvApplyState    string `xml:"invApplyState"`
	InvCheckFeedBack string `xml:"invCheckFeedBack"`
	InvNo            string `xml:"invNo"`
	InvCcy           string `xml:"invCcy"`
	InvAmt           string `xml:"invAmt"`
}

//// 主动推送 融资结果通知 b2e0083
//const BOCB2E_OPTS_FINANCING_RESULT_NOTICE = "financingResultNotice"

type TrnB2e0083Rq struct {
	XMLName   xml.Name     `xml:"trn-b2e0083-rq"`
	Text      string       `xml:",chardata"`
	B2e0083Rq *[]B2e0083Rq `xml:"b2e0083-rq"`
}

type B2e0083Rq struct {
	Text            string               `xml:",chardata"`
	GeneralInfo     *GeneralInfo0083     `xml:"generalInfo"`
	TransReply      *TransReply0083      `xml:"transReply"`
	TransReturnData *TransReturnData0083 `xml:"transReturnData"`
}

type GeneralInfo0083 struct {
	Text        string        `xml:",chardata"`
	MainRef     string        `xml:"mainRef"`
	MsgType     string        `xml:"msgType"`
	MsgSendTime string        `xml:"msgSendTime"`
	TransCode   string        `xml:"transCode"`
	PageInfo    *PageInfo0083 `xml:"pageInfo"`
}

type PageInfo0083 struct {
	Text     string `xml:",chardata"`
	TotalRec string `xml:"totalRec"`
	PageSize string `xml:"pageSize"`
	StartRec string `xml:"startRec"`
}

type FincApprv0083 struct {
	Text        string `xml:",chardata"`
	LoadApplyNo string `xml:"loadApplyNo"`
	BocnetRef   string `xml:"bocnetRef"`
}

type FincApprvData0083 struct {
	Text      string         `xml:",chardata"`
	CircleNum string         `xml:"circleNum"`
	FincApprv *FincApprv0083 `xml:"fincApprv"`
}

type AcrcFincRs0083 struct {
	Text            string             `xml:",chardata"`
	MainRef         string             `xml:"mainRef"`
	RelatedRef      string             `xml:"relatedRef"`
	TrxDate         string             `xml:"trxDate"`
	ProcessDate     string             `xml:"processDate"`
	ApplyState      string             `xml:"applyState"`
	ChannelCustID   string             `xml:"channelCustID"`
	ProcessFeedback string             `xml:"processFeedback"`
	FincApprvData   *FincApprvData0083 `xml:"fincApprvData"`
}

type TransReturnData0083 struct {
	Text       string          `xml:",chardata"`
	AcrcFincRs *AcrcFincRs0083 `xml:"acrcFincRs"`
}

type TransReply0083 struct {
	Text     string `xml:",chardata"`
	RtnCode  string `xml:"rtnCode"`
	ErrorMsg string `xml:"errorMsg"`
}

// b2e0057
// 提额 Raise
type RaiseRequestRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0057Rs *TrnB2e0057Rs `xml:"trn-b2e0057-rs"`
}

type TrnB2e0057Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0057-rs"`
	Text      string       `xml:",chardata"`
	Status    Status       `xml:"status"`
	B2e0057Rs *[]B2e0057Rs `xml:"b2e0057-rs"`
}

type B2e0057Rs struct {
	Text    string      `xml:",chardata"`
	Account Account0057 `xml:"account"`
	Status  Status      `xml:"status"`
}

// b2e0007
// 查询转账交易结果
type PaymentsQueryRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0007Rs *TrnB2e0007Rs `xml:"trn-b2e0007-rs"`
}

type TrnB2e0007Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0007-rs"`
	Text      string       `xml:",chardata"`
	Status    Status       `xml:"status"`
	B2e0007Rs *[]B2e0007Rs `xml:"b2e0007-rs"`
}

type B2e0007Rs struct {
	Text   string `xml:",chardata"`
	Status Status `xml:"status"`
	Insid  string `xml:"insid"`
	Obssid string `xml:"obssid"`
}

// b2e0058
// 限额参数管理 - 信息查询
type TransferLimitQueryRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0058Rs *TrnB2e0058Rs `xml:"trn-b2e0058-rs"`
}

type TrnB2e0058Rs struct {
	XMLName   xml.Name   `xml:"trn-b2e0058-rs"`
	Text      string     `xml:",chardata"`
	Status    Status     `xml:"status"`
	B2e0058Rs *B2e0058Rs `xml:"b2e0058-rs"`
}

type B2e0058Rs struct {
	Text      string           `xml:",chardata"`
	Account   Account0058Rs    `xml:"account"`
	Status    Status           `xml:"status"`
	Validate  string           `xml:"validate"`
	Allquotas *[]Allquotas0058 `xml:"allquotas"`
}

type Account0058Rs struct {
	Text   string `xml:",chardata"`
	Ibknum string `xml:"ibknum"`
	Actacn string `xml:"actacn"`
	Actcur string `xml:"actcur"`
}

type Allquotas0058 struct {
	Text      string           `xml:",chardata"`
	Cycletype string           `xml:"cycletype"`
	Period    string           `xml:"period"`
	Interval  string           `xml:"interval"`
	Startdt   string           `xml:"startdt"`
	Enddt     string           `xml:"enddt"`
	Totalnum  string           `xml:"totalnum"`
	Quotalist *[]Quotalist0058 `xml:"quotalist"`
}

type Quotalist0058 struct {
	Text        string `xml:",chardata"`
	Inuredate   string `xml:"inuredate"`
	Chltyp      string `xml:"chltyp"`
	Txntyp      string `xml:"txntyp"`
	Quttyp      string `xml:"quttyp"`
	Cycletype   string `xml:"cycletype"`
	Quota       string `xml:"quota"`
	Aquota      string `xml:"aquota"`
	Accutype    string `xml:"accutype"`
	Accuquota   string `xml:"accuquota"`
	Accustartdt string `xml:"accustartdt"`
	Uquota      string `xml:"uquota"`
	Tempquota   string `xml:"tempquota"`
}

// b2e0483 保理申请结果查询
type FactoringResultRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0483Rs *TrnB2e0483Rs `xml:"trn-b2e0483-rs"`
}

type TrnB2e0483Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0483-rs"`
	Text      string       `xml:",chardata"`
	Status    Status       `xml:"status"`
	B2e0483Rs *[]B2e0483Rs `xml:"b2e0483-rs"`
}

type B2e0483Rs struct {
	Text            string         `xml:",chardata"`
	Status          Status         `xml:"status"`
	Obssid          string         `xml:"obssid"`
	Relatedref      string         `xml:"relatedref"`
	Channelflg      string         `xml:"channelflg"`
	Channelcustid   string         `xml:"channelcustid"`
	Processdate     string         `xml:"processdate"`
	Applystate      string         `xml:"applystate"`
	Processfeedback string         `xml:"processfeedback"`
	Fincref         string         `xml:"fincref"`
	Fincccy         string         `xml:"fincccy"`
	Fincamt         string         `xml:"fincamt"`
	Fincrate        string         `xml:"fincrate"`
	Intmthd         string         `xml:"intmthd"`
	Intamt          string         `xml:"intamt"`
	Floatperiod     string         `xml:"floatperiod"`
	Intperiod       string         `xml:"intperiod"`
	Fincacno        string         `xml:"fincacno"`
	Fincacccy       string         `xml:"fincacccy"`
	Fincacamt       string         `xml:"fincacamt"`
	Paycirclenum    string         `xml:"paycirclenum"`
	Totalnum        string         `xml:"totalnum"`
	Circlenum       string         `xml:"circlenum"`
	Invoice         *[]Invoice0483 `xml:"invoice"`
}

type Invoice0483 struct {
	Text            string `xml:",chardata"`
	Applydtalref    string `xml:"applydtalref"`
	Applystate      string `xml:"applystate"`
	Processfeedback string `xml:"processfeedback"`
	Invref          string `xml:"invref"`
	Invno           string `xml:"invno"`
	Sellerid        string `xml:"sellerid"`
	Sellerorganid   string `xml:"sellerorganid"`
	Sellererpid     string `xml:"sellererpid"`
	Sellernm        string `xml:"sellernm"`
	Buyerid         string `xml:"buyerid"`
	Buyerorganid    string `xml:"buyerorganid"`
	Buyererpid      string `xml:"buyererpid"`
	Buyernm         string `xml:"buyernm"`
	Invccy          string `xml:"invccy"`
	Invamt          string `xml:"invamt"`
	Invfincccy      string `xml:"invfincccy"`
	Invfincamt      string `xml:"invfincamt"`
	Fincrate        string `xml:"fincrate"`
	Fincstartdate   string `xml:"fincstartdate"`
	Fincduedate     string `xml:"fincduedate"`
	Fincdays        string `xml:"fincdays"`
	Invintby        string `xml:"invintby"`
	Intcustid       string `xml:"intcustid"`
	Invintamt       string `xml:"invintamt"`
	Chgby           string `xml:"chgby"`
	Chgcustid       string `xml:"chgcustid"`
	Commamtrecv     string `xml:"commamtrecv"`
	Commamtaccr     string `xml:"commamtaccr"`
	Handingccy      string `xml:"handingccy"`
	Handingamtrecv  string `xml:"handingamtrecv"`
	Handingamtaccr  string `xml:"handingamtaccr"`
}

// b2e0232 补平用
type FillUpRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0232Rs *TrnB2e0232Rs `xml:"trn-b2e0232-rs"`
}
type TrnB2e0232Rs struct {
	XMLName   xml.Name   `xml:"trn-b2e0232-rs"`
	Text      string     `xml:",chardata"`
	Status    Status     `xml:"status"`
	B2e0232Rs *B2e0232Rs `xml:"b2e0232-rs"`
}

type B2e0232Rs struct {
	Text    string         `xml:",chardata"`
	Status  Status         `xml:"status"`
	Account *Account0232Rs `xml:"account"`
}

type Account0232Rs struct {
	Text   string `xml:",chardata"`
	Actacn string `xml:"actacn"`
	Ibknum string `xml:"ibknum"`
	Actcur string `xml:"actcur"`
}

// b2e0500 回单下载
type ReceiptDownloadRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0500Rs *TrnB2e0500Rs `xml:"trn-b2e0500-rs"`
}
type TrnB2e0500Rs struct {
	XMLName   xml.Name   `xml:"trn-b2e0500-rs"`
	Text      string     `xml:",chardata"`
	Status    Status     `xml:"status"`
	B2e0500Rs *B2e0500Rs `xml:"b2e0500-rs"`
}

type B2e0500Rs struct {
	Text     string `xml:",chardata"`
	Status   Status `xml:"status"`
	Filename string `xml:"filename"`
}

// b2e0358 余额对账单查询
type BalanceStatementsQueryRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0358Rs *TrnB2e0358Rs `xml:"trn-b2e0358-rs"`
}

type B2e0358Rs struct {
	Text        string `xml:",chardata"`
	Status      Status `xml:"status"`
	Cif         string `xml:"cif"`
	Stmtnum     string `xml:"stmtnum"`
	Checkmonth  string `xml:"checkmonth"`
	Overdueflag string `xml:"overdueflag"`
}

type TrnB2e0358Rs struct {
	XMLName   xml.Name     `xml:"trn-b2e0358-rs"`
	Text      string       `xml:",chardata"`
	Status    Status       `xml:"status"`
	Totalnum  string       `xml:"totalnum"`
	Backnum   string       `xml:"backnum"`
	B2e0358Rs *[]B2e0358Rs `xml:"b2e0358-rs"`
}

// b2e0359 对账单下载
type BalanceStatementsDownloadRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0359Rs *TrnB2e0359Rs `xml:"trn-b2e0359-rs"`
}
type TrnB2e0359Rs struct {
	XMLName   xml.Name   `xml:"trn-b2e0359-rs"`
	Text      string     `xml:",chardata"`
	Status    Status     `xml:"status"`
	B2e0359Rs *B2e0359Rs `xml:"b2e0359-rs"`
}

type B2e0359Rs struct {
	Text     string `xml:",chardata"`
	Status   Status `xml:"status"`
	Filename string `xml:"filename"`
}

// b2e0360 账单核对结果反馈
type BalanceStatementsFeedbackRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0360Rs *TrnB2e0360Rs `xml:"trn-b2e0360-rs"`
}

type TrnB2e0360Rs struct {
	XMLName   xml.Name   `xml:"trn-b2e0360-rs"`
	Text      string     `xml:",chardata"`
	Status    Status     `xml:"status"`
	B2e0360Rs *B2e0360Rs `xml:"b2e0360-rs"`
}

type B2e0360Rs struct {
	Text    string `xml:",chardata"`
	Status  Status `xml:"status"`
	Stmtnum string `xml:"stmtnum"`
}

// b2e0043 联行号查询结果
type CNAPSRspTrans struct {
	Text         string        `xml:",chardata"`
	TrnB2e0043Rs *TrnB2e0043Rs `xml:"trn-b2e0043-rs"`
}
type TrnB2e0043Rs struct {
	XMLName xml.Name `xml:"trn-b2e0043-rs"`
	Text    string   `xml:",chardata"`
	Status  Status   `xml:"status"`
	Data    string   `xml:"data"`
}
