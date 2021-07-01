package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthGetcodeAPIRequest
获取token API请求
alibaba.ailabs.tmallgenie.auth.getcode

获取天猫精灵authCode */
type AlibabaAilabsTmallgenieAuthGetcodeAPIRequest struct {
	model.Params
	// 授权参数
	_authParam *TopAuthReqDto
	// 设备参数
	_deviceParam *TopDeviceReqDto
}

// New
