package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详情 API请求
alibaba.alink.device.detail.get

阿里智能获取设备详情
*/
type AlibabaAlinkDeviceDetailGetRequest struct {
    model.Params
    // 设备id
    _uuid   string
}

// 初始化AlibabaAlinkDeviceDetailGetRequest对象
func NewAlibabaAlinkDeviceDetailGetRequest() *AlibabaAlinkDeviceDetailGetRequest{
    return &AlibabaAlinkDeviceDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceDetailGetRequest) GetApiMethodName() string {
    return "alibaba.alink.device.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceDetailGetRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAlinkDeviceDetailGetRequest) GetUuid() string {
    return r._uuid
}
