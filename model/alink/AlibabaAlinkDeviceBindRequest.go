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
type AlibabaAlinkDeviceBindRequest struct {
    model.Params
    // 设备id
    uuid   string
}

// 初始化AlibabaAlinkDeviceBindRequest对象
func NewAlibabaAlinkDeviceBindRequest() *AlibabaAlinkDeviceBindRequest{
    return &AlibabaAlinkDeviceBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceBindRequest) GetApiMethodName() string {
    return "alibaba.alink.device.bind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceBindRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAlinkDeviceBindRequest) GetUuid() string {
    return r.uuid
}
