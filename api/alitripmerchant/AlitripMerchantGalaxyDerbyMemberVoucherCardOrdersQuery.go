package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQuery 查询权益卡订单列表
// alitrip.merchant.galaxy.derby.member.voucher.card.orders.query
//
// 查询权益卡订单列表
func AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
