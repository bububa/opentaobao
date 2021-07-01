package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTxcsFinanceInvoiceInputAPIRequest
供应商发票录入 API请求
tmall.txcs.finance.invoice.input

提供天猫超市外部合作商家财务：供应商发票录入 */
type TmallTxcsFinanceInvoiceInputAPIRequest struct {
	model.Params
	// 门店ID
	_ouCode string
	// 发票内容
	_invoiceInputDTO1 []InvoiceInputDto
}

// NewTmallTxcsFinanceInvoiceInputRequest 初始化TmallTxcsFinanceInvoiceInputAPIRequest对象
func NewTmallTxcsFinanceInvoiceInputRequest() *TmallTxcsFinanceInvoiceInputAPIRequest {
	return &TmallTxcsFinanceInvoiceInputAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTxcsFinanceInvoiceInputAPIRequest) GetApiMethodName() string {
	return "tmall.txcs.finance.invoice.input"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTxcsFinanceInvoiceInputAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OuCode Setter
// 门店ID
func (r *TmallTxcsFinanceInvoiceInputAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// Get OuCode Getter
func (r TmallTxcsFinanceInvoiceInputAPIRequest) GetOuCode() string {
	return r._ouCode
}

// Set is InvoiceInputDTO1 Setter
// 发票内容
func (r *TmallTxcsFinanceInvoiceInputAPIRequest) SetInvoiceInputDTO1(_invoiceInputDTO1 []InvoiceInputDto) error {
	r._invoiceInputDTO1 = _invoiceInputDTO1
	r.Set("invoice_input_d_t_o1", _invoiceInputDTO1)
	return nil
}

// Get InvoiceInputDTO1 Getter
func (r TmallTxcsFinanceInvoiceInputAPIRequest) GetInvoiceInputDTO1() []InvoiceInputDto {
	return r._invoiceInputDTO1
}
