package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客户端授权页 APIRequest
alibaba.interact.sensor.authorize

客户端授权页
*/
type AlibabaInteractSensorAuthorizeRequest struct {
    model.Params

}

func NewAlibabaInteractSensorAuthorizeRequest() *AlibabaInteractSensorAuthorizeRequest{
    return &AlibabaInteractSensorAuthorizeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorAuthorizeRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.authorize"
}

func (r AlibabaInteractSensorAuthorizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


