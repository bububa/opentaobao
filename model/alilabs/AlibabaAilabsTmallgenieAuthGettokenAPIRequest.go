package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthGettokenAPIRequest
设备授权 API请求
alibaba.ailabs.tmallgenie.auth.gettoken

获取设备授权码 */
type AlibabaAilabsTmallgenieAuthGettokenAPIRequest struct {
	model.Params
	// clientId
	_clientId string
	// 授权码
	_authCode string
	// 授权类型，只支持authorization_code
	_grantType string
}

// New
