package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取登陆页面 APIRequest
alibaba.interact.sensor.login

获取登陆页面
*/
type AlibabaInteractSensorLoginRequest struct {
    model.Params

}

func NewAlibabaInteractSensorLoginRequest() *AlibabaInteractSensorLoginRequest{
    return &AlibabaInteractSensorLoginRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorLoginRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.login"
}

func (r AlibabaInteractSensorLoginRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


