package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceListAPIRequest
获取用户设备列表 API请求
alibaba.ailabs.tmallgenie.auth.device.list

通过此接口获取用户绑定的设备信息列表 */
type AlibabaAilabsTmallgenieAuthDeviceListAPIRequest struct {
	model.Params
	// 客户id
	_clientId string
	// 用户开放id
	_userOpenId string
}

// New
