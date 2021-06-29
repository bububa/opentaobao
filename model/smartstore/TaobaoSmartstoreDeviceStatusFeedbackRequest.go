package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备在线状态回流 API请求
taobao.smartstore.device.status.feedback

智能硬件设备状态回流
*/
type TaobaoSmartstoreDeviceStatusFeedbackRequest struct {
    model.Params
    // ONLINE_WITH_CONTENT("ONLINE_WITH_CONTENT", "设备在线"), OFFLINE("OFFLINE", "设备断线");
    status   string
    // 设备编码
    deviceCode   string
    // 当前状态的时间
    statusTime   string
}

// 初始化TaobaoSmartstoreDeviceStatusFeedbackRequest对象
func NewTaobaoSmartstoreDeviceStatusFeedbackRequest() *TaobaoSmartstoreDeviceStatusFeedbackRequest{
    return &TaobaoSmartstoreDeviceStatusFeedbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSmartstoreDeviceStatusFeedbackRequest) GetApiMethodName() string {
    return "taobao.smartstore.device.status.feedback"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSmartstoreDeviceStatusFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// ONLINE_WITH_CONTENT("ONLINE_WITH_CONTENT", "设备在线"), OFFLINE("OFFLINE", "设备断线");
func (r *TaobaoSmartstoreDeviceStatusFeedbackRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoSmartstoreDeviceStatusFeedbackRequest) GetStatus() string {
    return r.status
}
// DeviceCode Setter
// 设备编码
func (r *TaobaoSmartstoreDeviceStatusFeedbackRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

// DeviceCode Getter
func (r TaobaoSmartstoreDeviceStatusFeedbackRequest) GetDeviceCode() string {
    return r.deviceCode
}
// StatusTime Setter
// 当前状态的时间
func (r *TaobaoSmartstoreDeviceStatusFeedbackRequest) SetStatusTime(statusTime string) error {
    r.statusTime = statusTime
    r.Set("status_time", statusTime)
    return nil
}

// StatusTime Getter
func (r TaobaoSmartstoreDeviceStatusFeedbackRequest) GetStatusTime() string {
    return r.statusTime
}
