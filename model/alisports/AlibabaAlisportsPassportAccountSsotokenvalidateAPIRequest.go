package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest
sso_token验证 API请求
alibaba.alisports.passport.account.ssotokenvalidate

sso_token验证 */
type AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest struct {
	model.Params
	// sso_token
	_ssoToken string
	// 应用APPKEY
	_alispAppKey string
	// 当前时间戳[精确到秒，10位]
	_alispTime string
	// 签名
	_alispSign string
}

// New
