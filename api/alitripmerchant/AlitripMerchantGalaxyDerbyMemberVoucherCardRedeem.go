package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeem 根据兑换码兑换臻享卡接口
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem
//
// 根据兑换码兑换臻享卡接口
func AlitripMerchantGalaxyDerbyMemberVoucherCardRedeem(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
