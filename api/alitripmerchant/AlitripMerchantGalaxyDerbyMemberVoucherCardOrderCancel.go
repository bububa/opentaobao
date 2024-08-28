package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancel 取消订单
// alitrip.merchant.galaxy.derby.member.voucher.card.order.cancel
//
// 德比付费会员卡订单取消
func AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancel(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
