package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceCoreInvUploadAPIRequest 发票中台-发票结果回传 API请求
// alibaba.einvoice.core.inv.upload
//
// 发票回传接口适用于以下场景：
// ① 阿里发票平台向ISV提交原始发票申请，ISV开具发票成功后，基于申请ID(apply_id)回传发票至阿里发票平台进行归集与交付。
// ② 直接回传发票给阿里发票平台，进行归集，并交付给业务前台和用户。
type AlibabaEinvoiceCoreInvUploadAPIRequest struct {
	model.Params
	// 发票明细。source=apply时必填，其他source可为空
	_invoiceItems []InvoiceResultItemDto
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
	// 开票人
	_payeeOperator string
	// 购方税务登记证号，由大写字母或数字组成，长度要求15~20位。  开企业抬头时必填， 专票必填。
	_payerRegisterNo string
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
	// 发票板式文件数据，字节数据。  电票时必传。
	_invoiceFileData *model.File
	// 开票发票类型  可选值：  0: 电票  1：纸质普票  2：纸质专票
	_invoiceKind int64
	// 抬头类型。可选值：  0：个人  1：企业
	_businessType int64
}

// NewAlibabaEinvoiceCoreInvUploadRequest 初始化AlibabaEinvoiceCoreInvUploadAPIRequest对象
func NewAlibabaEinvoiceCoreInvUploadRequest() *AlibabaEinvoiceCoreInvUploadAPIRequest {
	return &AlibabaEinvoiceCoreInvUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.core.inv.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetInvoiceItems is InvoiceItems Setter
// 发票明细。source=apply时必填，其他source可为空
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetInvoiceItems(_invoiceItems []InvoiceResultItemDto) error {
	r._invoiceItems = _invoiceItems
	r.Set("invoice_items", _invoiceItems)
	return nil
}

// GetInvoiceItems InvoiceItems Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetInvoiceItems() []InvoiceResultItemDto {
	return r._invoiceItems
}

// SetSumTax is SumTax Setter
// 合计税额，格式为2位小数。  当开红票时，该字段为负数
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetSumTax(_sumTax string) error {
	r._sumTax = _sumTax
	r.Set("sum_tax", _sumTax)
	return nil
}

// GetSumTax SumTax Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetSumTax() string {
	return r._sumTax
}

// SetPayeePhone is PayeePhone Setter
// 销方联系电话。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayeePhone(_payeePhone string) error {
	r._payeePhone = _payeePhone
	r.Set("payee_phone", _payeePhone)
	return nil
}

// GetPayeePhone PayeePhone Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayeePhone() string {
	return r._payeePhone
}

// SetPlatformTid is PlatformTid Setter
// 业务平台发票申请对应的订单号。  用于source=upload时区分业务平台订单号。  source=apply时可空
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPlatformTid(_platformTid string) error {
	r._platformTid = _platformTid
	r.Set("platform_tid", _platformTid)
	return nil
}

// GetPlatformTid PlatformTid Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPlatformTid() string {
	return r._platformTid
}

// SetSumPrice is SumPrice Setter
// 合计金额（不含税），格式为2位小数。  当开红票时，该字段为负数
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetSumPrice(_sumPrice string) error {
	r._sumPrice = _sumPrice
	r.Set("sum_price", _sumPrice)
	return nil
}

// GetSumPrice SumPrice Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetSumPrice() string {
	return r._sumPrice
}

// SetInvoiceAmount is InvoiceAmount Setter
// 合计含税金额（开票金额），格式为2位小数。  当开红票时，该字段为负数。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetInvoiceAmount(_invoiceAmount string) error {
	r._invoiceAmount = _invoiceAmount
	r.Set("invoice_amount", _invoiceAmount)
	return nil
}

// GetInvoiceAmount InvoiceAmount Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetInvoiceAmount() string {
	return r._invoiceAmount
}

// SetSource is Source Setter
// 发票来源，可选值：  apply: 间连模式，服务商基于申请开具的发票；async：直连开票模式，ISV回传开票结果；upload：直接回传，进行归集与交付的发票；
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetSource() string {
	return r._source
}

// SetInvoiceFileType is InvoiceFileType Setter
// 发票板式文件类型。可选值：  PDF, OFD。  电票时必传。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetInvoiceFileType(_invoiceFileType string) error {
	r._invoiceFileType = _invoiceFileType
	r.Set("invoice_file_type", _invoiceFileType)
	return nil
}

// GetInvoiceFileType InvoiceFileType Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetInvoiceFileType() string {
	return r._invoiceFileType
}

// SetPayeeReceiver is PayeeReceiver Setter
// 收款人
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayeeReceiver(_payeeReceiver string) error {
	r._payeeReceiver = _payeeReceiver
	r.Set("payee_receiver", _payeeReceiver)
	return nil
}

