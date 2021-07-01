package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotDeviceIsexistAPIRequest
判断设备是否存在 API请求
alibaba.aliqin.fc.iot.device.isexist

判断设备是否存在 */
type AlibabaAliqinFcIotDeviceIsexistAPIRequest struct {
	model.Params
	// 设备编号
	_imei string
	// 设备类型（预留将来扩展）
	_deviceType string
	// 渠道扩展编码（预留）
	_midPatChannel string
}

// New
