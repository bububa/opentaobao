package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消叫车 API请求
alibaba.happytrip.taxi.order.cancel

取消叫车订单,行程中的订单不能取消
*/
type AlibabaHappytripTaxiOrderCancelRequest struct {
    model.Params
    // 订单id
    orderId   string
    // 是否强制取消(true 或 false)默认false
    force   string
    // 取消类型，0：系统取消，非0：用户取消
    type   int64
}

// 初始化AlibabaHappytripTaxiOrderCancelRequest对象
func NewAlibabaHappytripTaxiOrderCancelRequest() *AlibabaHappytripTaxiOrderCancelRequest{
    return &AlibabaHappytripTaxiOrderCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderCancelRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaHappytripTaxiOrderCancelRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiOrderCancelRequest) GetOrderId() string {
    return r.orderId
}
// Force Setter
// 是否强制取消(true 或 false)默认false
func (r *AlibabaHappytripTaxiOrderCancelRequest) SetForce(force string) error {
    r.force = force
    r.Set("force", force)
    return nil
}

// Force Getter
func (r AlibabaHappytripTaxiOrderCancelRequest) GetForce() string {
    return r.force
}
// Type Setter
// 取消类型，0：系统取消，非0：用户取消
func (r *AlibabaHappytripTaxiOrderCancelRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaHappytripTaxiOrderCancelRequest) GetType() int64 {
    return r.type
}
