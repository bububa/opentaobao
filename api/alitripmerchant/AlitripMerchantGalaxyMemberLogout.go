package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberLogout 星河-用户登出
// alitrip.merchant.galaxy.member.logout
//
// 星河=微信小程序用户登出
func AlitripMerchantGalaxyMemberLogout(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberLogoutAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberLogoutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
