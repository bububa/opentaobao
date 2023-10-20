package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltxcsfinanceinvoiceinputAPIRequest 供应商发票录入 API请求
// tmall.txcs.finance.invoice.input
//
// 提供天猫超市外部合作商家财务：供应商发票录入
type TmalltxcsfinanceinvoiceinputAPIRequest struct {
	model.Params
	// 发票内容
	_invoiceInputDTO1 []InvoiceInputDto
	// 门店ID
	_ouCode string
}

// NewTmalltxcsfinanceinvoiceinputRequest 初始化TmalltxcsfinanceinvoiceinputAPIRequest对象
func NewTmalltxcsfinanceinvoiceinputRequest() *TmalltxcsfinanceinvoiceinputAPIRequest {
	return &TmalltxcsfinanceinvoiceinputAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltxcsfinanceinvoiceinputAPIRequest) GetApiMethodName() string {
	return "tmall.txcs.finance.invoice.input"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltxcsfinanceinvoiceinputAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltxcsfinanceinvoiceinputAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceInputDTO1 is InvoiceInputDTO1 Setter
// 发票内容
func (r *TmalltxcsfinanceinvoiceinputAPIRequest) SetInvoiceInputDTO1(_invoiceInputDTO1 []InvoiceInputDto) error {
	r._invoiceInputDTO1 = _invoiceInputDTO1
	r.Set("invoice_input_d_t_o1", _invoiceInputDTO1)
	return nil
}

// GetInvoiceInputDTO1 InvoiceInputDTO1 Getter
func (r TmalltxcsfinanceinvoiceinputAPIRequest) GetInvoiceInputDTO1() []InvoiceInputDto {
	return r._invoiceInputDTO1
}

// SetOuCode is OuCode Setter
// 门店ID
func (r *TmalltxcsfinanceinvoiceinputAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r TmalltxcsfinanceinvoiceinputAPIRequest) GetOuCode() string {
	return r._ouCode
}
