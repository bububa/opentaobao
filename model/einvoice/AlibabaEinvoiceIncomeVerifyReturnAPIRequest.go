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

// NewAlibabaEinvoiceIncomeVerifyReturnRequest 初始化AlibabaEinvoiceIncomeVerifyReturnAPIRequest对象
func NewAlibabaEinvoiceIncomeVerifyReturnRequest() *AlibabaEinvoiceIncomeVerifyReturnAPIRequest {
	return &AlibabaEinvoiceIncomeVerifyReturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.verify.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Checksum Setter
// 校验码，success=true时必填
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetChecksum(_checksum string) error {
	r._checksum = _checksum
	r.Set("checksum", _checksum)
	return nil
}

// Get Checksum Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetChecksum() string {
	return r._checksum
}

// Set is ErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// Get ErrorCode Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// Set is ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// Get ErrorMessage Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}

// Set is ImageId Setter
// 发票影像编号，type=1时必填
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetImageId(_imageId string) error {
	r._imageId = _imageId
	r.Set("image_id", _imageId)
	return nil
}

// Get ImageId Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetImageId() string {
	return r._imageId
}

// Set is InvoiceAmount Setter
// 价税合计金额，success=true时必填，invoiceAmount=sumPrice+sumTax
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetInvoiceAmount(_invoiceAmount string) error {
	r._invoiceAmount = _invoiceAmount
	r.Set("invoice_amount", _invoiceAmount)
	return nil
}

// Get InvoiceAmount Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetInvoiceAmount() string {
	return r._invoiceAmount
}

// Set is InvoiceCode Setter
// 发票代码，success=true时必填
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetInvoiceCode(_invoiceCode string) error {
	r._invoiceCode = _invoiceCode
	r.Set("invoice_code", _invoiceCode)
	return nil
}

// Get InvoiceCode Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetInvoiceCode() string {
	return r._invoiceCode
}

// Set is InvoiceDate Setter
// 开票日期，格式为yyyy-MM-dd，success=true时必填
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetInvoiceDate(_invoiceDate string) error {
	r._invoiceDate = _invoiceDate
	r.Set("invoice_date", _invoiceDate)
	return nil
}

// Get InvoiceDate Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetInvoiceDate() string {
	return r._invoiceDate
}

// Set is InvoiceItems Setter
// 发票明细
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetInvoiceItems(_invoiceItems []InvoiceItem) error {
	r._invoiceItems = _invoiceItems
	r.Set("invoice_items", _invoiceItems)
	return nil
}

// Get InvoiceItems Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetInvoiceItems() []InvoiceItem {
	return r._invoiceItems
}

// Set is InvoiceMemo Setter
// 发票备注
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetInvoiceMemo(_invoiceMemo string) error {
	r._invoiceMemo = _invoiceMemo
	r.Set("invoice_memo", _invoiceMemo)
	return nil
}

// Get InvoiceMemo Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetInvoiceMemo() string {
	return r._invoiceMemo
}

// Set is InvoiceNo Setter
// 发票号码，success=true时必填
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetInvoiceNo(_invoiceNo string) error {
	r._invoiceNo = _invoiceNo
	r.Set("invoice_no", _invoiceNo)
	return nil
}

// Get InvoiceNo Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetInvoiceNo() string {
	return r._invoiceNo
}

// Set is MachineNo Setter
// 机器编号
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetMachineNo(_machineNo string) error {
	r._machineNo = _machineNo
	r.Set("machine_no", _machineNo)
	return nil
}

// Get MachineNo Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetMachineNo() string {
	return r._machineNo
}

// Set is PayeeAddressPhone Setter
// 销售方地址电话
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetPayeeAddressPhone(_payeeAddressPhone string) error {
	r._payeeAddressPhone = _payeeAddressPhone
	r.Set("payee_address_phone", _payeeAddressPhone)
	return nil
}

// Get PayeeAddressPhone Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetPayeeAddressPhone() string {
	return r._payeeAddressPhone
}

// Set is PayeeBankAccount Setter
// 销售方银行及账号
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetPayeeBankAccount(_payeeBankAccount string) error {
	r._payeeBankAccount = _payeeBankAccount
	r.Set("payee_bank_account", _payeeBankAccount)
	return nil
}

// Get PayeeBankAccount Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetPayeeBankAccount() string {
	return r._payeeBankAccount
}

// Set is PayeeChecker Setter
// 复核人
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetPayeeChecker(_payeeChecker string) error {
	r._payeeChecker = _payeeChecker
	r.Set("payee_checker", _payeeChecker)
	return nil
}

