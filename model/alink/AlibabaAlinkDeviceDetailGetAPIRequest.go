package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlinkDeviceDetailGetAPIRequest
获取设备详情 API请求
alibaba.alink.device.detail.get

阿里智能获取设备详情 */
type AlibabaAlinkDeviceDetailGetAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// New
