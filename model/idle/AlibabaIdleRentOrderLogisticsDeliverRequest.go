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
    orderId   int64
    // 物流信息
    logistics   *LogisticsDto
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
func (r *AlibabaIdleRentOrderLogisticsDeliverRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleRentOrderLogisticsDeliverRequest) GetOrderId() int64 {
    return r.orderId
}
// Logistics Setter
// 物流信息
func (r *AlibabaIdleRentOrderLogisticsDeliverRequest) SetLogistics(logistics *LogisticsDto) error {
    r.logistics = logistics
    r.Set("logistics", logistics)
    return nil
}

// Logistics Getter
func (r AlibabaIdleRentOrderLogisticsDeliverRequest) GetLogistics() *LogisticsDto {
    return r.logistics
}
