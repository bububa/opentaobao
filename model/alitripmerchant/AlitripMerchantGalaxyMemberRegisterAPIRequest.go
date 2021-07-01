package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyMemberRegisterAPIRequest
星河-微信小程序会员注册 API请求
alitrip.merchant.galaxy.member.register

星河产品=微信小程序注册雅高会员服务 */
type AlitripMerchantGalaxyMemberRegisterAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 注册入参
	_memberParam *MemberParam
}

// New
