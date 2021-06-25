package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘下单能力开放 APIRequest
alibaba.interact.sensor.trade.buy

交易流程鉴权
*/
type AlibabaInteractSensorTradeBuyRequest struct {
    model.Params

    // 系统自动生成
    id   string 

}

func NewAlibabaInteractSensorTradeBuyRequest() *AlibabaInteractSensorTradeBuyRequest{
    return &AlibabaInteractSensorTradeBuyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorTradeBuyRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.trade.buy"
}

func (r AlibabaInteractSensorTradeBuyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractSensorTradeBuyRequest) SetId(id string) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaInteractSensorTradeBuyRequest) GetId() string {
    return r.id
}

