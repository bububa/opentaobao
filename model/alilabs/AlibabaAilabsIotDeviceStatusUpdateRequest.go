package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ailabs iot 设备状态更新 API请求
alibaba.ailabs.iot.device.status.update

用于人工智能实验室IoT合作厂商上报三方接入IoT设备状态更新时的设备状态上报
*/
type AlibabaAilabsIotDeviceStatusUpdateRequest struct {
    model.Params
    // 入参设备信息
    _deviceStatusDTO   *DeviceStatusDTO
}

// 初始化AlibabaAilabsIotDeviceStatusUpdateRequest对象
func NewAlibabaAilabsIotDeviceStatusUpdateRequest() *AlibabaAilabsIotDeviceStatusUpdateRequest{
    return &AlibabaAilabsIotDeviceStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotDeviceStatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.device.status.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotDeviceStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceStatusDTO Setter
// 入参设备信息
func (r *AlibabaAilabsIotDeviceStatusUpdateRequest) SetDeviceStatusDTO(_deviceStatusDTO *DeviceStatusDTO) error {
    r._deviceStatusDTO = _deviceStatusDTO
    r.Set("device_status_d_t_o", _deviceStatusDTO)
    return nil
}

// DeviceStatusDTO Getter
func (r AlibabaAilabsIotDeviceStatusUpdateRequest) GetDeviceStatusDTO() *DeviceStatusDTO {
    return r._deviceStatusDTO
}
