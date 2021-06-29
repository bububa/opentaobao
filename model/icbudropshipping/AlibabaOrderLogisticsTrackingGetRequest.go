package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴订单物流轨迹查询 APIRequest
alibaba.order.logistics.tracking.get

阿里巴巴订单物流轨迹查询
*/
type AlibabaOrderLogisticsTrackingGetRequest struct {
    model.Params

    // order id
    tradeId   int64 

}

func NewAlibabaOrderLogisticsTrackingGetRequest() *AlibabaOrderLogisticsTrackingGetRequest{
    return &AlibabaOrderLogisticsTrackingGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOrderLogisticsTrackingGetRequest) GetApiMethodName() string {
    return "alibaba.order.logistics.tracking.get"
}

func (r AlibabaOrderLogisticsTrackingGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaOrderLogisticsTrackingGetRequest) SetTradeId(tradeId int64) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

func (r AlibabaOrderLogisticsTrackingGetRequest) GetTradeId() int64 {
    return r.tradeId
}

