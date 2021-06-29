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
type AlibabaAliqinFcIotDeviceIsexistRequest struct {
    model.Params
    // 设备编号
    imei   string
    // 设备类型（预留将来扩展）
    deviceType   string
    // 渠道扩展编码（预留）
    midPatChannel   string
}

// 初始化AlibabaAliqinFcIotDeviceIsexistRequest对象
func NewAlibabaAliqinFcIotDeviceIsexistRequest() *AlibabaAliqinFcIotDeviceIsexistRequest{
    return &AlibabaAliqinFcIotDeviceIsexistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotDeviceIsexistRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.device.isexist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotDeviceIsexistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Imei Setter
// 设备编号
func (r *AlibabaAliqinFcIotDeviceIsexistRequest) SetImei(imei string) error {
    r.imei = imei
    r.Set("imei", imei)
    return nil
}

// Imei Getter
func (r AlibabaAliqinFcIotDeviceIsexistRequest) GetImei() string {
    return r.imei
}
// DeviceType Setter
// 设备类型（预留将来扩展）
func (r *AlibabaAliqinFcIotDeviceIsexistRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaAliqinFcIotDeviceIsexistRequest) GetDeviceType() string {
    return r.deviceType
}
// MidPatChannel Setter
// 渠道扩展编码（预留）
func (r *AlibabaAliqinFcIotDeviceIsexistRequest) SetMidPatChannel(midPatChannel string) error {
    r.midPatChannel = midPatChannel
    r.Set("mid_pat_channel", midPatChannel)
    return nil
}

// MidPatChannel Getter
func (r AlibabaAliqinFcIotDeviceIsexistRequest) GetMidPatChannel() string {
    return r.midPatChannel
}
