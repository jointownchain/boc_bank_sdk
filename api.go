package bocBankSDK

import (
	"github.com/jointownchain/boc_bank_sdk/Modules"
	"github.com/pkg/errors"
)

// 包含全部 API 接口

// 1. 构造 xml 请求头
// NewXMLRequestHead
// 2. 构造 xml 请求体
// new(TrnB2e0009Rq)/或者别的
// 3. 调用 HospitalPayment/或者别的 方法

// HospitalPayment 医院付款接口 0009
func HospitalPayment(bh *Modules.Head, trnRq *Modules.TrnB2e0009Rq) (*Modules.HospitalPaymentRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, errors.Wrap(err, "HospitalPayment NewXMLRequest")
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_HOSPITAL_PAYMENT)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.HospitalPaymentRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0009Rs.Status.Rspcod) {
		return nil, errors.Errorf("HospitalPayment TRN error: code : %s, msg: %s",
			rsp.TrnB2e0009Rs.Status.Rspcod, rsp.TrnB2e0009Rs.Status.Rspmsg)
	}
	status := (*rsp.TrnB2e0009Rs.B2e0009Rs)[0].Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("HospitalPayment B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

// CompanyFactoring 企业保理接口 0477
func CompanyFactoring(bh *Modules.Head, trnRq *Modules.TrnB2e0477Rq) (*Modules.CompanyFactoringRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_COMPANY_FACTORING)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.CompanyFactoringRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0477Rs.Status.Rspcod) {
		return nil, errors.Errorf("CompanyFactoring TRN error: code : %s, msg: %s",
			rsp.TrnB2e0477Rs.Status.Rspcod, rsp.TrnB2e0477Rs.Status.Rspmsg)
	}
	status := (*rsp.TrnB2e0477Rs.B2e0477Rs)[0].Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("CompanyFactoring B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

// FactoringResult 保理结果查询 b2e0483
func FactoringResult(bh *Modules.Head, trnRq *Modules.TrnB2e0483Rq) (*Modules.FactoringResultRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_FACTORING_RESULT)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.FactoringResultRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0483Rs.Status.Rspcod) {
		return nil, errors.Errorf("FactoringResult TRN error: code : %s, msg: %s",
			rsp.TrnB2e0483Rs.Status.Rspcod, rsp.TrnB2e0483Rs.Status.Rspmsg)
	}
	status := (*rsp.TrnB2e0483Rs.B2e0483Rs)[0].Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("FactoringResult B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

// RaiseRequest 提额申请 b2e0057
func RaiseRequest(bh *Modules.Head, trnRq *Modules.TrnB2e0057Rq) (*Modules.RaiseRequestRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_RAISE_REQUEST)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.RaiseRequestRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0057Rs.Status.Rspcod) {
		return nil, errors.Errorf("RaiseRequest TRN error: code : %s, msg: %s",
			rsp.TrnB2e0057Rs.Status.Rspcod, rsp.TrnB2e0057Rs.Status.Rspmsg)
	}
	status := (*rsp.TrnB2e0057Rs.B2e0057Rs)[0].Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("RaiseRequest B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

// TransferLimitQuery 限额查询 b2e0058
func TransferLimitQuery(bh *Modules.Head, trnRq *Modules.TrnB2e0058Rq) (*Modules.TransferLimitQueryRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_TRANSFER_LIMIT_QUERY)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.TransferLimitQueryRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0058Rs.Status.Rspcod) {
		return nil, errors.Errorf("TransferLimitQuery TRN error: code : %s, msg: %s",
			rsp.TrnB2e0058Rs.Status.Rspcod, rsp.TrnB2e0058Rs.Status.Rspmsg)
	}
	status := rsp.TrnB2e0058Rs.B2e0058Rs.Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("TransferLimitQuery B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

// PaymentsQuery 支付结果查询 b2e0007
func PaymentsQuery(bh *Modules.Head, trnRq *Modules.TrnB2e0007Rq) (*Modules.PaymentsQueryRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_PAYMENTS_QUERY)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.PaymentsQueryRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0007Rs.Status.Rspcod) {
		return nil, errors.Errorf("PaymentsQuery TRN error: code : %s, msg: %s",
			rsp.TrnB2e0007Rs.Status.Rspcod, rsp.TrnB2e0007Rs.Status.Rspmsg)
	}
	// 报文交互正常不应该返回error，另外b2e0007支持多笔查询，所以必须要去掉以下逻辑
	// status := (*rsp.TrnB2e0007Rs.B2e0007Rs)[0].Status
	// if !IsRspCodeOK(status.Rspcod) {
	// 	return nil, errors.Errorf("PaymentsQuery B2E error: code : %s, msg: %s",
	// 		status.Rspcod, status.Rspmsg)
	// }
	return rsp, nil
}

