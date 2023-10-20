package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderListQuery 星河-订单列表查询
// alitrip.merchant.galaxy.order.list.query
//
// 为C端用户提供酒店预订订单列表查询服务，包括订单支付状态、订单日期
func AlitripMerchantGalaxyOrderListQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderListQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyOrderListQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
