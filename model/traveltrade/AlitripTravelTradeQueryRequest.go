package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单详情查询接口 API请求
alitrip.travel.trade.query

飞猪度假订单详情查询接口
*/
type AlitripTravelTradeQueryAPIRequest struct {
    model.Params
    // 主订单id
    _orderId   int64
}

// 初始化AlitripTravelTradeQueryAPIRequest对象
func NewAlitripTravelTradeQueryRequest() *AlitripTravelTradeQueryAPIRequest{
    return &AlitripTravelTradeQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelTradeQueryAPIRequest) GetApiMethodName() string {
    return "alitrip.travel.trade.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelTradeQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 主订单id
func (r *AlitripTravelTradeQueryAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlitripTravelTradeQueryAPIRequest) GetOrderId() int64 {
    return r._orderId
}
