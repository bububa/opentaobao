package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Weex页面设置或读取剪切板 APIRequest
alibaba.interact.sensor.clipbroad

Weex页面设置或读取剪切板
*/
type AlibabaInteractSensorClipbroadRequest struct {
    model.Params

    // 客户端鉴权使用，实际不会发送或接收数据
    unNamed   string 

}

func NewAlibabaInteractSensorClipbroadRequest() *AlibabaInteractSensorClipbroadRequest{
    return &AlibabaInteractSensorClipbroadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorClipbroadRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.clipbroad"
}

func (r AlibabaInteractSensorClipbroadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractSensorClipbroadRequest) SetUnNamed(unNamed string) error {
    r.unNamed = unNamed
    r.Set("un_named", unNamed)
    return nil
}

func (r AlibabaInteractSensorClipbroadRequest) GetUnNamed() string {
    return r.unNamed
}

