package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
gmedia APIRequest
alibaba.interact.sensor.gmedia

媒体功能
*/
type AlibabaInteractSensorGmediaRequest struct {
    model.Params

}

func NewAlibabaInteractSensorGmediaRequest() *AlibabaInteractSensorGmediaRequest{
    return &AlibabaInteractSensorGmediaRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorGmediaRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.gmedia"
}

func (r AlibabaInteractSensorGmediaRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


