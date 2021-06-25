package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
toast APIRequest
alibaba.interact.sensor.toast

toast提示
*/
type AlibabaInteractSensorToastRequest struct {
    model.Params

}

func NewAlibabaInteractSensorToastRequest() *AlibabaInteractSensorToastRequest{
    return &AlibabaInteractSensorToastRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorToastRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.toast"
}

func (r AlibabaInteractSensorToastRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


