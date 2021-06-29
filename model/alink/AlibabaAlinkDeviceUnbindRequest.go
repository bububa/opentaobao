package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解绑设备 API请求
alibaba.alink.device.unbind

阿里智能解绑设备
*/
type AlibabaAlinkDeviceUnbindRequest struct {
    model.Params
    // 设备id
    _uuid   string
}

// 初始化AlibabaAlinkDeviceUnbindRequest对象
func NewAlibabaAlinkDeviceUnbindRequest() *AlibabaAlinkDeviceUnbindRequest{
    return &AlibabaAlinkDeviceUnbindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceUnbindRequest) GetApiMethodName() string {
    return "alibaba.alink.device.unbind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceUnbindRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAlinkDeviceUnbindRequest) GetUuid() string {
    return r._uuid
}
