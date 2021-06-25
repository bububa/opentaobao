package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
popwindow APIRequest
alibaba.interact.sensor.popwindow

popwindow
*/
type AlibabaInteractSensorPopwindowRequest struct {
    model.Params

}

func NewAlibabaInteractSensorPopwindowRequest() *AlibabaInteractSensorPopwindowRequest{
    return &AlibabaInteractSensorPopwindowRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorPopwindowRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.popwindow"
}

func (r AlibabaInteractSensorPopwindowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


