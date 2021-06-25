package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
延长交易收货时间 APIRequest
taobao.trade.receivetime.delay

延长交易收货时间
*/
type TaobaoTradeReceivetimeDelayRequest struct {
    model.Params

    // 主订单号
    tid   int64 

    // 延长收货的天数，可选值为：3, 5, 7, 10。
    days   int64 

}

func NewTaobaoTradeReceivetimeDelayRequest() *TaobaoTradeReceivetimeDelayRequest{
    return &TaobaoTradeReceivetimeDelayRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeReceivetimeDelayRequest) GetApiMethodName() string {
    return "taobao.trade.receivetime.delay"
}

func (r TaobaoTradeReceivetimeDelayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeReceivetimeDelayRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoTradeReceivetimeDelayRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoTradeReceivetimeDelayRequest) SetDays(days int64) error {
    r.days = days
    r.Set("days", days)
    return nil
}

func (r TaobaoTradeReceivetimeDelayRequest) GetDays() int64 {
    return r.days
}

