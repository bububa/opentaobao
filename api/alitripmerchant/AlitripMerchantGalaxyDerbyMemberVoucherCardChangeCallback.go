package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallback v5.0德比付费会员卡通知
// alitrip.merchant.galaxy.derby.member.voucher.card.change.callback
//
// v5.0德比付费会员卡通知
func AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallback(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
