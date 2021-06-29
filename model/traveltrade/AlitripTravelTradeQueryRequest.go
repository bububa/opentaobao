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
type AlitripTravelTradeQueryRequest struct {
    model.Params
    // 主订单id
    orderId   int64
}

// 初始化AlitripTravelTradeQueryRequest对象
func NewAlitripTravelTradeQueryRequest() *AlitripTravelTradeQueryRequest{
    return &AlitripTravelTradeQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelTradeQueryRequest) GetApiMethodName() string {
    return "alitrip.travel.trade.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelTradeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 主订单id
func (r *AlitripTravelTradeQueryRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlitripTravelTradeQueryRequest) GetOrderId() int64 {
    return r.orderId
}