// GetPayeeReceiver PayeeReceiver Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayeeReceiver() string {
	return r._payeeReceiver
}

// SetRedNoticeNo is RedNoticeNo Setter
// 红字通知单号
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetRedNoticeNo(_redNoticeNo string) error {
	r._redNoticeNo = _redNoticeNo
	r.Set("red_notice_no", _redNoticeNo)
	return nil
}

// GetRedNoticeNo RedNoticeNo Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetRedNoticeNo() string {
	return r._redNoticeNo
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 销方税务登记证号。由大写字母或数字组成，长度要求15~20位。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetReceiveMobile is ReceiveMobile Setter
// 购方手机号码，用于收票
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetReceiveMobile(_receiveMobile string) error {
	r._receiveMobile = _receiveMobile
	r.Set("receive_mobile", _receiveMobile)
	return nil
}

// GetReceiveMobile ReceiveMobile Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetReceiveMobile() string {
	return r._receiveMobile
}

// SetApplyId is ApplyId Setter
// 发票申请ID, 由阿里发票平台生成。  source=apply时 必填。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetApplyId() string {
	return r._applyId
}

// SetPayeeName is PayeeName Setter
// 销方名称
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayeeName(_payeeName string) error {
	r._payeeName = _payeeName
	r.Set("payee_name", _payeeName)
	return nil
}

// GetPayeeName PayeeName Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayeeName() string {
	return r._payeeName
}

// SetQrCode is QrCode Setter
// 二维码
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetQrCode(_qrCode string) error {
	r._qrCode = _qrCode
	r.Set("qr_code", _qrCode)
	return nil
}

// GetQrCode QrCode Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetQrCode() string {
	return r._qrCode
}

// SetLevyType is LevyType Setter
// 征税方式，0普通征收，1减按征收，2差额征收
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetLevyType(_levyType string) error {
	r._levyType = _levyType
	r.Set("levy_type", _levyType)
	return nil
}

// GetLevyType LevyType Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetLevyType() string {
	return r._levyType
}

// SetInvoiceType is InvoiceType Setter
// 发票(开票)类型，可选值：  blue: 蓝票  red: 红票
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetInvoiceType(_invoiceType string) error {
	r._invoiceType = _invoiceType
	r.Set("invoice_type", _invoiceType)
	return nil
}

// GetInvoiceType InvoiceType Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetInvoiceType() string {
	return r._invoiceType
}

// SetPayerName is PayerName Setter
// 购方抬头
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayerName(_payerName string) error {
	r._payerName = _payerName
	r.Set("payer_name", _payerName)
	return nil
}

// GetPayerName PayerName Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayerName() string {
	return r._payerName
}

// SetInvoiceNo is InvoiceNo Setter
// 发票号码
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetInvoiceNo(_invoiceNo string) error {
	r._invoiceNo = _invoiceNo
	r.Set("invoice_no", _invoiceNo)
	return nil
}

// GetInvoiceNo InvoiceNo Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetInvoiceNo() string {
	return r._invoiceNo
}

// SetInvoiceMemo is InvoiceMemo Setter
// 发票备注，会显示在票面
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetInvoiceMemo(_invoiceMemo string) error {
	r._invoiceMemo = _invoiceMemo
	r.Set("invoice_memo", _invoiceMemo)
	return nil
}

// GetInvoiceMemo InvoiceMemo Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetInvoiceMemo() string {
	return r._invoiceMemo
}

// SetPayerEmail is PayerEmail Setter
// 购方电子邮箱，需满足邮箱格式。  格式要求：\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayerEmail(_payerEmail string) error {
	r._payerEmail = _payerEmail
	r.Set("payer_email", _payerEmail)
	return nil
}

// GetPayerEmail PayerEmail Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayerEmail() string {
	return r._payerEmail
}

// SetAntiFakeCode is AntiFakeCode Setter
// 发票防伪码/密码
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetAntiFakeCode(_antiFakeCode string) error {
	r._antiFakeCode = _antiFakeCode
	r.Set("anti_fake_code", _antiFakeCode)
	return nil
}

// GetAntiFakeCode AntiFakeCode Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetAntiFakeCode() string {
	return r._antiFakeCode
}

// SetPayerBankAccountId is PayerBankAccountId Setter
// 购方银行账号， 专票必填。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayerBankAccountId(_payerBankAccountId string) error {
	r._payerBankAccountId = _payerBankAccountId
	r.Set("payer_bank_account_id", _payerBankAccountId)
	return nil
}

// GetPayerBankAccountId PayerBankAccountId Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayerBankAccountId() string {
	return r._payerBankAccountId
}

