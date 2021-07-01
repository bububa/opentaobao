package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsPassportAccountBindthirdidAPIRequest
阿里体育三方ID绑定接口 API请求
alibaba.alisports.passport.account.bindthirdid

阿里体育三方ID绑定接口 */
type AlibabaAlisportsPassportAccountBindthirdidAPIRequest struct {
	model.Params
	// 业务方appkey
	_alispAppKey string
	// 时间戳精确到秒
	_alispTime string
	// 接口签名
	_alispSign string
	// 阿里体育用户ID
	_aliuid string
	// 三方ID
	_appUid string
	// 手机号
	_mobile string
}

// New
