package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsPassportOauthTokenvalidateAPIRequest
阿里体育会员系统--获取登录信息接口 API请求
alibaba.alisports.passport.oauth.tokenvalidate

阿里体育会员系统--获取登录信息接口 */
type AlibabaAlisportsPassportOauthTokenvalidateAPIRequest struct {
	model.Params
	// 登录时返回给前端的token
	_token string
	// 时间戳
	_alispTime string
	// 应用的appkey
	_alispAppKey string
	// 参数加密之后的串
	_alispSign string
}

// New
