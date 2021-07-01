package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsPassportAccountSsotokenrefreshAPIRequest
sso_token刷新 API请求
alibaba.alisports.passport.account.ssotokenrefresh

sso_token刷新 */
type AlibabaAlisportsPassportAccountSsotokenrefreshAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 当前时间戳[精确到秒,10位]
	_alispTime string
	// 签名
	_alispSign string
	// 登录成功返回的access_token(access_token是TOP保留关键字，只能改名，望谅解)
	_secret string
}

// New
