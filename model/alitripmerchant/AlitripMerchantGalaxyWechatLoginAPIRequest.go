package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyWechatLoginAPIRequest
星河-用户使用微信登陆 API请求
alitrip.merchant.galaxy.wechat.login

星河产品=用户微信小程序登陆 */
type AlitripMerchantGalaxyWechatLoginAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 微信小程序登陆请求参数
	_loginParam *LoginParam
}

// New
