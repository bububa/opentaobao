package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlinkDeviceUnbindAPIRequest
解绑设备 API请求
alibaba.alink.device.unbind

阿里智能解绑设备 */
type AlibabaAlinkDeviceUnbindAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// New
