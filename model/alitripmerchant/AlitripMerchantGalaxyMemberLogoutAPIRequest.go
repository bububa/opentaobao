package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyMemberLogoutAPIRequest
星河-用户登出 API请求
alitrip.merchant.galaxy.member.logout

星河=微信小程序用户登出 */
type AlitripMerchantGalaxyMemberLogoutAPIRequest struct {
	model.Params
	// 租户信息
	_tenantKey string
	// 用户登录token
	_token string
}

// New
