package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单 API请求
alibaba.idle.rent.order.query

查询订单信息
*/
type AlibabaIdleRentOrderQueryRequest struct {
    model.Params
    // 订单id
    _orderId   int64
}

// 初始化AlibabaIdleRentOrderQueryRequest对象
func NewAlibabaIdleRentOrderQueryRequest() *AlibabaIdleRentOrderQueryRequest{
    return &AlibabaIdleRentOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderQueryRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleRentOrderQueryRequest) GetOrderId() int64 {
    return r._orderId
}
