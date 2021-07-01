package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlinkDeviceBindAPIRequest
绑定设备 API请求
alibaba.alink.device.bind

阿里智能解绑设备 */
type AlibabaAlinkDeviceBindAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// New