// SetPayerBankName is PayerBankName Setter
// 购方开户行名称，  专票必填。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayerBankName(_payerBankName string) error {
	r._payerBankName = _payerBankName
	r.Set("payer_bank_name", _payerBankName)
	return nil
}

// GetPayerBankName PayerBankName Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayerBankName() string {
	return r._payerBankName
}

// SetPayeeBankAccountId is PayeeBankAccountId Setter
// 销方银行账号
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayeeBankAccountId(_payeeBankAccountId string) error {
	r._payeeBankAccountId = _payeeBankAccountId
	r.Set("payee_bank_account_id", _payeeBankAccountId)
	return nil
}

// GetPayeeBankAccountId PayeeBankAccountId Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayeeBankAccountId() string {
	return r._payeeBankAccountId
}

// SetPayeeChecker is PayeeChecker Setter
// 复核人
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayeeChecker(_payeeChecker string) error {
	r._payeeChecker = _payeeChecker
	r.Set("payee_checker", _payeeChecker)
	return nil
}

// GetPayeeChecker PayeeChecker Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayeeChecker() string {
	return r._payeeChecker
}

// SetPayerPhone is PayerPhone Setter
// 购方联系电话，  专票必填。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayerPhone(_payerPhone string) error {
	r._payerPhone = _payerPhone
	r.Set("payer_phone", _payerPhone)
	return nil
}

// GetPayerPhone PayerPhone Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayerPhone() string {
	return r._payerPhone
}

// SetNormalInvoiceCode is NormalInvoiceCode Setter
// 原发票代码(开红票时必须)
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetNormalInvoiceCode(_normalInvoiceCode string) error {
	r._normalInvoiceCode = _normalInvoiceCode
	r.Set("normal_invoice_code", _normalInvoiceCode)
	return nil
}

// GetNormalInvoiceCode NormalInvoiceCode Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetNormalInvoiceCode() string {
	return r._normalInvoiceCode
}

// SetDeviceNo is DeviceNo Setter
// 开票分机号/机器编号
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetDeviceNo(_deviceNo string) error {
	r._deviceNo = _deviceNo
	r.Set("device_no", _deviceNo)
	return nil
}

// GetDeviceNo DeviceNo Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetDeviceNo() string {
	return r._deviceNo
}

// SetInvoiceDate is InvoiceDate Setter
// 开票日期，格式 yyyy-MM-dd
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetInvoiceDate(_invoiceDate string) error {
	r._invoiceDate = _invoiceDate
	r.Set("invoice_date", _invoiceDate)
	return nil
}

// GetInvoiceDate InvoiceDate Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetInvoiceDate() string {
	return r._invoiceDate
}

// SetInvoiceCode is InvoiceCode Setter
// 发票代码
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetInvoiceCode(_invoiceCode string) error {
	r._invoiceCode = _invoiceCode
	r.Set("invoice_code", _invoiceCode)
	return nil
}

// GetInvoiceCode InvoiceCode Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetInvoiceCode() string {
	return r._invoiceCode
}

// SetCheckCode is CheckCode Setter
// 校验码
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetCheckCode(_checkCode string) error {
	r._checkCode = _checkCode
	r.Set("check_code", _checkCode)
	return nil
}

// GetCheckCode CheckCode Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetCheckCode() string {
	return r._checkCode
}

// SetPayeeAddress is PayeeAddress Setter
// 销方地址。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayeeAddress(_payeeAddress string) error {
	r._payeeAddress = _payeeAddress
	r.Set("payee_address", _payeeAddress)
	return nil
}

// GetPayeeAddress PayeeAddress Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayeeAddress() string {
	return r._payeeAddress
}

// SetNormalInvoiceNo is NormalInvoiceNo Setter
// 原发票号码(开红票时必须)
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetNormalInvoiceNo(_normalInvoiceNo string) error {
	r._normalInvoiceNo = _normalInvoiceNo
	r.Set("normal_invoice_no", _normalInvoiceNo)
	return nil
}

// GetNormalInvoiceNo NormalInvoiceNo Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetNormalInvoiceNo() string {
	return r._normalInvoiceNo
}

// SetPayerAddress is PayerAddress Setter
// 购方地址，  专票必填。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayerAddress(_payerAddress string) error {
	r._payerAddress = _payerAddress
	r.Set("payer_address", _payerAddress)
	return nil
}

// GetPayerAddress PayerAddress Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayerAddress() string {
	return r._payerAddress
}

// SetPayeeBankName is PayeeBankName Setter
// 销方开户行名称。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayeeBankName(_payeeBankName string) error {
	r._payeeBankName = _payeeBankName
	r.Set("payee_bank_name", _payeeBankName)
	return nil
}

// GetPayeeBankName PayeeBankName Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayeeBankName() string {
	return r._payeeBankName
}

