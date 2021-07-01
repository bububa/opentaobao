package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyMemberProviderRegisterAPIRequest
对外提供会员注册服务 API请求
alitrip.merchant.galaxy.member.provider.register

星河产品=对外提供注册雅高会员服务 */
type AlitripMerchantGalaxyMemberProviderRegisterAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 注册入参
	_providerMemberParam *ProviderMemberParam
}

// New
