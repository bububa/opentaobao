package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberProviderRegister 对外提供会员注册服务
// alitrip.merchant.galaxy.member.provider.register
//
// 星河产品=对外提供注册雅高会员服务
func AlitripMerchantGalaxyMemberProviderRegister(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberProviderRegisterAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberProviderRegisterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
