package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
吹气 APIRequest
alibaba.interact.sensor.blow

客户端吹气
*/
type AlibabaInteractSensorBlowRequest struct {
    model.Params

}

func NewAlibabaInteractSensorBlowRequest() *AlibabaInteractSensorBlowRequest{
    return &AlibabaInteractSensorBlowRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorBlowRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.blow"
}

func (r AlibabaInteractSensorBlowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


