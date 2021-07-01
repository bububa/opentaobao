package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest
天猫精灵淘宝登录授权绑定接口 API请求
alibaba.ailabs.tmallgenie.auth.taobaoauth

厂商获取用户淘宝授权之后，通过此接口获取天猫精灵授权，并绑定一台设备 */
type AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest struct {
	model.Params
	// 授权信息
	_authParam *TopAuthReqDto
	// 设备信息
	_deviceParam *TopDeviceReqDto
}

// New
