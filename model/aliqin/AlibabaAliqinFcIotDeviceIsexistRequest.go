package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断设备是否存在 API请求
alibaba.aliqin.fc.iot.device.isexist

判断设备是否存在
*/
type AlibabaAliqinFcIotDeviceIsexistAPIRequest struct {
    model.Params
    // 设备编号
    _imei   string
    // 设备类型（预留将来扩展）
    _deviceType   string
    // 渠道扩展编码（预留）
    _midPatChannel   string
}

// 初始化AlibabaAliqinFcIotDeviceIsexistAPIRequest对象
func NewAlibabaAliqinFcIotDeviceIsexistRequest() *AlibabaAliqinFcIotDeviceIsexistAPIRequest{
    return &AlibabaAliqinFcIotDeviceIsexistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotDeviceIsexistAPIRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.device.isexist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotDeviceIsexistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Imei Setter
// 设备编号
func (r *AlibabaAliqinFcIotDeviceIsexistAPIRequest) SetImei(_imei string) error {
    r._imei = _imei
    r.Set("imei", _imei)
    return nil
}

// Imei Getter
func (r AlibabaAliqinFcIotDeviceIsexistAPIRequest) GetImei() string {
    return r._imei
}
// DeviceType Setter
// 设备类型（预留将来扩展）
func (r *AlibabaAliqinFcIotDeviceIsexistAPIRequest) SetDeviceType(_deviceType string) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaAliqinFcIotDeviceIsexistAPIRequest) GetDeviceType() string {
    return r._deviceType
}
// MidPatChannel Setter
// 渠道扩展编码（预留）
func (r *AlibabaAliqinFcIotDeviceIsexistAPIRequest) SetMidPatChannel(_midPatChannel string) error {
    r._midPatChannel = _midPatChannel
    r.Set("mid_pat_channel", _midPatChannel)
    return nil
}

// MidPatChannel Getter
func (r AlibabaAliqinFcIotDeviceIsexistAPIRequest) GetMidPatChannel() string {
    return r._midPatChannel
}
