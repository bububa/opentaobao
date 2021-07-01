package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyMemberTokenAPIRequest
星河-校验token API请求
alitrip.merchant.galaxy.member.token

校验或者刷新token */
type AlitripMerchantGalaxyMemberTokenAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 访问携带toekn
	_token string
}

// New
