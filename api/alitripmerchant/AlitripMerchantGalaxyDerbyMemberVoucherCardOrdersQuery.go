package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQuery 查询权益卡订单列表
// alitrip.merchant.galaxy.derby.member.voucher.card.orders.query
//
// 查询权益卡订单列表
func AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQuery(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
