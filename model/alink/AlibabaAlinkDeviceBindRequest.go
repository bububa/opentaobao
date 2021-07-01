package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
绑定设备 API请求
alibaba.alink.device.bind

阿里智能解绑设备
*/
type AlibabaAlinkDeviceBindAPIRequest struct {
    model.Params
    // 设备id
    _uuid   string
}

// 初始化AlibabaAlinkDeviceBindAPIRequest对象
func NewAlibabaAlinkDeviceBindRequest() *AlibabaAlinkDeviceBindAPIRequest{
    return &AlibabaAlinkDeviceBindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceBindAPIRequest) GetApiMethodName() string {
    return "alibaba.alink.device.bind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceBindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceBindAPIRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAlinkDeviceBindAPIRequest) GetUuid() string {
    return r._uuid
}
