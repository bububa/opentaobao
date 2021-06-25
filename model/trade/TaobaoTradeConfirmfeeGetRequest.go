package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取交易确认收货费用 APIRequest
taobao.trade.confirmfee.get

获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
*/
type TaobaoTradeConfirmfeeGetRequest struct {
    model.Params

    // 交易主订单或子订单ID
    tid   int64 

}

func NewTaobaoTradeConfirmfeeGetRequest() *TaobaoTradeConfirmfeeGetRequest{
    return &TaobaoTradeConfirmfeeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeConfirmfeeGetRequest) GetApiMethodName() string {
    return "taobao.trade.confirmfee.get"
}

func (r TaobaoTradeConfirmfeeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeConfirmfeeGetRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoTradeConfirmfeeGetRequest) GetTid() int64 {
    return r.tid
}

