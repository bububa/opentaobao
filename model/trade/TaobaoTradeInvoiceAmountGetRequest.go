package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单应开票金额 APIRequest
taobao.trade.invoice.amount.get

订单应开票金额计算
*/
type TaobaoTradeInvoiceAmountGetRequest struct {
    model.Params

    // 业务订单ID
    tid   int64 

}

func NewTaobaoTradeInvoiceAmountGetRequest() *TaobaoTradeInvoiceAmountGetRequest{
    return &TaobaoTradeInvoiceAmountGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeInvoiceAmountGetRequest) GetApiMethodName() string {
    return "taobao.trade.invoice.amount.get"
}

func (r TaobaoTradeInvoiceAmountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeInvoiceAmountGetRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoTradeInvoiceAmountGetRequest) GetTid() int64 {
    return r.tid
}