// SetPayeeOperator is PayeeOperator Setter
// 开票人
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayeeOperator(_payeeOperator string) error {
	r._payeeOperator = _payeeOperator
	r.Set("payee_operator", _payeeOperator)
	return nil
}

// GetPayeeOperator PayeeOperator Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayeeOperator() string {
	return r._payeeOperator
}

// SetPayerRegisterNo is PayerRegisterNo Setter
// 购方税务登记证号，由大写字母或数字组成，长度要求15~20位。  开企业抬头时必填， 专票必填。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPayerRegisterNo(_payerRegisterNo string) error {
	r._payerRegisterNo = _payerRegisterNo
	r.Set("payer_register_no", _payerRegisterNo)
	return nil
}

// GetPayerRegisterNo PayerRegisterNo Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPayerRegisterNo() string {
	return r._payerRegisterNo
}

// SetSpecialFlag is SpecialFlag Setter
// 特殊票种标识，可选值：  02: 农产品收购票
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetSpecialFlag(_specialFlag string) error {
	r._specialFlag = _specialFlag
	r.Set("special_flag", _specialFlag)
	return nil
}

// GetSpecialFlag SpecialFlag Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetSpecialFlag() string {
	return r._specialFlag
}

// SetPlatformCode is PlatformCode Setter
// 业务平台Code, 由发票中台分配。  用于source=upload时标识需交付发票的业务平台。  source=apply时可空
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPlatformCode(_platformCode string) error {
	r._platformCode = _platformCode
	r.Set("platform_code", _platformCode)
	return nil
}

// GetPlatformCode PlatformCode Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPlatformCode() string {
	return r._platformCode
}

// SetPlatformUserId is PlatformUserId Setter
// 业务平台uid
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetPlatformUserId(_platformUserId string) error {
	r._platformUserId = _platformUserId
	r.Set("platform_user_id", _platformUserId)
	return nil
}

// GetPlatformUserId PlatformUserId Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetPlatformUserId() string {
	return r._platformUserId
}

// SetBizErrorMsg is BizErrorMsg Setter
// 开票失败错误信息，  开票失败(create_result=fail)时必填。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetBizErrorMsg(_bizErrorMsg string) error {
	r._bizErrorMsg = _bizErrorMsg
	r.Set("biz_error_msg", _bizErrorMsg)
	return nil
}

// GetBizErrorMsg BizErrorMsg Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetBizErrorMsg() string {
	return r._bizErrorMsg
}

// SetBizErrorCode is BizErrorCode Setter
// 开票失败错误码，  开票失败(create_result=fail)时必填。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetBizErrorCode(_bizErrorCode string) error {
	r._bizErrorCode = _bizErrorCode
	r.Set("biz_error_code", _bizErrorCode)
	return nil
}

// GetBizErrorCode BizErrorCode Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetBizErrorCode() string {
	return r._bizErrorCode
}

// SetCreateResult is CreateResult Setter
// 开票结果，枚举值：  success: 发票开具成功；  fail: 开票失败；  source=async时必填，传实际的开票结果。其他source可不传，默认为success
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetCreateResult(_createResult string) error {
	r._createResult = _createResult
	r.Set("create_result", _createResult)
	return nil
}

// GetCreateResult CreateResult Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetCreateResult() string {
	return r._createResult
}

// SetSerialNo is SerialNo Setter
// 开票流水号/序列号，唯一标志一笔开票请求，由于阿里发票中台生成。  source=async时必填，其他source可为空
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// SetInvoiceFileData is InvoiceFileData Setter
// 发票板式文件数据，字节数据。  电票时必传。
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetInvoiceFileData(_invoiceFileData *model.File) error {
	r._invoiceFileData = _invoiceFileData
	r.Set("invoice_file_data", _invoiceFileData)
	return nil
}

// GetInvoiceFileData InvoiceFileData Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetInvoiceFileData() *model.File {
	return r._invoiceFileData
}

// SetInvoiceKind is InvoiceKind Setter
// 开票发票类型  可选值：  0: 电票  1：纸质普票  2：纸质专票
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetInvoiceKind(_invoiceKind int64) error {
	r._invoiceKind = _invoiceKind
	r.Set("invoice_kind", _invoiceKind)
	return nil
}

// GetInvoiceKind InvoiceKind Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetInvoiceKind() int64 {
	return r._invoiceKind
}

// SetBusinessType is BusinessType Setter
// 抬头类型。可选值：  0：个人  1：企业
func (r *AlibabaEinvoiceCoreInvUploadAPIRequest) SetBusinessType(_businessType int64) error {
	r._businessType = _businessType
	r.Set("business_type", _businessType)
	return nil
}

// GetBusinessType BusinessType Getter
func (r AlibabaEinvoiceCoreInvUploadAPIRequest) GetBusinessType() int64 {
	return r._businessType
}
