package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
摇一摇 APIRequest
alibaba.interact.sensor.shake

摇一摇
*/
type AlibabaInteractSensorShakeRequest struct {
    model.Params

}

func NewAlibabaInteractSensorShakeRequest() *AlibabaInteractSensorShakeRequest{
    return &AlibabaInteractSensorShakeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorShakeRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.shake"
}

func (r AlibabaInteractSensorShakeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


