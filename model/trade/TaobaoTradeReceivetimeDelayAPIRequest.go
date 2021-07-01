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
type TaobaoTradeReceivetimeDelayAPIRequest struct {
    model.Params
    // 主订单号
    _tid   int64
    // 延长收货的天数，可选值为：3, 5, 7, 10。
    _days   int64
}

// 初始化TaobaoTradeReceivetimeDelayAPIRequest对象
func NewTaobaoTradeReceivetimeDelayRequest() *TaobaoTradeReceivetimeDelayAPIRequest{
    return &TaobaoTradeReceivetimeDelayAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeReceivetimeDelayAPIRequest) GetApiMethodName() string {
    return "taobao.trade.receivetime.delay"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeReceivetimeDelayAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 主订单号
func (r *TaobaoTradeReceivetimeDelayAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoTradeReceivetimeDelayAPIRequest) GetTid() int64 {
    return r._tid
}
// Days Setter
// 延长收货的天数，可选值为：3, 5, 7, 10。
func (r *TaobaoTradeReceivetimeDelayAPIRequest) SetDays(_days int64) error {
    r._days = _days
    r.Set("days", _days)
    return nil
}

// Days Getter
func (r TaobaoTradeReceivetimeDelayAPIRequest) GetDays() int64 {
    return r._days
}
