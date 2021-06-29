package txcs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商发票录入 APIRequest
tmall.txcs.finance.invoice.input

提供天猫超市外部合作商家财务：供应商发票录入
*/
type TmallTxcsFinanceInvoiceInputRequest struct {
    model.Params

    // 门店ID
    ouCode   string 

    // 发票内容
    invoiceInputDTO1   []InvoiceInputDto 

}

func NewTmallTxcsFinanceInvoiceInputRequest() *TmallTxcsFinanceInvoiceInputRequest{
    return &TmallTxcsFinanceInvoiceInputRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTxcsFinanceInvoiceInputRequest) GetApiMethodName() string {
    return "tmall.txcs.finance.invoice.input"
}

func (r TmallTxcsFinanceInvoiceInputRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTxcsFinanceInvoiceInputRequest) SetOuCode(ouCode string) error {
    r.ouCode = ouCode
    r.Set("ou_code", ouCode)
    return nil
}

func (r TmallTxcsFinanceInvoiceInputRequest) GetOuCode() string {
    return r.ouCode
}

func (r *TmallTxcsFinanceInvoiceInputRequest) SetInvoiceInputDTO1(invoiceInputDTO1 []InvoiceInputDto) error {
    r.invoiceInputDTO1 = invoiceInputDTO1
    r.Set("invoice_input_d_t_o1", invoiceInputDTO1)
    return nil
}

func (r TmallTxcsFinanceInvoiceInputRequest) GetInvoiceInputDTO1() []InvoiceInputDto {
    return r.invoiceInputDTO1
}

