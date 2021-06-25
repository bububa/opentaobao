package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
基本ui操作 APIRequest
alibaba.interact.sensor.ui

Weex 基本UI操作
*/
type AlibabaInteractSensorUiRequest struct {
    model.Params

    // 仅作客户端鉴权使用，不会发送接收请求
    unNamed   string 

}

func NewAlibabaInteractSensorUiRequest() *AlibabaInteractSensorUiRequest{
    return &AlibabaInteractSensorUiRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorUiRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.ui"
}

func (r AlibabaInteractSensorUiRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractSensorUiRequest) SetUnNamed(unNamed string) error {
    r.unNamed = unNamed
    r.Set("un_named", unNamed)
    return nil
}

func (r AlibabaInteractSensorUiRequest) GetUnNamed() string {
    return r.unNamed
}

