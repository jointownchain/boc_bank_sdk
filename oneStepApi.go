package bocBankSDK

import (
	"context"
	"encoding/xml"
	"github.com/pkg/errors"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/Modules"
)

// 根据 function name 决定调用哪个 api
//// proto 中的 银行相关参数
//// 银行参数基于 XML, 对应的 struct 中有很多标签和无用字段, 使用 xml.Marshal 传 string
//message BankRequest {
//// 请求头 xml.marshal
//string head = 1;
//// 请求体 xml.marshal
//string body = 2;
//// 请求方法 <- 按照 bankSDK 内定义的 constant opts 来
//string function = 3;

//// 通用银行信息返回(返回对应结构体 xml marshal)
//message BankResponse {
//string body = 1;
//}

func CallBank(ctx context.Context, headStr, bodyStr, funcName string) (rspBody []byte, err error) {
	head := &Modules.Head{}
	err = xml.Unmarshal([]byte(headStr), head)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal head error")
	}
	switch funcName {
	//// 医疗机构专户付款接口 b2e0009
	//const BOCB2E_OPTS_HOSPITAL_PAYMENT = "hospitalPayment"
	case Modules.BOCB2E_OPTS_HOSPITAL_PAYMENT:
		trn := &Modules.TrnB2e0009Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0009Rq error")
		}
		rsp, err := HospitalPayment(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call HospitalPayment error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal HospitalPayment result error")
		}
		return rspBody, err
	//// 企业专户申请保理 b2e0477
	//const BOCB2E_OPTS_COMPANY_FACTORING = "companyFactoring"
	case Modules.BOCB2E_OPTS_COMPANY_FACTORING:
		trn := &Modules.TrnB2e0477Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0477Rq error")
		}
		rsp, err := CompanyFactoring(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call CompanyFactoring error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal CompanyFactoring result error")
		}
		return rspBody, err
	//// 配送企业专户偿还保理 b2e0481
	//const BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT = "companyFactoringRepayment"
	case Modules.BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT:
		trn := &Modules.TrnB2e0481Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0481Rq error")
		}
		rsp, err := CompanyFactoringRepayment(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call CompanyFactoringRepayment error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal CompanyFactoringRepayment result error")
		}
		return rspBody, err
	//// 保理订单偿还进度通知 b2e0486
	//const BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_NOTICE = "companyFactoringRepaymentNotice"
	case Modules.BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_NOTICE:
		trn := &Modules.TrnB2e0486Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0486Rq error")
		}
		rsp, err := CompanyFactoringRepaymentNotice(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call CompanyFactoringRepaymentNotice error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal CompanyFactoringRepaymentNotice result error")
		}
		return rspBody, err
	//// 保理订单偿还进度查询 b2e0487
	//const BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_RESULT = "companyFactoringRepaymentResult"
	case Modules.BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_RESULT:
		trn := &Modules.TrnB2e0487Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0487Rq error")
		}
		rsp, err := CompanyFactoringRepaymentResult(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call companyFactoringRepaymentResult error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal companyFactoringRepaymentResult result error")
		}
		return rspBody, err
	//// 专户交易流水查询接口 b2e0035
	//const BOCB2E_OPTS_ACCOUNT_TRANSACTIONS_QUERY = "accountTransactionsQuery"
	case Modules.BOCB2E_OPTS_ACCOUNT_TRANSACTIONS_QUERY:
		trn := &Modules.TrnB2e0035Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0035Rq error")
		}
		rsp, err := AccountTransactionsQuery(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call AccountTransactionsQuery error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal AccountTransactionsQuery result error")
		}
		return rspBody, err
	//// 医疗机构专户余额接口 b2e0005
	//const BOCB2E_OPTS_HOSPITAL_ACCOUNT_BALANCE = "hospitalAccountBalance"
	case Modules.BOCB2E_OPTS_HOSPITAL_ACCOUNT_BALANCE:
		trn := &Modules.TrnB2e0005Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0005Rq error")
		}
		rsp, err := HospitalAccountBalance(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call HospitalAccountBalance error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal HospitalAccountBalance result error")
		}
		return rspBody, err
	//// 发票信息提交 b2e0080
	//const BOCB2E_OPTS_INVOICE_SUBMIT = "invoiceSubmit"
	case Modules.BOCB2E_OPTS_INVOICE_SUBMIT:
		trn := &Modules.TrnB2e0080Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0080Rq error")
		}
		rsp, err := InvoiceSubmit(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call InvoiceSubmit error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal InvoiceSubmit result error")
		}
		return rspBody, err
	//// 付款申请 b2e0081
	//const BOCB2E_OPTS_PAYMENT_REQUEST = "paymentRequest"
	case Modules.BOCB2E_OPTS_PAYMENT_REQUEST:
		trn := &Modules.TrnB2e0081Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0081Rq error")
		}
		rsp, err := PaymentRequest(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call PaymentRequest error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal PaymentRequest result error")
		}
		return rspBody, err
	//// 提额 b2e0057
	//const BOCB2E_OPTS_RAISE_REQUEST = "raiseRequest"
	case Modules.BOCB2E_OPTS_RAISE_REQUEST:
		trn := &Modules.TrnB2e0057Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0057Rq error")
		}
		rsp, err := RaiseRequest(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call RaiseRequest error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal RaiseRequest result error")
		}
		return rspBody, err
	//// 查询 b2e0007
	case Modules.BOCB2E_OPTS_PAYMENTS_QUERY:
		trn := &Modules.TrnB2e0007Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0057Rq error")
		}
		rsp, err := PaymentsQuery(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call RaiseRequest error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal RaiseRequest result error")
		}
		return rspBody, err
	// 查询限额 b2e0058
	case Modules.BOCB2E_OPTS_TRANSFER_LIMIT_QUERY:
		trn := &Modules.TrnB2e0058Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0058Rq error")
		}
		rsp, err := TransferLimitQuery(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call TransferLimitQuery error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal TransferLimitQuery result error")
		}
		return rspBody, err
	// 查询保理结果 b2e0483
	case Modules.BOCB2E_OPTS_FACTORING_RESULT:
		trn := &Modules.TrnB2e0483Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0483Rq error")
		}
		rsp, err := FactoringResult(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call FactoringResult error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal FactoringResult result error")
		}
		return rspBody, err
	//BOCB2E_OPTS_FILL_UP 填平 b2e0232
	case Modules.BOCB2E_OPTS_FILL_UP:
		trn := &Modules.TrnB2e0232Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0232Rq error")
		}
		rsp, err := FillUp(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call FillUp error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal FillUp result error")
		}
		return rspBody, err

	//	BOCB2E_OPTS_RECEIPT_DOWNLOAD:                   "b2e0500",
	case Modules.BOCB2E_OPTS_RECEIPT_DOWNLOAD:
		trn := &Modules.TrnB2e0500Rq{}
		err = xml.Unmarshal([]byte(bodyStr), trn)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal TrnB2e0500Rq error")
		}
		rsp, err := ReceiptDownload(head, trn)
		if err != nil {
			return nil, errors.Wrap(err, "call ReceiptDownload error")
		}
		rspBody, err = xml.Marshal(rsp)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal ReceiptDownload result error")
		}
		return rspBody, err
	default:
		return nil, errors.New("no func find by funcName:" + funcName)
	}
}
