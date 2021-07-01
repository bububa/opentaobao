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
type AlibabaInteractSensorUiAPIRequest struct {
    model.Params
    // 仅作客户端鉴权使用，不会发送接收请求
    _unNamed   string
}

// 初始化AlibabaInteractSensorUiAPIRequest对象
func NewAlibabaInteractSensorUiRequest() *AlibabaInteractSensorUiAPIRequest{
    return &AlibabaInteractSensorUiAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorUiAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.ui"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorUiAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnNamed Setter
// 仅作客户端鉴权使用，不会发送接收请求
func (r *AlibabaInteractSensorUiAPIRequest) SetUnNamed(_unNamed string) error {
    r._unNamed = _unNamed
    r.Set("un_named", _unNamed)
    return nil
}

// UnNamed Getter
func (r AlibabaInteractSensorUiAPIRequest) GetUnNamed() string {
    return r._unNamed
}
