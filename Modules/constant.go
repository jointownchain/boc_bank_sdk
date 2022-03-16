package Modules

const BOCB2E_ERROR_XML_PATH = "./bocb2e/trans"

const BOCB2E_ERROR_XML_KEY = "trn-b2eerror-rs"

const BOCB2E_BANK_URL_ENV = "BOCB2E_BANK_URL_ENV"

const BOCB2E_RESPONSE_CODE_OK = "B001"
const BOCB2E_RESPONSE_CODE_HANDLING = "B054"
const BOCB2E_RESPONSE_CODE_SUBMITTED = "B059"

// code : B002, msg: 成功,未完
const BOCB2E_REPONSE_CODE_TBC = "B002"

// 医疗机构专户付款接口 b2e0009
const BOCB2E_OPTS_HOSPITAL_PAYMENT = "hospitalPayment"

// 企业专户申请保理 b2e0477
const BOCB2E_OPTS_COMPANY_FACTORING = "companyFactoring"

// 配送企业专户偿还保理 b2e0481
const BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT = "companyFactoringRepayment"

// 保理订单偿还进度通知 b2e0486
const BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_NOTICE = "companyFactoringRepaymentNotice"

// 保理订单偿还进度查询 b2e0487
const BOCB2E_OPTS_COMPANY_FACTORING_REPAYMENT_RESULT = "companyFactoringRepaymentResult"

// 专户交易流水查询接口 b2e0035
const BOCB2E_OPTS_ACCOUNT_TRANSACTIONS_QUERY = "accountTransactionsQuery"

// 医疗机构专户余额接口 b2e0005
const BOCB2E_OPTS_HOSPITAL_ACCOUNT_BALANCE = "hospitalAccountBalance"

// 发票信息提交 b2e0080
const BOCB2E_OPTS_INVOICE_SUBMIT = "invoiceSubmit"

// 付款申请 b2e0081
const BOCB2E_OPTS_PAYMENT_REQUEST = "paymentRequest"

// 提额 b2e0057
const BOCB2E_OPTS_RAISE_REQUEST = "raiseRequest"

// 主动推送 应收账款 转让结果 b2e0082
const BOCB2E_OPTS_TRANSFER_RESULT_OF_ACCOUNTS_RECEIVABLE = "transferResultOfAccountsReceivable"

// 主动推送 融资结果通知 b2e0083
const BOCB2E_OPTS_FINANCING_RESULT_NOTICE = "financingResultNotice"

// 企业专户申请保理结果查询 b2e0492
const BOCB2E_OPTS_COMPANY_FACTORING_RESULT = "companyFactoringResult"

// 主动查询转账结果 b2e0007
const BOCB2E_OPTS_PAYMENTS_QUERY = "paymentsQuery"

// 查询 限额 b2e0058
const BOCB2E_OPTS_TRANSFER_LIMIT_QUERY = "transferLimitQuery"

// 查询保理结果 b2e0483
const BOCB2E_OPTS_FACTORING_RESULT = "factoringResult"

// 发起填平操作 b2e0232
const BOCB2E_OPTS_FILL_UP = "fillUp"

// 回单下载 b2e0500
const BOCB2E_OPTS_RECEIPT_DOWNLOAD = "receiptDownload"

// b2e0358 对账服务-余额对账单查询
const BOCB2E_OPTS_BALANCE_STATEMENT_QUERY = "balanceStatementsQuery"

// b2e0359 对账单下载
const BOCB2E_OPTS_BALANCE_STATEMENT_DOWNLOAD = "balanceStatementsDownload"

// b2e0360 账单核对结果反馈
const BOCB2E_OPTS_BALANCE_STATEMENT_FEEDBACK = "balanceStatementsFeedback"

// b2e0043 联行号查询结果
const BOCB2E_OPTS_CNAPS_QUERY = "CNAPSQuery"
