package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsIotDeviceListGetAPIRequest
获取iot设备列表 API请求
alibaba.ailabs.iot.device.list.get

通过此接口获取用户名下的iot设备列表 */
type AlibabaAilabsIotDeviceListGetAPIRequest struct {
	model.Params
	// 用户open id
	_userOpenId string
	// client id
	_clientId string
}

// New
