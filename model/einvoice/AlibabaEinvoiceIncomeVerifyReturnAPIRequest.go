package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceincomeverifyreturnAPIRequest 服务商回传发票查验的结果 API请求
// alibaba.einvoice.income.verify.return
//
// 服务商回传发票查验的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的查验回传
type AlibabaeinvoiceincomeverifyreturnAPIRequest struct {
	model.Params
	// 发票明细
	_invoiceItems []InvoiceItem
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
	// 发票不含税金额，success=true时必填
	_sumPrice string
	// 发票税额，success=true时必填
	_sumTax string
	// 请求类型，0=阿里主动发起的查验，1=扫描驱动服务商主动回传查验结果
	_type int64
	// 发票状态，0=无效（作废），1=有效
	_invoiceStatus int64
	// 查验结果，true=成功，false=失败
	_success bool
}

// NewAlibabaeinvoiceincomeverifyreturnRequest 初始化AlibabaeinvoiceincomeverifyreturnAPIRequest对象
func NewAlibabaeinvoiceincomeverifyreturnRequest() *AlibabaeinvoiceincomeverifyreturnAPIRequest {
	return &AlibabaeinvoiceincomeverifyreturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.verify.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceItems is InvoiceItems Setter
// 发票明细
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetInvoiceItems(_invoiceItems []InvoiceItem) error {
	r._invoiceItems = _invoiceItems
	r.Set("invoice_items", _invoiceItems)
	return nil
}

// GetInvoiceItems InvoiceItems Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetInvoiceItems() []InvoiceItem {
	return r._invoiceItems
}

// SetChecksum is Checksum Setter
// 校验码，success=true时必填
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetChecksum(_checksum string) error {
	r._checksum = _checksum
	r.Set("checksum", _checksum)
	return nil
}

// GetChecksum Checksum Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetChecksum() string {
	return r._checksum
}

// SetErrorCode is ErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// GetErrorCode ErrorCode Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// SetErrorMessage is ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// GetErrorMessage ErrorMessage Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}

// SetImageId is ImageId Setter
// 发票影像编号，type=1时必填
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetImageId(_imageId string) error {
	r._imageId = _imageId
	r.Set("image_id", _imageId)
	return nil
}

// GetImageId ImageId Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetImageId() string {
	return r._imageId
}

// SetInvoiceAmount is InvoiceAmount Setter
// 价税合计金额，success=true时必填，invoiceAmount=sumPrice+sumTax
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetInvoiceAmount(_invoiceAmount string) error {
	r._invoiceAmount = _invoiceAmount
	r.Set("invoice_amount", _invoiceAmount)
	return nil
}

// GetInvoiceAmount InvoiceAmount Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetInvoiceAmount() string {
	return r._invoiceAmount
}

// SetInvoiceCode is InvoiceCode Setter
// 发票代码，success=true时必填
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetInvoiceCode(_invoiceCode string) error {
	r._invoiceCode = _invoiceCode
	r.Set("invoice_code", _invoiceCode)
	return nil
}

// GetInvoiceCode InvoiceCode Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetInvoiceCode() string {
	return r._invoiceCode
}

// SetInvoiceDate is InvoiceDate Setter
// 开票日期，格式为yyyy-MM-dd，success=true时必填
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetInvoiceDate(_invoiceDate string) error {
	r._invoiceDate = _invoiceDate
	r.Set("invoice_date", _invoiceDate)
	return nil
}

// GetInvoiceDate InvoiceDate Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetInvoiceDate() string {
	return r._invoiceDate
}

// SetInvoiceMemo is InvoiceMemo Setter
// 发票备注
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetInvoiceMemo(_invoiceMemo string) error {
	r._invoiceMemo = _invoiceMemo
	r.Set("invoice_memo", _invoiceMemo)
	return nil
}

// GetInvoiceMemo InvoiceMemo Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetInvoiceMemo() string {
	return r._invoiceMemo
}

// SetInvoiceNo is InvoiceNo Setter
// 发票号码，success=true时必填
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetInvoiceNo(_invoiceNo string) error {
	r._invoiceNo = _invoiceNo
	r.Set("invoice_no", _invoiceNo)
	return nil
}

// GetInvoiceNo InvoiceNo Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetInvoiceNo() string {
	return r._invoiceNo
}

// SetMachineNo is MachineNo Setter
// 机器编号
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetMachineNo(_machineNo string) error {
	r._machineNo = _machineNo
	r.Set("machine_no", _machineNo)
	return nil
}

// GetMachineNo MachineNo Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetMachineNo() string {
	return r._machineNo
}

// SetPayeeAddressPhone is PayeeAddressPhone Setter
// 销售方地址电话
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetPayeeAddressPhone(_payeeAddressPhone string) error {
	r._payeeAddressPhone = _payeeAddressPhone
	r.Set("payee_address_phone", _payeeAddressPhone)
	return nil
}

// GetPayeeAddressPhone PayeeAddressPhone Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetPayeeAddressPhone() string {
	return r._payeeAddressPhone
}

