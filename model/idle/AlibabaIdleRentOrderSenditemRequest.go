package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认发货 APIRequest
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

func NewAlibabaIdleRentOrderSenditemRequest() *AlibabaIdleRentOrderSenditemRequest{
    return &AlibabaIdleRentOrderSenditemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRentOrderSenditemRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.senditem"
}

func (r AlibabaIdleRentOrderSenditemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRentOrderSenditemRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaIdleRentOrderSenditemRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *AlibabaIdleRentOrderSenditemRequest) SetLogisticsList(logisticsList []LogisticsDto) error {
    r.logisticsList = logisticsList
    r.Set("logistics_list", logisticsList)
    return nil
}

func (r AlibabaIdleRentOrderSenditemRequest) GetLogisticsList() []LogisticsDto {
    return r.logisticsList
}

