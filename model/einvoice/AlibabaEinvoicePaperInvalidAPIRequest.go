package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePaperInvalidAPIRequest 纸票作废接口 API请求
// alibaba.einvoice.paper.invalid
//
// 作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票
type AlibabaEinvoicePaperInvalidAPIRequest struct {
	model.Params
	// 作废操作人
	_invalidOperator string
	// 发票代码，空白作废时必填
	_invoiceCode string
	// 发票号码，空白作废时必填
	_invoiceNo string
	// 销售方纳税人识别号
	_payeeRegisterNo string
	// 开票流水号
	_serialNo string
	// 作废类型, 0=空白发票(有残缺 的纸张发票，不能做为有效报销)作废, 1=已开发票作废
	_invalidType int64
}

// NewAlibabaEinvoicePaperInvalidRequest 初始化AlibabaEinvoicePaperInvalidAPIRequest对象
func NewAlibabaEinvoicePaperInvalidRequest() *AlibabaEinvoicePaperInvalidAPIRequest {
	return &AlibabaEinvoicePaperInvalidAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoicePaperInvalidAPIRequest) Reset() {
	r._invalidOperator = ""
	r._invoiceCode = ""
	r._invoiceNo = ""
	r._payeeRegisterNo = ""
	r._serialNo = ""
	r._invalidType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePaperInvalidAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.paper.invalid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePaperInvalidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoicePaperInvalidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvalidOperator is InvalidOperator Setter
// 作废操作人
func (r *AlibabaEinvoicePaperInvalidAPIRequest) SetInvalidOperator(_invalidOperator string) error {
	r._invalidOperator = _invalidOperator
	r.Set("invalid_operator", _invalidOperator)
	return nil
}

// GetInvalidOperator InvalidOperator Getter
func (r AlibabaEinvoicePaperInvalidAPIRequest) GetInvalidOperator() string {
	return r._invalidOperator
}

// SetInvoiceCode is InvoiceCode Setter
// 发票代码，空白作废时必填
func (r *AlibabaEinvoicePaperInvalidAPIRequest) SetInvoiceCode(_invoiceCode string) error {
	r._invoiceCode = _invoiceCode
	r.Set("invoice_code", _invoiceCode)
	return nil
}

// GetInvoiceCode InvoiceCode Getter
func (r AlibabaEinvoicePaperInvalidAPIRequest) GetInvoiceCode() string {
	return r._invoiceCode
}

// SetInvoiceNo is InvoiceNo Setter
// 发票号码，空白作废时必填
func (r *AlibabaEinvoicePaperInvalidAPIRequest) SetInvoiceNo(_invoiceNo string) error {
	r._invoiceNo = _invoiceNo
	r.Set("invoice_no", _invoiceNo)
	return nil
}

// GetInvoiceNo InvoiceNo Getter
func (r AlibabaEinvoicePaperInvalidAPIRequest) GetInvoiceNo() string {
	return r._invoiceNo
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 销售方纳税人识别号
func (r *AlibabaEinvoicePaperInvalidAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoicePaperInvalidAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetSerialNo is SerialNo Setter
// 开票流水号
func (r *AlibabaEinvoicePaperInvalidAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaEinvoicePaperInvalidAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// SetInvalidType is InvalidType Setter
// 作废类型, 0=空白发票(有残缺 的纸张发票，不能做为有效报销)作废, 1=已开发票作废
func (r *AlibabaEinvoicePaperInvalidAPIRequest) SetInvalidType(_invalidType int64) error {
	r._invalidType = _invalidType
	r.Set("invalid_type", _invalidType)
	return nil
}

// GetInvalidType InvalidType Getter
func (r AlibabaEinvoicePaperInvalidAPIRequest) GetInvalidType() int64 {
	return r._invalidType
}

var poolAlibabaEinvoicePaperInvalidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoicePaperInvalidRequest()
	},
}

// GetAlibabaEinvoicePaperInvalidRequest 从 sync.Pool 获取 AlibabaEinvoicePaperInvalidAPIRequest
func GetAlibabaEinvoicePaperInvalidAPIRequest() *AlibabaEinvoicePaperInvalidAPIRequest {
	return poolAlibabaEinvoicePaperInvalidAPIRequest.Get().(*AlibabaEinvoicePaperInvalidAPIRequest)
}

// ReleaseAlibabaEinvoicePaperInvalidAPIRequest 将 AlibabaEinvoicePaperInvalidAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoicePaperInvalidAPIRequest(v *AlibabaEinvoicePaperInvalidAPIRequest) {
	v.Reset()
	poolAlibabaEinvoicePaperInvalidAPIRequest.Put(v)
}
