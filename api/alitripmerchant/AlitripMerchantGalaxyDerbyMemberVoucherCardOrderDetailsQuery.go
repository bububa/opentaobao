package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQuery 订单详情查询接口
// alitrip.merchant.galaxy.derby.member.voucher.card.order.details.query
//
// 订单详情查询接口
func AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQuery(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
