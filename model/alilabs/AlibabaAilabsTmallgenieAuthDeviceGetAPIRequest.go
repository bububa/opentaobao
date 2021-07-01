package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest
获取设备详情 API请求
alibaba.ailabs.tmallgenie.auth.device.get

通过此接口获取设备详情 */
type AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest struct {
	model.Params
	// 客户id
	_clientId string
	// 用户开放id
	_userOpenId string
	// 设备uuid
	_uuid string
}

// New
