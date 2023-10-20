package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderQueryOrderCount 查询各种状态订单的总数
// alitrip.merchant.galaxy.order.query.order.count
//
// 调用查询接口整合各个订单类型总数
func AlitripMerchantGalaxyOrderQueryOrderCount(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
