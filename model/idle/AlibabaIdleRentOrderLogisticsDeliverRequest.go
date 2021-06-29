package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建揽收物流 APIRequest
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

func NewAlibabaIdleRentOrderLogisticsDeliverRequest() *AlibabaIdleRentOrderLogisticsDeliverRequest{
    return &AlibabaIdleRentOrderLogisticsDeliverRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRentOrderLogisticsDeliverRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.logistics.deliver"
}

func (r AlibabaIdleRentOrderLogisticsDeliverRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRentOrderLogisticsDeliverRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaIdleRentOrderLogisticsDeliverRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *AlibabaIdleRentOrderLogisticsDeliverRequest) SetLogistics(logistics *LogisticsDto) error {
    r.logistics = logistics
    r.Set("logistics", logistics)
    return nil
}

func (r AlibabaIdleRentOrderLogisticsDeliverRequest) GetLogistics() *LogisticsDto {
    return r.logistics
}

