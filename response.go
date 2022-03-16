package bocBankSDK

import (
	"encoding/xml"
	"github.com/pkg/errors"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/Modules"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/rspError"
)

// 解析 oneStepAPI 返回
func ParseResponseXML(xmlBytes []byte, opts string) (outStruct interface{}, err error) {
	if err := rspError.ParseError(xmlBytes); err != nil {
		return nil, errors.Wrap(err, "ParseResponseXML")
	}
	var trans interface{}
	switch opts {

	// 医院付款返回 xml
	case Modules.BOCB2E_OPTS_HOSPITAL_PAYMENT:
		trans = &Modules.HospitalPaymentRspTrans{}

	// 企业保理返回
	case Modules.BOCB2E_OPTS_COMPANY_FACTORING:
		trans = &Modules.CompanyFactoringRspTrans{}

	// 企业保理偿还返回
	case Modules.BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT:
		trans = &Modules.CompanyFactoringRepaymentRspTrans{}

	// 保理订单偿还进度通知
	case Modules.BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_NOTICE:
		trans = &Modules.CompanyFactoringRepaymentNoticeRspTrans{}

	// 保理订单偿还进度查询 b2e0487
	case Modules.BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_RESULT:
		trans = &Modules.CompanyFactoringRepaymentResultRspTrans{}

	// 专户交易流水查询接口 b2e0035
	case Modules.BOCB2E_OPTS_ACCOUNT_TRANSACTIONS_QUERY:
		trans = &Modules.AccountTransactionsQueryRspTrans{}

	// 医疗机构专户余额接口 b2e0005
	case Modules.BOCB2E_OPTS_HOSPITAL_ACCOUNT_BALANCE:
		trans = &Modules.HospitalAccountBalanceRspTrans{}

	// 发票信息提交 b2e0080
	//const BOCB2E_OPTS_INVOICE_SUBMIT = "invoiceSubmit"
	case Modules.BOCB2E_OPTS_INVOICE_SUBMIT:
		trans = &Modules.InvoiceSubmitRspTrans{}

	// 付款申请 b2e0081
	//const BOCB2E_OPTS_PAYMENT_REQUEST = "paymentRequest"
	case Modules.BOCB2E_OPTS_PAYMENT_REQUEST:
		trans = &Modules.PaymentRequestRspTrans{}

	// 提额 b2e0057
	case Modules.BOCB2E_OPTS_RAISE_REQUEST:
		trans = &Modules.RaiseRequestRspTrans{}

	// 转账查询 b2e0007
	case Modules.BOCB2E_OPTS_PAYMENTS_QUERY:
		trans = &Modules.PaymentsQueryRspTrans{}

	// 限额查询 b2e0058
	case Modules.BOCB2E_OPTS_TRANSFER_LIMIT_QUERY:
		trans = &Modules.TransferLimitQueryRspTrans{}

	// 查询保理结果 b2e0483
	case Modules.BOCB2E_OPTS_FACTORING_RESULT:
		trans = &Modules.FactoringResultRspTrans{}

	// fillUp 填平 b2e0232
	case Modules.BOCB2E_OPTS_FILL_UP:
		trans = &Modules.FillUpRspTrans{}

	// receiptDownload 回单下载 b2e0500
	case Modules.BOCB2E_OPTS_RECEIPT_DOWNLOAD:
		trans = &Modules.ReceiptDownloadRspTrans{}

	// b2e0043 联行号查询结果
	case Modules.BOCB2E_OPTS_CNAPS_QUERY:
		trans = &Modules.CNAPSRspTrans{}

	// 银企对账 b2e 0358-0360
	// b2e0358 对账服务-余额对账单查询
	case Modules.BOCB2E_OPTS_BALANCE_STATEMENT_QUERY:
		trans = &Modules.BalanceStatementsQueryRspTrans{}

	// b2e0359 对账单下载
	case Modules.BOCB2E_OPTS_BALANCE_STATEMENT_DOWNLOAD:
		trans = &Modules.BalanceStatementsDownloadRspTrans{}

	// b2e0360 账单核对结果反馈
	case Modules.BOCB2E_OPTS_BALANCE_STATEMENT_FEEDBACK:
		trans = &Modules.BalanceStatementsFeedbackRspTrans{}

	default:
		return nil, errors.New("No satisfied type found.")
	}

	err = xml.Unmarshal(xmlBytes, trans)
	if err != nil {
		return nil, err
	}
	return trans, nil
}

