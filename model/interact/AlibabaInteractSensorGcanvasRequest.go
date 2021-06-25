package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
gcanvas APIRequest
alibaba.interact.sensor.gcanvas

gcanvas 功能
*/
type AlibabaInteractSensorGcanvasRequest struct {
    model.Params

}

func NewAlibabaInteractSensorGcanvasRequest() *AlibabaInteractSensorGcanvasRequest{
    return &AlibabaInteractSensorGcanvasRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorGcanvasRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.gcanvas"
}

func (r AlibabaInteractSensorGcanvasRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


