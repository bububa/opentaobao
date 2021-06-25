package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易组件 APIRequest
alibaba.interact.sensor.trade

交易流程
*/
type AlibabaInteractSensorTradeRequest struct {
    model.Params

    // 系统自动生成
    id   string 

}

func NewAlibabaInteractSensorTradeRequest() *AlibabaInteractSensorTradeRequest{
    return &AlibabaInteractSensorTradeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorTradeRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.trade"
}

func (r AlibabaInteractSensorTradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractSensorTradeRequest) SetId(id string) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaInteractSensorTradeRequest) GetId() string {
    return r.id
}

