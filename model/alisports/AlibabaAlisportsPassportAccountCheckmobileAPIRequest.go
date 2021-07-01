package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsPassportAccountCheckmobileAPIRequest
阿里体育会员系统--手机号验证接口 API请求
alibaba.alisports.passport.account.checkmobile

验证三方用户的手机号，获取对应的阿里体育会员id */
type AlibabaAlisportsPassportAccountCheckmobileAPIRequest struct {
	model.Params
	// 业务appkey
	_alispAppKey string
	// 调用时间戳
	_alispTime string
	// 签名字符串
	_alispSign string
	// 合作方用户ID
	_appUid string
	// 用户呢称
	_nick string
	// 手机号
	_mobile string
	// 性别 0未设置 1男 2女 3保密
	_gender string
	// 生日
	_birthday string
}

// New
