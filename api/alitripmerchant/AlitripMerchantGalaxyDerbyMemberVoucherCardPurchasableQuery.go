package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQuery 德比付费会员卡可购查询
// alitrip.merchant.galaxy.derby.member.voucher.card.purchasable.query
//
// 德比付费会员卡可购查询
func AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQuery(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
