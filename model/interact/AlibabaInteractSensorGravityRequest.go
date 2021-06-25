package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
重力感应 APIRequest
alibaba.interact.sensor.gravity

native获取重力感应
*/
type AlibabaInteractSensorGravityRequest struct {
    model.Params

}

func NewAlibabaInteractSensorGravityRequest() *AlibabaInteractSensorGravityRequest{
    return &AlibabaInteractSensorGravityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorGravityRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.gravity"
}

func (r AlibabaInteractSensorGravityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


