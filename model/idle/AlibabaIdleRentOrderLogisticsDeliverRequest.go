package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建揽收物流 API请求
alibaba.idle.rent.order.logistics.deliver

创建揽收物流
商家去物流公司创建物流订单
*/
type AlibabaIdleRentOrderLogisticsDeliverRequest struct {
    model.Params
    // 订单id
    _orderId   int64
    // 物流信息
    _logistics   *LogisticsDTO
}

// 初始化AlibabaIdleRentOrderLogisticsDeliverRequest对象
func NewAlibabaIdleRentOrderLogisticsDeliverRequest() *AlibabaIdleRentOrderLogisticsDeliverRequest{
    return &AlibabaIdleRentOrderLogisticsDeliverRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderLogisticsDeliverRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.logistics.deliver"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderLogisticsDeliverRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderLogisticsDeliverRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleRentOrderLogisticsDeliverRequest) GetOrderId() int64 {
    return r._orderId
}
// Logistics Setter
// 物流信息
func (r *AlibabaIdleRentOrderLogisticsDeliverRequest) SetLogistics(_logistics *LogisticsDTO) error {
    r._logistics = _logistics
    r.Set("logistics", _logistics)
    return nil
}

// Logistics Getter
func (r AlibabaIdleRentOrderLogisticsDeliverRequest) GetLogistics() *LogisticsDTO {
    return r._logistics
}
