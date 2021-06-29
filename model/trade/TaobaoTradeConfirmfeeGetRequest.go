package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取交易确认收货费用 API请求
taobao.trade.confirmfee.get

获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
*/
type TaobaoTradeConfirmfeeGetRequest struct {
    model.Params
    // 交易主订单或子订单ID
    _tid   int64
}

// 初始化TaobaoTradeConfirmfeeGetRequest对象
func NewTaobaoTradeConfirmfeeGetRequest() *TaobaoTradeConfirmfeeGetRequest{
    return &TaobaoTradeConfirmfeeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeConfirmfeeGetRequest) GetApiMethodName() string {
    return "taobao.trade.confirmfee.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeConfirmfeeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 交易主订单或子订单ID
func (r *TaobaoTradeConfirmfeeGetRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoTradeConfirmfeeGetRequest) GetTid() int64 {
    return r._tid
}
