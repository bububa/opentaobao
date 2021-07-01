package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest
解绑设备 API请求
alibaba.ailabs.tmallgenie.auth.device.unbind

通过此接口解绑天猫精灵设备 */
type AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest struct {
	model.Params
	// 客户id
	_clientId string
	// 用户开放id
	_userOpenId string
	// 设备uuid
	_uuid string
}

// New
