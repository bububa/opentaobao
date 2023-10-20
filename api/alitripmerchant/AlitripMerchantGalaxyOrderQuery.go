package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderQuery 星河-单个订单详细信息查询
// alitrip.merchant.galaxy.order.query
//
// 为用户提供酒店订单的详细信息查询能力
func AlitripMerchantGalaxyOrderQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
