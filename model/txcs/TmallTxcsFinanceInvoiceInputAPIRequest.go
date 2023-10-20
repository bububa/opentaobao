package txcs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTxcsFinanceInvoiceInputAPIRequest 供应商发票录入 API请求
// tmall.txcs.finance.invoice.input
//
// 提供天猫超市外部合作商家财务：供应商发票录入
type TmallTxcsFinanceInvoiceInputAPIRequest struct {
	model.Params
	// 发票内容
	_invoiceInputDTO1 []InvoiceInputDto
	// 门店ID
	_ouCode string
}

// NewTmallTxcsFinanceInvoiceInputRequest 初始化TmallTxcsFinanceInvoiceInputAPIRequest对象
func NewTmallTxcsFinanceInvoiceInputRequest() *TmallTxcsFinanceInvoiceInputAPIRequest {
	return &TmallTxcsFinanceInvoiceInputAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTxcsFinanceInvoiceInputAPIRequest) Reset() {
	r._invoiceInputDTO1 = r._invoiceInputDTO1[:0]
	r._ouCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTxcsFinanceInvoiceInputAPIRequest) GetApiMethodName() string {
	return "tmall.txcs.finance.invoice.input"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTxcsFinanceInvoiceInputAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTxcsFinanceInvoiceInputAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceInputDTO1 is InvoiceInputDTO1 Setter
// 发票内容
func (r *TmallTxcsFinanceInvoiceInputAPIRequest) SetInvoiceInputDTO1(_invoiceInputDTO1 []InvoiceInputDto) error {
	r._invoiceInputDTO1 = _invoiceInputDTO1
	r.Set("invoice_input_d_t_o1", _invoiceInputDTO1)
	return nil
}

// GetInvoiceInputDTO1 InvoiceInputDTO1 Getter
func (r TmallTxcsFinanceInvoiceInputAPIRequest) GetInvoiceInputDTO1() []InvoiceInputDto {
	return r._invoiceInputDTO1
}

// SetOuCode is OuCode Setter
// 门店ID
func (r *TmallTxcsFinanceInvoiceInputAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r TmallTxcsFinanceInvoiceInputAPIRequest) GetOuCode() string {
	return r._ouCode
}

var poolTmallTxcsFinanceInvoiceInputAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTxcsFinanceInvoiceInputRequest()
	},
}

// GetTmallTxcsFinanceInvoiceInputRequest 从 sync.Pool 获取 TmallTxcsFinanceInvoiceInputAPIRequest
func GetTmallTxcsFinanceInvoiceInputAPIRequest() *TmallTxcsFinanceInvoiceInputAPIRequest {
	return poolTmallTxcsFinanceInvoiceInputAPIRequest.Get().(*TmallTxcsFinanceInvoiceInputAPIRequest)
}

// ReleaseTmallTxcsFinanceInvoiceInputAPIRequest 将 TmallTxcsFinanceInvoiceInputAPIRequest 放入 sync.Pool
func ReleaseTmallTxcsFinanceInvoiceInputAPIRequest(v *TmallTxcsFinanceInvoiceInputAPIRequest) {
	v.Reset()
	poolTmallTxcsFinanceInvoiceInputAPIRequest.Put(v)
}
