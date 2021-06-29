package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取欢行统一订单模型 API请求
alibaba.happytrip.order.get

通过订单id获取欢行统一订单模型数据
*/
type AlibabaHappytripOrderGetRequest struct {
    model.Params
    // 订单id
    _orderId   int64
}

// 初始化AlibabaHappytripOrderGetRequest对象
func NewAlibabaHappytripOrderGetRequest() *AlibabaHappytripOrderGetRequest{
    return &AlibabaHappytripOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripOrderGetRequest) GetApiMethodName() string {
    return "alibaba.happytrip.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaHappytripOrderGetRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripOrderGetRequest) GetOrderId() int64 {
    return r._orderId
}