// CompanyFactoringRepayment 配送企业专户偿还保理接口
func CompanyFactoringRepayment(bh *Modules.Head, trnRq *Modules.TrnB2e0481Rq) (*Modules.CompanyFactoringRepaymentRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.CompanyFactoringRepaymentRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0481Rs.Status.Rspcod) {
		return nil, errors.Errorf("CompanyFactoringRepayment TRN error: code : %s, msg: %s",
			rsp.TrnB2e0481Rs.Status.Rspcod, rsp.TrnB2e0481Rs.Status.Rspmsg)
	}
	status := (*rsp.TrnB2e0481Rs.B2e0481Rs)[0].Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("CompanyFactoringRepayment B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

// CompanyFactoringRepaymentNotice 保理订单偿还进度通知 b2e0486
func CompanyFactoringRepaymentNotice(bh *Modules.Head, trnRq *Modules.TrnB2e0486Rq) (*Modules.CompanyFactoringRepaymentNoticeRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_NOTICE)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.CompanyFactoringRepaymentNoticeRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0486Rs.Status.Rspcod) {
		return nil, errors.Errorf("CompanyFactoringRepaymentNotice TRN error: code : %s, msg: %s",
			rsp.TrnB2e0486Rs.Status.Rspcod, rsp.TrnB2e0486Rs.Status.Rspmsg)
	}
	status := (*rsp.TrnB2e0486Rs.B2e0486Rs)[0].Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("CompanyFactoringRepaymentNotice B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

// CompanyFactoringRepaymentResult 保理订单偿还进度查询 b2e0487
func CompanyFactoringRepaymentResult(bh *Modules.Head, trnRq *Modules.TrnB2e0487Rq) (*Modules.CompanyFactoringRepaymentResultRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_RESULT)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.CompanyFactoringRepaymentResultRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0487Rs.Status.Rspcod) {
		return nil, errors.Errorf("CompanyFactoringRepaymentResult TRN error: code : %s, msg: %s",
			rsp.TrnB2e0487Rs.Status.Rspcod, rsp.TrnB2e0487Rs.Status.Rspmsg)
	}
	status := (*rsp.TrnB2e0487Rs.B2e0487Rs)[0].Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("CompanyFactoringRepaymentResult B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

// AccountTransactionsQuery 专户交易流水查询接口 b2e0035
func AccountTransactionsQuery(bh *Modules.Head, trnRq *Modules.TrnB2e0035Rq) (*Modules.AccountTransactionsQueryRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_ACCOUNT_TRANSACTIONS_QUERY)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.AccountTransactionsQueryRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0035Rs.Status.Rspcod) {
		return nil, errors.Errorf("AccountTransactionsQuery TRN error: code : %s, msg: %s",
			rsp.TrnB2e0035Rs.Status.Rspcod, rsp.TrnB2e0035Rs.Status.Rspmsg)
	}
	status := (*rsp.TrnB2e0035Rs.B2e0035Rs)[0].Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("AccountTransactionsQuery B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

// HospitalAccountBalance 医疗机构专户余额接口 b2e0005
//const BOCB2E_OPTS_HOSPITAL_ACCOUNT_BALANCE = "hospitalAccountBalance"
func HospitalAccountBalance(bh *Modules.Head, trnRq *Modules.TrnB2e0005Rq) (*Modules.HospitalAccountBalanceRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_HOSPITAL_ACCOUNT_BALANCE)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.HospitalAccountBalanceRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0005Rs.Status.Rspcod) {
		return nil, errors.Errorf("HospitalAccountBalance TRN error: code : %s, msg: %s",
			rsp.TrnB2e0005Rs.Status.Rspcod, rsp.TrnB2e0005Rs.Status.Rspmsg)
	}
	status := (*rsp.TrnB2e0005Rs.B2e0005Rs)[0].Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("HospitalAccountBalance B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

//// 发票信息提交 b2e0080
//const BOCB2E_OPTS_INVOICE_SUBMIT = "invoiceSubmit"
func InvoiceSubmit(bh *Modules.Head, trnRq *Modules.TrnB2e0080Rq) (*Modules.InvoiceSubmitRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_INVOICE_SUBMIT)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.InvoiceSubmitRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0080Rs.Status.Rspcod) {
		return nil, errors.Errorf("InvoiceSubmit TRN error: code : %s, msg: %s",
			rsp.TrnB2e0080Rs.Status.Rspcod, rsp.TrnB2e0080Rs.Status.Rspmsg)
	}
	status := (*rsp.TrnB2e0080Rs.B2e0080Rs)[0].Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("InvoiceSubmit B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

//// 付款申请 b2e0081
//const BOCB2E_OPTS_PAYMENT_REQUEST = "paymentRequest"
func PaymentRequest(bh *Modules.Head, trnRq *Modules.TrnB2e0081Rq) (*Modules.PaymentRequestRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_PAYMENT_REQUEST)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.PaymentRequestRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0081Rs.Status.Rspcod) {
		return nil, errors.Errorf("PaymentRequest TRN error: code : %s, msg: %s",
			rsp.TrnB2e0081Rs.Status.Rspcod, rsp.TrnB2e0081Rs.Status.Rspmsg)
	}
	status := (*rsp.TrnB2e0081Rs.B2e0081Rs)[0].Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("PaymentRequest B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

//// 填平 b2e0232
//const BOCB2E_OPTS_FILL_UP = "fillUp"
func FillUp(bh *Modules.Head, trnRq *Modules.TrnB2e0232Rq) (*Modules.FillUpRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_FILL_UP)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.FillUpRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0232Rs.Status.Rspcod) {
		return nil, errors.Errorf("FillUp TRN error: code : %s, msg: %s",
			rsp.TrnB2e0232Rs.Status.Rspcod, rsp.TrnB2e0232Rs.Status.Rspmsg)
	}
	status := rsp.TrnB2e0232Rs.B2e0232Rs.Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("FillUp B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

// 回单下载 b2e0500
//const BOCB2E_OPTS_RECEIPT_DOWNLOAD
func ReceiptDownload(bh *Modules.Head, trnRq *Modules.TrnB2e0500Rq) (*Modules.ReceiptDownloadRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_RECEIPT_DOWNLOAD)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.ReceiptDownloadRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0500Rs.Status.Rspcod) {
		return nil, errors.Errorf("ReceiptDownload TRN error: code : %s, msg: %s",
			rsp.TrnB2e0500Rs.Status.Rspcod, rsp.TrnB2e0500Rs.Status.Rspmsg)
	}
	status := rsp.TrnB2e0500Rs.B2e0500Rs.Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("ReceiptDownload B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

// // b2e0358 对账服务-余额对账单查询
//const BOCB2E_OPTS_BALANCE_STATEMENT_QUERY = "BalanceStatementsQueryRspTrans"
func BalanceStatementsQuery(bh *Modules.Head, trnRq *Modules.TrnB2e0358Rq) (*Modules.BalanceStatementsQueryRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_BALANCE_STATEMENT_QUERY)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.BalanceStatementsQueryRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0358Rs.Status.Rspcod) {
		return nil, errors.Errorf("BalanceStatementsQuery TRN error: code : %s, msg: %s",
			rsp.TrnB2e0358Rs.Status.Rspcod, rsp.TrnB2e0358Rs.Status.Rspmsg)
	}
	status := (*rsp.TrnB2e0358Rs.B2e0358Rs)[0].Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("BalanceStatementsQuery B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

//// b2e0359 对账单下载
//const BOCB2E_OPTS_BALANCE_STATEMENT_DOWNLOAD = "BalanceStatementsDownloadRspTrans"
func BalanceStatementsDownload(bh *Modules.Head, trnRq *Modules.TrnB2e0359Rq) (*Modules.BalanceStatementsDownloadRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_BALANCE_STATEMENT_DOWNLOAD)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.BalanceStatementsDownloadRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0359Rs.Status.Rspcod) {
		return nil, errors.Errorf("BalanceStatementsDownload TRN error: code : %s, msg: %s",
			rsp.TrnB2e0359Rs.Status.Rspcod, rsp.TrnB2e0359Rs.Status.Rspmsg)
	}
	status := rsp.TrnB2e0359Rs.B2e0359Rs.Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("BalanceStatementsDownload B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

//// b2e0360 账单核对结果反馈
//const BOCB2E_OPTS_BALANCE_STATEMENT_FEEDBACK = "BalanceStatementsFeedbackRspTrans"
func BalanceStatementsFeedback(bh *Modules.Head, trnRq *Modules.TrnB2e0360Rq) (*Modules.BalanceStatementsFeedbackRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_BALANCE_STATEMENT_FEEDBACK)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.BalanceStatementsFeedbackRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0360Rs.Status.Rspcod) {
		return nil, errors.Errorf("BalanceStatementsFeedback TRN error: code : %s, msg: %s",
			rsp.TrnB2e0360Rs.Status.Rspcod, rsp.TrnB2e0360Rs.Status.Rspmsg)
	}
	status := rsp.TrnB2e0360Rs.B2e0360Rs.Status
	if !IsRspCodeOK(status.Rspcod) {
		return nil, errors.Errorf("BalanceStatementsFeedback B2E error: code : %s, msg: %s",
			status.Rspcod, status.Rspmsg)
	}
	return rsp, nil
}

// b2e0043 联行号查询结果
func CNAPSQuery(bh *Modules.Head, trnRq *Modules.TrnB2e0043Rq) (*Modules.CNAPSRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_CNAPS_QUERY)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.CNAPSRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if !IsRspCodeOK(rsp.TrnB2e0043Rs.Status.Rspcod) {
		return nil, errors.Errorf("CNAPSQuery TRN error: code : %s, msg: %s",
			rsp.TrnB2e0043Rs.Status.Rspcod, rsp.TrnB2e0043Rs.Status.Rspmsg)
	}
	return rsp, nil
}
