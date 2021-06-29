package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认发货 API请求
alibaba.idle.rent.order.senditem

确认发货
*/
type AlibabaIdleRentOrderSenditemRequest struct {
    model.Params
    // 订单id
    _orderId   int64
    // 物流信息
    _logisticsList   []LogisticsDTO
}

// 初始化AlibabaIdleRentOrderSenditemRequest对象
func NewAlibabaIdleRentOrderSenditemRequest() *AlibabaIdleRentOrderSenditemRequest{
    return &AlibabaIdleRentOrderSenditemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderSenditemRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.senditem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderSenditemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderSenditemRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleRentOrderSenditemRequest) GetOrderId() int64 {
    return r._orderId
}
// LogisticsList Setter
// 物流信息
func (r *AlibabaIdleRentOrderSenditemRequest) SetLogisticsList(_logisticsList []LogisticsDTO) error {
    r._logisticsList = _logisticsList
    r.Set("logistics_list", _logisticsList)
    return nil
}

// LogisticsList Getter
func (r AlibabaIdleRentOrderSenditemRequest) GetLogisticsList() []LogisticsDTO {
    return r._logisticsList
}
