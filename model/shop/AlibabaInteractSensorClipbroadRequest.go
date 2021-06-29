package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Weex页面设置或读取剪切板 API请求
alibaba.interact.sensor.clipbroad

Weex页面设置或读取剪切板
*/
type AlibabaInteractSensorClipbroadRequest struct {
    model.Params
    // 客户端鉴权使用，实际不会发送或接收数据
    unNamed   string
}

// 初始化AlibabaInteractSensorClipbroadRequest对象
func NewAlibabaInteractSensorClipbroadRequest() *AlibabaInteractSensorClipbroadRequest{
    return &AlibabaInteractSensorClipbroadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorClipbroadRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.clipbroad"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorClipbroadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnNamed Setter
// 客户端鉴权使用，实际不会发送或接收数据
func (r *AlibabaInteractSensorClipbroadRequest) SetUnNamed(unNamed string) error {
    r.unNamed = unNamed
    r.Set("un_named", unNamed)
    return nil
}

// UnNamed Getter
func (r AlibabaInteractSensorClipbroadRequest) GetUnNamed() string {
    return r.unNamed
}
