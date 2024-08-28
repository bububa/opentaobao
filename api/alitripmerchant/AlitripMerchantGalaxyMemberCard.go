package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberCard 星河-获取会员卡信息
// alitrip.merchant.galaxy.member.card
//
// 星河=根据会员等级获取会员的权益
func AlitripMerchantGalaxyMemberCard(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberCardAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberCardAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
