package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认揽收商品 APIRequest
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

func NewAlibabaIdleRentOrderPackageRequest() *AlibabaIdleRentOrderPackageRequest{
    return &AlibabaIdleRentOrderPackageRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRentOrderPackageRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.package"
}

func (r AlibabaIdleRentOrderPackageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRentOrderPackageRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaIdleRentOrderPackageRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *AlibabaIdleRentOrderPackageRequest) SetLogistics(logistics *LogisticsDto) error {
    r.logistics = logistics
    r.Set("logistics", logistics)
    return nil
}

func (r AlibabaIdleRentOrderPackageRequest) GetLogistics() *LogisticsDto {
    return r.logistics
}

