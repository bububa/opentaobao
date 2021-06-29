package txcs

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/txcs"
)

/* 
供应商发票录入 
tmall.txcs.finance.invoice.input

提供天猫超市外部合作商家财务：供应商发票录入
*/
func TmallTxcsFinanceInvoiceInput(clt *core.SDKClient, req *txcs.TmallTxcsFinanceInvoiceInputRequest, session string) (*txcs.TmallTxcsFinanceInvoiceInputAPIResponse, error) {
    var resp txcs.TmallTxcsFinanceInvoiceInputAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
