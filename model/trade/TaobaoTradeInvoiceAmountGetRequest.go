package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单应开票金额 API请求
taobao.trade.invoice.amount.get

订单应开票金额计算
*/
type TaobaoTradeInvoiceAmountGetRequest struct {
    model.Params
    // 业务订单ID
    tid   int64
}

// 初始化TaobaoTradeInvoiceAmountGetRequest对象
func NewTaobaoTradeInvoiceAmountGetRequest() *TaobaoTradeInvoiceAmountGetRequest{
    return &TaobaoTradeInvoiceAmountGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeInvoiceAmountGetRequest) GetApiMethodName() string {
    return "taobao.trade.invoice.amount.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeInvoiceAmountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 业务订单ID
func (r *TaobaoTradeInvoiceAmountGetRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoTradeInvoiceAmountGetRequest) GetTid() int64 {
    return r.tid
}
