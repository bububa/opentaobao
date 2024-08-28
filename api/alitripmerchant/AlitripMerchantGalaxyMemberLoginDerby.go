package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberLoginDerby 小程序通过德比登入（会员认证）
// alitrip.merchant.galaxy.member.login.derby
//
// 会员认证(德比切换接口，包含授权接口)   返回认证及授权状态
func AlitripMerchantGalaxyMemberLoginDerby(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberLoginDerbyAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberLoginDerbyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
