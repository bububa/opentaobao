package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberRegister 星河-微信小程序会员注册
// alitrip.merchant.galaxy.member.register
//
// 星河产品=微信小程序注册雅高会员服务
func AlitripMerchantGalaxyMemberRegister(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberRegisterAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberRegisterAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
