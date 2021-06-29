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
    orderId   int64
    // 物流信息
    logisticsList   []LogisticsDto
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
func (r *AlibabaIdleRentOrderSenditemRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleRentOrderSenditemRequest) GetOrderId() int64 {
    return r.orderId
}
// LogisticsList Setter
// 物流信息
func (r *AlibabaIdleRentOrderSenditemRequest) SetLogisticsList(logisticsList []LogisticsDto) error {
    r.logisticsList = logisticsList
    r.Set("logistics_list", logisticsList)
    return nil
}

// LogisticsList Getter
func (r AlibabaIdleRentOrderSenditemRequest) GetLogisticsList() []LogisticsDto {
    return r.logisticsList
}
