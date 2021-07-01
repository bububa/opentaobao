package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsPassportOauthAlipaygrantAPIRequest
阿里体育会员系统-支付宝授权接口 API请求
alibaba.alisports.passport.oauth.alipaygrant

开放给乐心运动使用的支付宝授权接口 */
type AlibabaAlisportsPassportOauthAlipaygrantAPIRequest struct {
	model.Params
	// 阿里体育分配的用户appkey
	_alispAppKey string
	// 请求接口的时间戳
	_alispTime string
	// 签名字符串
	_alispSign string
	// 调用支付宝登录sdk返回的code
	_authCode string
	// 固定为rich_sports
	_partnerMode string
	// 支付宝的appid
	_appid string
	// 合作方的userid，即用户唯一的id标识
	_appUid string
}

// New
