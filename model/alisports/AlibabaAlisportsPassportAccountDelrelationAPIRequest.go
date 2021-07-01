package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsPassportAccountDelrelationAPIRequest
阿里体育会员系统--取消三方关联接口 API请求
alibaba.alisports.passport.account.delrelation

阿里体育会员系统--取消三方关联接口 */
type AlibabaAlisportsPassportAccountDelrelationAPIRequest struct {
	model.Params
	// 业务appkey
	_alispAppKey string
	// 调用时间戳
	_alispTime string
	// 签名字符串
	_alispSign string
	// 合作方用户ID
	_appUid string
	// 阿里体育会员id
	_aliuid string
}

// New
