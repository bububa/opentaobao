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
    _orderId   int64
    // 物流信息
    _logistics   *LogisticsDTO
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
func (r *AlibabaIdleRentOrderPackageRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleRentOrderPackageRequest) GetOrderId() int64 {
    return r._orderId
}
// Logistics Setter
// 物流信息
func (r *AlibabaIdleRentOrderPackageRequest) SetLogistics(_logistics *LogisticsDTO) error {
    r._logistics = _logistics
    r.Set("logistics", _logistics)
    return nil
}

// Logistics Getter
func (r AlibabaIdleRentOrderPackageRequest) GetLogistics() *LogisticsDTO {
    return r._logistics
}
