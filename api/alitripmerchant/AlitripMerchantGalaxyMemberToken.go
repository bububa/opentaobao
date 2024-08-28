package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberToken 星河-校验token
// alitrip.merchant.galaxy.member.token
//
// 校验或者刷新token
func AlitripMerchantGalaxyMemberToken(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberTokenAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberTokenAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
