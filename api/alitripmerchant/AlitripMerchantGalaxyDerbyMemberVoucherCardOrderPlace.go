package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlace 德比付费会员卡下单
// alitrip.merchant.galaxy.derby.member.voucher.card.order.place
//
// 德比付费会员卡下单
func AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlace(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
