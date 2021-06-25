package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客户端震动 APIRequest
alibaba.interact.sensor.vibrate

客户端震动
*/
type AlibabaInteractSensorVibrateRequest struct {
    model.Params

}

func NewAlibabaInteractSensorVibrateRequest() *AlibabaInteractSensorVibrateRequest{
    return &AlibabaInteractSensorVibrateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorVibrateRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.vibrate"
}

func (r AlibabaInteractSensorVibrateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


