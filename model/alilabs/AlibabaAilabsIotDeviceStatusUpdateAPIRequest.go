package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotDeviceStatusUpdateAPIRequest ailabs iot 设备状态更新 API请求
// alibaba.ailabs.iot.device.status.update
//
// 用于人工智能实验室IoT合作厂商上报三方接入IoT设备状态更新时的设备状态上报
type AlibabaAilabsIotDeviceStatusUpdateAPIRequest struct {
	model.Params
	// 入参设备信息
	_deviceStatusDTO *DeviceStatusDto
}

// NewAlibabaAilabsIotDeviceStatusUpdateRequest 初始化AlibabaAilabsIotDeviceStatusUpdateAPIRequest对象
func NewAlibabaAilabsIotDeviceStatusUpdateRequest() *AlibabaAilabsIotDeviceStatusUpdateAPIRequest {
	return &AlibabaAilabsIotDeviceStatusUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotDeviceStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.device.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotDeviceStatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsIotDeviceStatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceStatusDTO is DeviceStatusDTO Setter
// 入参设备信息
func (r *AlibabaAilabsIotDeviceStatusUpdateAPIRequest) SetDeviceStatusDTO(_deviceStatusDTO *DeviceStatusDto) error {
	r._deviceStatusDTO = _deviceStatusDTO
	r.Set("device_status_d_t_o", _deviceStatusDTO)
	return nil
}

// GetDeviceStatusDTO DeviceStatusDTO Getter
func (r AlibabaAilabsIotDeviceStatusUpdateAPIRequest) GetDeviceStatusDTO() *DeviceStatusDto {
	return r._deviceStatusDTO
}