// Get PayeeChecker Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetPayeeChecker() string {
	return r._payeeChecker
}

// Set is PayeeName Setter
// 销售方名称，success=true时必填
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetPayeeName(_payeeName string) error {
	r._payeeName = _payeeName
	r.Set("payee_name", _payeeName)
	return nil
}

// Get PayeeName Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetPayeeName() string {
	return r._payeeName
}

// Set is PayeeOperator Setter
// 开票人
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetPayeeOperator(_payeeOperator string) error {
	r._payeeOperator = _payeeOperator
	r.Set("payee_operator", _payeeOperator)
	return nil
}

// Get PayeeOperator Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetPayeeOperator() string {
	return r._payeeOperator
}

// Set is PayeeReceiver Setter
// 收款人
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetPayeeReceiver(_payeeReceiver string) error {
	r._payeeReceiver = _payeeReceiver
	r.Set("payee_receiver", _payeeReceiver)
	return nil
}

// Get PayeeReceiver Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetPayeeReceiver() string {
	return r._payeeReceiver
}

// Set is PayeeRegisterNo Setter
// 销售方纳税人识别号，success=true时必填
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// Get PayeeRegisterNo Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// Set is PayerAddressPhone Setter
// 购买方地址电话
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetPayerAddressPhone(_payerAddressPhone string) error {
	r._payerAddressPhone = _payerAddressPhone
	r.Set("payer_address_phone", _payerAddressPhone)
	return nil
}

// Get PayerAddressPhone Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetPayerAddressPhone() string {
	return r._payerAddressPhone
}

// Set is PayerBankAccount Setter
// 购买方银行及账号
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetPayerBankAccount(_payerBankAccount string) error {
	r._payerBankAccount = _payerBankAccount
	r.Set("payer_bank_account", _payerBankAccount)
	return nil
}

// Get PayerBankAccount Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetPayerBankAccount() string {
	return r._payerBankAccount
}

// Set is PayerName Setter
// 购买方名称，即发票抬头，success=true时必填
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetPayerName(_payerName string) error {
	r._payerName = _payerName
	r.Set("payer_name", _payerName)
	return nil
}

// Get PayerName Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetPayerName() string {
	return r._payerName
}

// Set is PayerRegisterNo Setter
// 购买方纳税人识别号
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetPayerRegisterNo(_payerRegisterNo string) error {
	r._payerRegisterNo = _payerRegisterNo
	r.Set("payer_register_no", _payerRegisterNo)
	return nil
}

// Get PayerRegisterNo Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetPayerRegisterNo() string {
	return r._payerRegisterNo
}

// Set is ReqIndex Setter
// 开票请求标识，扫描驱动回传type=1时填批次号
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetReqIndex(_reqIndex string) error {
	r._reqIndex = _reqIndex
	r.Set("req_index", _reqIndex)
	return nil
}

// Get ReqIndex Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetReqIndex() string {
	return r._reqIndex
}

// Set is Success Setter
// 查验结果，true=成功，false=失败
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// Get Success Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetSuccess() bool {
	return r._success
}

// Set is SumPrice Setter
// 发票不含税金额，success=true时必填
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetSumPrice(_sumPrice string) error {
	r._sumPrice = _sumPrice
	r.Set("sum_price", _sumPrice)
	return nil
}

// Get SumPrice Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetSumPrice() string {
	return r._sumPrice
}

// Set is SumTax Setter
// 发票税额，success=true时必填
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetSumTax(_sumTax string) error {
	r._sumTax = _sumTax
	r.Set("sum_tax", _sumTax)
	return nil
}

// Get SumTax Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetSumTax() string {
	return r._sumTax
}

// Set is Type Setter
// 请求类型，0=阿里主动发起的查验，1=扫描驱动服务商主动回传查验结果
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetType() int64 {
	return r._type
}

// Set is InvoiceStatus Setter
// 发票状态，0=无效（作废），1=有效
func (r *AlibabaEinvoiceIncomeVerifyReturnAPIRequest) SetInvoiceStatus(_invoiceStatus int64) error {
	r._invoiceStatus = _invoiceStatus
	r.Set("invoice_status", _invoiceStatus)
	return nil
}

// Get InvoiceStatus Getter
func (r AlibabaEinvoiceIncomeVerifyReturnAPIRequest) GetInvoiceStatus() int64 {
	return r._invoiceStatus
}
