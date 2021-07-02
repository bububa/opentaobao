package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberToken 星河-校验token
// alitrip.merchant.galaxy.member.token
//
// 校验或者刷新token
func AlitripMerchantGalaxyMemberToken(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberTokenAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyMemberTokenAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyMemberTokenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
