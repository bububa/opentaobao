package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
基本ui操作 API请求
alibaba.interact.sensor.ui

Weex 基本UI操作
*/
type AlibabaInteractSensorUiRequest struct {
    model.Params
    // 仅作客户端鉴权使用，不会发送接收请求
    _unNamed   string
}

// 初始化AlibabaInteractSensorUiRequest对象
func NewAlibabaInteractSensorUiRequest() *AlibabaInteractSensorUiRequest{
    return &AlibabaInteractSensorUiRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorUiRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.ui"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorUiRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnNamed Setter
// 仅作客户端鉴权使用，不会发送接收请求
func (r *AlibabaInteractSensorUiRequest) SetUnNamed(_unNamed string) error {
    r._unNamed = _unNamed
    r.Set("un_named", _unNamed)
    return nil
}

// UnNamed Getter
func (r AlibabaInteractSensorUiRequest) GetUnNamed() string {
    return r._unNamed
}
