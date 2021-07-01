package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsPassportAccountTokenvalidateAPIRequest
阿里体育会员系统帐号登录注册token验证接口 API请求
alibaba.alisports.passport.account.tokenvalidate

阿里体育会员系统帐号登录注册token验证接口 */
type AlibabaAlisportsPassportAccountTokenvalidateAPIRequest struct {
	model.Params
	// 业务方appkey
	_alispAppKey string
	// 签名
	_alispSign string
	// token
	_token string
	// 注册用户类型
	_userType int64
	// 时间戳
	_alispTime string
	// 一键登录参数
	_secret string
	// json字符串，传入扩展字段
	_extInfo string
	// 选填，调用百川登录接口的appkey，百川登录时，需要传此字段
	_mtopAppkey string
}

// New
