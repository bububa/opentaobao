package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardActive 权益卡订单激活
// alitrip.merchant.galaxy.derby.member.voucher.card.active
//
// 权益卡订单激活
func AlitripMerchantGalaxyDerbyMemberVoucherCardActive(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
