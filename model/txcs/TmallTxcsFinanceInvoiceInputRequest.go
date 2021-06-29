package txcs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商发票录入 API请求
tmall.txcs.finance.invoice.input

提供天猫超市外部合作商家财务：供应商发票录入
*/
type TmallTxcsFinanceInvoiceInputRequest struct {
    model.Params
    // 门店ID
    _ouCode   string
    // 发票内容
    _invoiceInputDTO1   []InvoiceInputDto
}

// 初始化TmallTxcsFinanceInvoiceInputRequest对象
func NewTmallTxcsFinanceInvoiceInputRequest() *TmallTxcsFinanceInvoiceInputRequest{
    return &TmallTxcsFinanceInvoiceInputRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTxcsFinanceInvoiceInputRequest) GetApiMethodName() string {
    return "tmall.txcs.finance.invoice.input"
}

// IRequest interface 方法, 获取API参数
func (r TmallTxcsFinanceInvoiceInputRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuCode Setter
// 门店ID
func (r *TmallTxcsFinanceInvoiceInputRequest) SetOuCode(_ouCode string) error {
    r._ouCode = _ouCode
    r.Set("ou_code", _ouCode)
    return nil
}

// OuCode Getter
func (r TmallTxcsFinanceInvoiceInputRequest) GetOuCode() string {
    return r._ouCode
}
// InvoiceInputDTO1 Setter
// 发票内容
func (r *TmallTxcsFinanceInvoiceInputRequest) SetInvoiceInputDTO1(_invoiceInputDTO1 []InvoiceInputDto) error {
    r._invoiceInputDTO1 = _invoiceInputDTO1
    r.Set("invoice_input_d_t_o1", _invoiceInputDTO1)
    return nil
}

// InvoiceInputDTO1 Getter
func (r TmallTxcsFinanceInvoiceInputRequest) GetInvoiceInputDTO1() []InvoiceInputDto {
    return r._invoiceInputDTO1
}