// SetPayeeBankAccount is PayeeBankAccount Setter
// 销售方银行及账号
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetPayeeBankAccount(_payeeBankAccount string) error {
	r._payeeBankAccount = _payeeBankAccount
	r.Set("payee_bank_account", _payeeBankAccount)
	return nil
}

// GetPayeeBankAccount PayeeBankAccount Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetPayeeBankAccount() string {
	return r._payeeBankAccount
}

// SetPayeeChecker is PayeeChecker Setter
// 复核人
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetPayeeChecker(_payeeChecker string) error {
	r._payeeChecker = _payeeChecker
	r.Set("payee_checker", _payeeChecker)
	return nil
}

// GetPayeeChecker PayeeChecker Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetPayeeChecker() string {
	return r._payeeChecker
}

// SetPayeeName is PayeeName Setter
// 销售方名称，success=true时必填
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetPayeeName(_payeeName string) error {
	r._payeeName = _payeeName
	r.Set("payee_name", _payeeName)
	return nil
}

// GetPayeeName PayeeName Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetPayeeName() string {
	return r._payeeName
}

// SetPayeeOperator is PayeeOperator Setter
// 开票人
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetPayeeOperator(_payeeOperator string) error {
	r._payeeOperator = _payeeOperator
	r.Set("payee_operator", _payeeOperator)
	return nil
}

// GetPayeeOperator PayeeOperator Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetPayeeOperator() string {
	return r._payeeOperator
}

// SetPayeeReceiver is PayeeReceiver Setter
// 收款人
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetPayeeReceiver(_payeeReceiver string) error {
	r._payeeReceiver = _payeeReceiver
	r.Set("payee_receiver", _payeeReceiver)
	return nil
}

// GetPayeeReceiver PayeeReceiver Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetPayeeReceiver() string {
	return r._payeeReceiver
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 销售方纳税人识别号，success=true时必填
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetPayerAddressPhone is PayerAddressPhone Setter
// 购买方地址电话
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetPayerAddressPhone(_payerAddressPhone string) error {
	r._payerAddressPhone = _payerAddressPhone
	r.Set("payer_address_phone", _payerAddressPhone)
	return nil
}

// GetPayerAddressPhone PayerAddressPhone Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetPayerAddressPhone() string {
	return r._payerAddressPhone
}

// SetPayerBankAccount is PayerBankAccount Setter
// 购买方银行及账号
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetPayerBankAccount(_payerBankAccount string) error {
	r._payerBankAccount = _payerBankAccount
	r.Set("payer_bank_account", _payerBankAccount)
	return nil
}

// GetPayerBankAccount PayerBankAccount Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetPayerBankAccount() string {
	return r._payerBankAccount
}

// SetPayerName is PayerName Setter
// 购买方名称，即发票抬头，success=true时必填
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetPayerName(_payerName string) error {
	r._payerName = _payerName
	r.Set("payer_name", _payerName)
	return nil
}

// GetPayerName PayerName Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetPayerName() string {
	return r._payerName
}

// SetPayerRegisterNo is PayerRegisterNo Setter
// 购买方纳税人识别号
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetPayerRegisterNo(_payerRegisterNo string) error {
	r._payerRegisterNo = _payerRegisterNo
	r.Set("payer_register_no", _payerRegisterNo)
	return nil
}

// GetPayerRegisterNo PayerRegisterNo Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetPayerRegisterNo() string {
	return r._payerRegisterNo
}

// SetReqIndex is ReqIndex Setter
// 开票请求标识，扫描驱动回传type=1时填批次号
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetReqIndex(_reqIndex string) error {
	r._reqIndex = _reqIndex
	r.Set("req_index", _reqIndex)
	return nil
}

// GetReqIndex ReqIndex Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetReqIndex() string {
	return r._reqIndex
}

// SetSumPrice is SumPrice Setter
// 发票不含税金额，success=true时必填
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetSumPrice(_sumPrice string) error {
	r._sumPrice = _sumPrice
	r.Set("sum_price", _sumPrice)
	return nil
}

// GetSumPrice SumPrice Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetSumPrice() string {
	return r._sumPrice
}

// SetSumTax is SumTax Setter
// 发票税额，success=true时必填
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetSumTax(_sumTax string) error {
	r._sumTax = _sumTax
	r.Set("sum_tax", _sumTax)
	return nil
}

// GetSumTax SumTax Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetSumTax() string {
	return r._sumTax
}

// SetType is Type Setter
// 请求类型，0=阿里主动发起的查验，1=扫描驱动服务商主动回传查验结果
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetType() int64 {
	return r._type
}

// SetInvoiceStatus is InvoiceStatus Setter
// 发票状态，0=无效（作废），1=有效
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetInvoiceStatus(_invoiceStatus int64) error {
	r._invoiceStatus = _invoiceStatus
	r.Set("invoice_status", _invoiceStatus)
	return nil
}

// GetInvoiceStatus InvoiceStatus Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetInvoiceStatus() int64 {
	return r._invoiceStatus
}

// SetSuccess is Success Setter
// 查验结果，true=成功，false=失败
func (r *AlibabaeinvoiceincomeverifyreturnAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// GetSuccess Success Getter
func (r AlibabaeinvoiceincomeverifyreturnAPIRequest) GetSuccess() bool {
	return r._success
}
