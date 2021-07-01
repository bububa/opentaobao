package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotDevicePostAPIRequest
商家提交设备信息 API请求
alibaba.aliqin.fc.iot.device.post

物联网商家设备信息录入 */
type AlibabaAliqinFcIotDevicePostAPIRequest struct {
	model.Params
	// 15位imei号
	_imei string
	// 设备类型（将来扩展）
	_deviceType string
	// 备注
	_comments string
	// 扩展字段
	_midPatChannel string
}

// New
