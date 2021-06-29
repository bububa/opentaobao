package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认揽收商品 API请求
alibaba.idle.rent.order.package

确认揽收
*/
type AlibabaIdleRentOrderPackageRequest struct {
    model.Params
    // 订单id
    orderId   int64
    // 物流信息
    logistics   *LogisticsDto
}

// 初始化AlibabaIdleRentOrderPackageRequest对象
func NewAlibabaIdleRentOrderPackageRequest() *AlibabaIdleRentOrderPackageRequest{
    return &AlibabaIdleRentOrderPackageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderPackageRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.package"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderPackageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderPackageRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleRentOrderPackageRequest) GetOrderId() int64 {
    return r.orderId
}
// Logistics Setter
// 物流信息
func (r *AlibabaIdleRentOrderPackageRequest) SetLogistics(logistics *LogisticsDto) error {
    r.logistics = logistics
    r.Set("logistics", logistics)
    return nil
}

// Logistics Getter
func (r AlibabaIdleRentOrderPackageRequest) GetLogistics() *LogisticsDto {
    return r.logistics
}
