package healthnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康新零售物流详情接口 APIRequest
alibaba.health.nr.logistics.query

对阿里健康o2o对接的商户提供查询物流单详情的能力
*/
type AlibabaHealthNrLogisticsQueryRequest struct {
    model.Params

    // 订单id
    orderId   int64 

}

func NewAlibabaHealthNrLogisticsQueryRequest() *AlibabaHealthNrLogisticsQueryRequest{
    return &AlibabaHealthNrLogisticsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthNrLogisticsQueryRequest) GetApiMethodName() string {
    return "alibaba.health.nr.logistics.query"
}

func (r AlibabaHealthNrLogisticsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthNrLogisticsQueryRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHealthNrLogisticsQueryRequest) GetOrderId() int64 {
    return r.orderId
}

