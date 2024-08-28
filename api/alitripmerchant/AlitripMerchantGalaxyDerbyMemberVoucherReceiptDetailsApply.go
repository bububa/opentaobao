package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApply v5.0付费会员卡开发订单开票详情申请
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.apply
//
// v5.0付费会员卡开发订单开票详情申请
func AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApply(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
