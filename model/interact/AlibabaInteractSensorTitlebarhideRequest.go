package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
隐藏titleBar APIRequest
alibaba.interact.sensor.titlebarhide

隐藏titleBar
*/
type AlibabaInteractSensorTitlebarhideRequest struct {
    model.Params

}

func NewAlibabaInteractSensorTitlebarhideRequest() *AlibabaInteractSensorTitlebarhideRequest{
    return &AlibabaInteractSensorTitlebarhideRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorTitlebarhideRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.titlebarhide"
}

func (r AlibabaInteractSensorTitlebarhideRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


