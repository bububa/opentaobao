package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyMemberCardAPIRequest
星河-获取会员卡信息 API请求
alitrip.merchant.galaxy.member.card

星河=根据会员等级获取会员的权益 */
type AlitripMerchantGalaxyMemberCardAPIRequest struct {
	model.Params
	// 租户信息
	_tenantKey string
	// 飞猪等级
	_fliggyLevel string
	// 卡类型
	_cardType string
}

// New
