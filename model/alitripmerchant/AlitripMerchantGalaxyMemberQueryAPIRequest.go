package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyMemberQueryAPIRequest
星河-获取登录用户的信息 API请求
alitrip.merchant.galaxy.member.query

获取登录用户的信息 */
type AlitripMerchantGalaxyMemberQueryAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// toekn
	_token string
}

// New
