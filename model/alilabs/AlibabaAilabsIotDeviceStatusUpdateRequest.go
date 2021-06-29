package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ailabs iot 设备状态更新 APIRequest
alibaba.ailabs.iot.device.status.update

用于人工智能实验室IoT合作厂商上报三方接入IoT设备状态更新时的设备状态上报
*/
type AlibabaAilabsIotDeviceStatusUpdateRequest struct {
    model.Params

    // 入参设备信息
    deviceStatusDTO   *DeviceStatusDto 

}

func NewAlibabaAilabsIotDeviceStatusUpdateRequest() *AlibabaAilabsIotDeviceStatusUpdateRequest{
    return &AlibabaAilabsIotDeviceStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsIotDeviceStatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.device.status.update"
}

func (r AlibabaAilabsIotDeviceStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsIotDeviceStatusUpdateRequest) SetDeviceStatusDTO(deviceStatusDTO *DeviceStatusDto) error {
    r.deviceStatusDTO = deviceStatusDTO
    r.Set("device_status_d_t_o", deviceStatusDTO)
    return nil
}

func (r AlibabaAilabsIotDeviceStatusUpdateRequest) GetDeviceStatusDTO() *DeviceStatusDto {
    return r.deviceStatusDTO
}

