package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlinkDeviceInfoUpdateAPIRequest
更新设备昵称等信息 API请求
alibaba.alink.device.info.update

更新设备昵称等信息 */
type AlibabaAlinkDeviceInfoUpdateAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
	// 设备昵称
	_nickName string
}

// New
