package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceIncomeVerifyReturnAPIRequest
服务商回传发票查验的结果 API请求
alibaba.einvoice.income.verify.return

服务商回传发票查验的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的查验回传 */
type AlibabaEinvoiceIncomeVerifyReturnAPIRequest struct {
	model.Params
	// 校验码，success=true时必填
	_checksum string
	// 错误码，success=false时必填
	_errorCode string
	// 错误信息，success=false时必填
	_errorMessage string
	// 发票影像编号，type=1时必填
	_imageId string
	// 价税合计金额，success=true时必填，invoiceAmount=sumPrice+sumTax
	_invoiceAmount string
	// 发票代码，success=true时必填
	_invoiceCode string
	// 开票日期，格式为yyyy-MM-dd，success=true时必填
	_invoiceDate string
	// 发票明细
	_invoiceItems []InvoiceItem
	// 发票备注
	_invoiceMemo string
	// 发票号码，success=true时必填
	_invoiceNo string
	// 机器编号
	_machineNo string
	// 销售方地址电话
	_payeeAddressPhone string
	// 销售方银行及账号
	_payeeBankAccount string
	// 复核人
	_payeeChecker string
	// 销售方名称，success=true时必填
	_payeeName string
	// 开票人
	_payeeOperator string
	// 收款人
	_payeeReceiver string
	// 销售方纳税人识别号，success=true时必填
	_payeeRegisterNo string
	// 购买方地址电话
	_payerAddressPhone string
	// 购买方银行及账号
	_payerBankAccount string
	// 购买方名称，即发票抬头，success=true时必填
	_payerName string
	// 购买方纳税人识别号
	_payerRegisterNo string
	// 开票请求标识，扫描驱动回传type=1时填批次号
	_reqIndex string
	// 查验结果，true=成功，false=失败
	_success bool
	// 发票不含税金额，success=true时必填
	_sumPrice string
	// 发票税额，success=true时必填
	_sumTax string
	// 请求类型，0=阿里主动发起的查验，1=扫描驱动服务商主动回传查验结果
	_type int64
	// 发票状态，0=无效（作废），1=有效
	_invoiceStatus int64
}

// New
