package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistory 查询会员兑换臻享卡历史记录
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem.history
//
// 查询会员兑换臻享卡历史记录
func AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistory(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
