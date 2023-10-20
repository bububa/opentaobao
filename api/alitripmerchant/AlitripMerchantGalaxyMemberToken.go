package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberToken 星河-校验token
// alitrip.merchant.galaxy.member.token
//
// 校验或者刷新token
func AlitripMerchantGalaxyMemberToken(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberTokenAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberTokenAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