// 包内使用, 解析全部 返回的 xml
func parseResponseXML(xmlBytes []byte, opts string) (outStruct interface{}, err error) {
	if err := rspError.ParseError(xmlBytes); err != nil {
		return nil, errors.Wrap(err, "ParseResponseXML")
	}
	var trans interface{}
	switch opts {

	// 医院付款返回 xml
	case Modules.BOCB2E_OPTS_HOSPITAL_PAYMENT:
		trans = &Modules.HospitalPaymentRspTrans{}

	// 企业保理返回
	case Modules.BOCB2E_OPTS_COMPANY_FACTORING:
		trans = &Modules.CompanyFactoringRspTrans{}

	// 企业保理偿还返回
	case Modules.BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT:
		trans = &Modules.CompanyFactoringRepaymentRspTrans{}

	// 保理订单偿还进度通知
	case Modules.BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_NOTICE:
		trans = &Modules.CompanyFactoringRepaymentNoticeRspTrans{}

	// 保理订单偿还进度查询 b2e0487
	case Modules.BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_RESULT:
		trans = &Modules.CompanyFactoringRepaymentResultRspTrans{}

	// 专户交易流水查询接口 b2e0035
	case Modules.BOCB2E_OPTS_ACCOUNT_TRANSACTIONS_QUERY:
		trans = &Modules.AccountTransactionsQueryRspTrans{}

	// 医疗机构专户余额接口 b2e0005
	case Modules.BOCB2E_OPTS_HOSPITAL_ACCOUNT_BALANCE:
		trans = &Modules.HospitalAccountBalanceRspTrans{}

	// 发票信息提交 b2e0080
	//const BOCB2E_OPTS_INVOICE_SUBMIT = "invoiceSubmit"
	case Modules.BOCB2E_OPTS_INVOICE_SUBMIT:
		trans = &Modules.InvoiceSubmitRspTrans{}

	// 付款申请 b2e0081
	//const BOCB2E_OPTS_PAYMENT_REQUEST = "paymentRequest"
	case Modules.BOCB2E_OPTS_PAYMENT_REQUEST:
		trans = &Modules.PaymentRequestRspTrans{}

	// 提额 b2e0057
	case Modules.BOCB2E_OPTS_RAISE_REQUEST:
		trans = &Modules.RaiseRequestRspTrans{}

	// 转账查询 b2e0007
	case Modules.BOCB2E_OPTS_PAYMENTS_QUERY:
		trans = &Modules.PaymentsQueryRspTrans{}

	// 限额查询 b2e0058
	case Modules.BOCB2E_OPTS_TRANSFER_LIMIT_QUERY:
		trans = &Modules.TransferLimitQueryRspTrans{}

	// 查询保理结果 b2e0483
	case Modules.BOCB2E_OPTS_FACTORING_RESULT:
		trans = &Modules.FactoringResultRspTrans{}

	// fillUp 填平 b2e0232
	case Modules.BOCB2E_OPTS_FILL_UP:
		trans = &Modules.FillUpRspTrans{}

	// receiptDownload 回单下载 b2e0500
	case Modules.BOCB2E_OPTS_RECEIPT_DOWNLOAD:
		trans = &Modules.ReceiptDownloadRspTrans{}

	// b2e0043 联行号查询结果
	case Modules.BOCB2E_OPTS_CNAPS_QUERY:
		trans = &Modules.CNAPSRspTrans{}

	// 银企对账 b2e 0358-0360
	// b2e0358 对账服务-余额对账单查询
	case Modules.BOCB2E_OPTS_BALANCE_STATEMENT_QUERY:
		trans = &Modules.BalanceStatementsQueryRspTrans{}

	// b2e0359 对账单下载
	case Modules.BOCB2E_OPTS_BALANCE_STATEMENT_DOWNLOAD:
		trans = &Modules.BalanceStatementsDownloadRspTrans{}

	// b2e0360 账单核对结果反馈
	case Modules.BOCB2E_OPTS_BALANCE_STATEMENT_FEEDBACK:
		trans = &Modules.BalanceStatementsFeedbackRspTrans{}

	default:
		return nil, errors.New("No satisfied type found.")
	}

	rspXML := &Modules.Bocb2eResponseRoot{
		Head:  &Modules.Head{},
		Trans: trans,
	}

	err = xml.Unmarshal(xmlBytes, rspXML)
	if err != nil {
		return nil, err
	}
	return rspXML.Trans, nil
}
