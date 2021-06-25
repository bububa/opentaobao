package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单查询接口 APIRequest
alibaba.health.nr.logistics.waybill.get

商家登录后根据订单号查询物流单号及电子面单信息
*/
type AlibabaHealthNrLogisticsWaybillGetRequest struct {
    model.Params

    // 订单id
    orderId   int64 

}

func NewAlibabaHealthNrLogisticsWaybillGetRequest() *AlibabaHealthNrLogisticsWaybillGetRequest{
    return &AlibabaHealthNrLogisticsWaybillGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthNrLogisticsWaybillGetRequest) GetApiMethodName() string {
    return "alibaba.health.nr.logistics.waybill.get"
}

func (r AlibabaHealthNrLogisticsWaybillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthNrLogisticsWaybillGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHealthNrLogisticsWaybillGetRequest) GetOrderId() int64 {
    return r.orderId
}

