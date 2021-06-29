package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家提交设备信息 API请求
alibaba.aliqin.fc.iot.device.post

物联网商家设备信息录入
*/
type AlibabaAliqinFcIotDevicePostRequest struct {
    model.Params
    // 15位imei号
    imei   string
    // 设备类型（将来扩展）
    deviceType   string
    // 备注
    comments   string
    // 扩展字段
    midPatChannel   string
}

// 初始化AlibabaAliqinFcIotDevicePostRequest对象
func NewAlibabaAliqinFcIotDevicePostRequest() *AlibabaAliqinFcIotDevicePostRequest{
    return &AlibabaAliqinFcIotDevicePostRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotDevicePostRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.device.post"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotDevicePostRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Imei Setter
// 15位imei号
func (r *AlibabaAliqinFcIotDevicePostRequest) SetImei(imei string) error {
    r.imei = imei
    r.Set("imei", imei)
    return nil
}

// Imei Getter
func (r AlibabaAliqinFcIotDevicePostRequest) GetImei() string {
    return r.imei
}
// DeviceType Setter
// 设备类型（将来扩展）
func (r *AlibabaAliqinFcIotDevicePostRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaAliqinFcIotDevicePostRequest) GetDeviceType() string {
    return r.deviceType
}
// Comments Setter
// 备注
func (r *AlibabaAliqinFcIotDevicePostRequest) SetComments(comments string) error {
    r.comments = comments
    r.Set("comments", comments)
    return nil
}

// Comments Getter
func (r AlibabaAliqinFcIotDevicePostRequest) GetComments() string {
    return r.comments
}
// MidPatChannel Setter
// 扩展字段
func (r *AlibabaAliqinFcIotDevicePostRequest) SetMidPatChannel(midPatChannel string) error {
    r.midPatChannel = midPatChannel
    r.Set("mid_pat_channel", midPatChannel)
    return nil
}

// MidPatChannel Getter
func (r AlibabaAliqinFcIotDevicePostRequest) GetMidPatChannel() string {
    return r.midPatChannel
}
