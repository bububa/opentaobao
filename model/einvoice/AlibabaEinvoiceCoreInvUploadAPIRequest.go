package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceCoreInvUploadAPIRequest
发票中台-发票结果回传 API请求
alibaba.einvoice.core.inv.upload

发票回传接口适用于以下场景：
① 阿里发票平台向ISV提交原始发票申请，ISV开具发票成功后，基于申请ID(apply_id)回传发票至阿里发票平台进行归集与交付。
② 直接回传发票给阿里发票平台，进行归集，并交付给业务前台和用户。 */
type AlibabaEinvoiceCoreInvUploadAPIRequest struct {
	model.Params
	// 合计税额，格式为2位小数。  当开红票时，该字段为负数
	_sumTax string
	// 销方联系电话。
	_payeePhone string
	// 业务平台发票申请对应的订单号。  用于source=upload时区分业务平台订单号。  source=apply时可空
	_platformTid string
	// 合计金额（不含税），格式为2位小数。  当开红票时，该字段为负数
	_sumPrice string
	// 合计含税金额（开票金额），格式为2位小数。  当开红票时，该字段为负数。
	_invoiceAmount string
	// 发票来源，可选值：  apply: 间连模式，服务商基于申请开具的发票；async：直连开票模式，ISV回传开票结果；upload：直接回传，进行归集与交付的发票；
	_source string
	// 发票板式文件类型。可选值：  PDF, OFD。  电票时必传。
	_invoiceFileType string
	// 收款人
	_payeeReceiver string
	// 红字通知单号
	_redNoticeNo string
	// 销方税务登记证号。由大写字母或数字组成，长度要求15~20位。
	_payeeRegisterNo string
	// 购方手机号码，用于收票
	_receiveMobile string
	// 发票申请ID, 由阿里发票平台生成。  source=apply时 必填。
	_applyId string
	// 销方名称
	_payeeName string
	// 二维码
	_qrCode string
	// 征税方式，0普通征收，1减按征收，2差额征收
	_levyType string
	// 发票(开票)类型，可选值：  blue: 蓝票  red: 红票
	_invoiceType string
	// 购方抬头
	_payerName string
	// 发票号码
	_invoiceNo string
	// 发票备注，会显示在票面
	_invoiceMemo string
	// 购方电子邮箱，需满足邮箱格式。  格式要求：\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*
	_payerEmail string
	// 发票防伪码/密码
	_antiFakeCode string
	// 购方银行账号， 专票必填。
	_payerBankAccountId string
	// 购方开户行名称，  专票必填。
	_payerBankName string
	// 销方银行账号
	_payeeBankAccountId string
	// 复核人
	_payeeChecker string
	// 购方联系电话，  专票必填。
	_payerPhone string
	// 原发票代码(开红票时必须)
	_normalInvoiceCode string
	// 开票分机号/机器编号
	_deviceNo string
	// 开票日期，格式 yyyy-MM-dd
	_invoiceDate string
	// 发票代码
	_invoiceCode string
	// 校验码
	_checkCode string
	// 销方地址。
	_payeeAddress string
	// 原发票号码(开红票时必须)
	_normalInvoiceNo string
	// 购方地址，  专票必填。
	_payerAddress string
	// 销方开户行名称。
	_payeeBankName string
	// 发票明细。source=apply时必填，其他source可为空
	_invoiceItems []InvoiceResultItemDto
	// 开票人
	_payeeOperator string
	// 发票板式文件数据，字节数据。  电票时必传。
	_invoiceFileData *model.File
	// 购方税务登记证号，由大写字母或数字组成，长度要求15~20位。  开企业抬头时必填， 专票必填。
	_payerRegisterNo string
	// 开票发票类型  可选值：  0: 电票  1：纸质普票  2：纸质专票
	_invoiceKind int64
	// 抬头类型。可选值：  0：个人  1：企业
	_businessType int64
	// 特殊票种标识，可选值：  02: 农产品收购票
	_specialFlag string
	// 业务平台Code, 由发票中台分配。  用于source=upload时标识需交付发票的业务平台。  source=apply时可空
	_platformCode string
	// 业务平台uid
	_platformUserId string
	// 开票失败错误信息，  开票失败(create_result=fail)时必填。
	_bizErrorMsg string
	// 开票失败错误码，  开票失败(create_result=fail)时必填。
	_bizErrorCode string
	// 开票结果，枚举值：  success: 发票开具成功；  fail: 开票失败；  source=async时必填，传实际的开票结果。其他source可不传，默认为success
	_createResult string
	// 开票流水号/序列号，唯一标志一笔开票请求，由于阿里发票中台生成。  source=async时必填，其他source可为空
	_serialNo string
}

// New
