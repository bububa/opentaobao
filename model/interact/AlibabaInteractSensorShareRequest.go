package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分享 APIRequest
alibaba.interact.sensor.share

客户端分享
*/
type AlibabaInteractSensorShareRequest struct {
    model.Params

}

func NewAlibabaInteractSensorShareRequest() *AlibabaInteractSensorShareRequest{
    return &AlibabaInteractSensorShareRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorShareRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.share"
}

func (r AlibabaInteractSensorShareRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


