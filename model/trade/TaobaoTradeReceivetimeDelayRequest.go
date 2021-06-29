package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
延长交易收货时间 API请求
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

// 初始化TaobaoTradeReceivetimeDelayRequest对象
func NewTaobaoTradeReceivetimeDelayRequest() *TaobaoTradeReceivetimeDelayRequest{
    return &TaobaoTradeReceivetimeDelayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeReceivetimeDelayRequest) GetApiMethodName() string {
    return "taobao.trade.receivetime.delay"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeReceivetimeDelayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 主订单号
func (r *TaobaoTradeReceivetimeDelayRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoTradeReceivetimeDelayRequest) GetTid() int64 {
    return r.tid
}
// Days Setter
// 延长收货的天数，可选值为：3, 5, 7, 10。
func (r *TaobaoTradeReceivetimeDelayRequest) SetDays(days int64) error {
    r.days = days
    r.Set("days", days)
    return nil
}

// Days Getter
func (r TaobaoTradeReceivetimeDelayRequest) GetDays() int64 {
    return r.days
}
