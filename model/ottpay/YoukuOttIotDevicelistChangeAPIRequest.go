package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttIotDevicelistChangeAPIRequest
iot设备列表变化接口 API请求
youku.ott.iot.devicelist.change

iot设备列表变化接口 */
type YoukuOttIotDevicelistChangeAPIRequest struct {
	model.Params
	// 变更信息
	_changeInfo string
}

// New
