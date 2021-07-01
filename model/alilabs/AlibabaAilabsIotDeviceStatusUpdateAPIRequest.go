package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsIotDeviceStatusUpdateAPIRequest
ailabs iot 设备状态更新 API请求
alibaba.ailabs.iot.device.status.update

用于人工智能实验室IoT合作厂商上报三方接入IoT设备状态更新时的设备状态上报 */
type AlibabaAilabsIotDeviceStatusUpdateAPIRequest struct {
	model.Params
	// 入参设备信息
	_deviceStatusDTO *DeviceStatusDto
}

// New
