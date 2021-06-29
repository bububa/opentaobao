package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零配方案上报设备 API请求
taobao.ailab.aicloud.smarthome.top.genielink.reportdevice

零配方案中设备联网成功之后上报设备
*/
type TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest struct {
    model.Params
    // 供应商id
    vendorId   int64
    // 设备id
    deviceId   string
    // 设备状态，online上线，offline下线
    status   string
    // 保留字段json字符串
    extensions   string
}

// 初始化TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest对象
func NewTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest() *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest{
    return &TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.smarthome.top.genielink.reportdevice"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VendorId Setter
// 供应商id
func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) SetVendorId(vendorId int64) error {
    r.vendorId = vendorId
    r.Set("vendor_id", vendorId)
    return nil
}

// VendorId Getter
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) GetVendorId() int64 {
    return r.vendorId
}
// DeviceId Setter
// 设备id
func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) GetDeviceId() string {
    return r.deviceId
}
// Status Setter
// 设备状态，online上线，offline下线
func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) GetStatus() string {
    return r.status
}
// Extensions Setter
// 保留字段json字符串
func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) SetExtensions(extensions string) error {
    r.extensions = extensions
    r.Set("extensions", extensions)
    return nil
}

// Extensions Getter
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) GetExtensions() string {
    return r.extensions
}
