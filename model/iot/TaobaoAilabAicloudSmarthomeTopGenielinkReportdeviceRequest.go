package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零配方案上报设备 APIRequest
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

func NewTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest() *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest{
    return &TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.smarthome.top.genielink.reportdevice"
}

func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) SetVendorId(vendorId int64) error {
    r.vendorId = vendorId
    r.Set("vendor_id", vendorId)
    return nil
}

func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) GetVendorId() int64 {
    return r.vendorId
}

func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) SetExtensions(extensions string) error {
    r.extensions = extensions
    r.Set("extensions", extensions)
    return nil
}

func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest) GetExtensions() string {
    return r.extensions
}

