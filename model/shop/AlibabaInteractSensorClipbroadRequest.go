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
type AlibabaInteractSensorClipbroadAPIRequest struct {
    model.Params
    // 客户端鉴权使用，实际不会发送或接收数据
    _unNamed   string
}

// 初始化AlibabaInteractSensorClipbroadAPIRequest对象
func NewAlibabaInteractSensorClipbroadRequest() *AlibabaInteractSensorClipbroadAPIRequest{
    return &AlibabaInteractSensorClipbroadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorClipbroadAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.clipbroad"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorClipbroadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnNamed Setter
// 客户端鉴权使用，实际不会发送或接收数据
func (r *AlibabaInteractSensorClipbroadAPIRequest) SetUnNamed(_unNamed string) error {
    r._unNamed = _unNamed
    r.Set("un_named", _unNamed)
    return nil
}

// UnNamed Getter
func (r AlibabaInteractSensorClipbroadAPIRequest) GetUnNamed() string {
    return r._unNamed
}
