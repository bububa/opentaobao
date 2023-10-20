package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicepaperinvalidAPIRequest 纸票作废接口 API请求
// alibaba.einvoice.paper.invalid
//
// 作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票
type AlibabaeinvoicepaperinvalidAPIRequest struct {
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

// NewAlibabaeinvoicepaperinvalidRequest 初始化AlibabaeinvoicepaperinvalidAPIRequest对象
func NewAlibabaeinvoicepaperinvalidRequest() *AlibabaeinvoicepaperinvalidAPIRequest {
	return &AlibabaeinvoicepaperinvalidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicepaperinvalidAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.paper.invalid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicepaperinvalidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicepaperinvalidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvalidOperator is InvalidOperator Setter
// 作废操作人
func (r *AlibabaeinvoicepaperinvalidAPIRequest) SetInvalidOperator(_invalidOperator string) error {
	r._invalidOperator = _invalidOperator
	r.Set("invalid_operator", _invalidOperator)
	return nil
}

// GetInvalidOperator InvalidOperator Getter
func (r AlibabaeinvoicepaperinvalidAPIRequest) GetInvalidOperator() string {
	return r._invalidOperator
}

// SetInvoiceCode is InvoiceCode Setter
// 发票代码，空白作废时必填
func (r *AlibabaeinvoicepaperinvalidAPIRequest) SetInvoiceCode(_invoiceCode string) error {
	r._invoiceCode = _invoiceCode
	r.Set("invoice_code", _invoiceCode)
	return nil
}

// GetInvoiceCode InvoiceCode Getter
func (r AlibabaeinvoicepaperinvalidAPIRequest) GetInvoiceCode() string {
	return r._invoiceCode
}

// SetInvoiceNo is InvoiceNo Setter
// 发票号码，空白作废时必填
func (r *AlibabaeinvoicepaperinvalidAPIRequest) SetInvoiceNo(_invoiceNo string) error {
	r._invoiceNo = _invoiceNo
	r.Set("invoice_no", _invoiceNo)
	return nil
}

// GetInvoiceNo InvoiceNo Getter
func (r AlibabaeinvoicepaperinvalidAPIRequest) GetInvoiceNo() string {
	return r._invoiceNo
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 销售方纳税人识别号
func (r *AlibabaeinvoicepaperinvalidAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoicepaperinvalidAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetSerialNo is SerialNo Setter
// 开票流水号
func (r *AlibabaeinvoicepaperinvalidAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaeinvoicepaperinvalidAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// SetInvalidType is InvalidType Setter
// 作废类型, 0=空白发票(有残缺 的纸张发票，不能做为有效报销)作废, 1=已开发票作废
func (r *AlibabaeinvoicepaperinvalidAPIRequest) SetInvalidType(_invalidType int64) error {
	r._invalidType = _invalidType
	r.Set("invalid_type", _invalidType)
	return nil
}

// GetInvalidType InvalidType Getter
func (r AlibabaeinvoicepaperinvalidAPIRequest) GetInvalidType() int64 {
	return r._invalidType
}
