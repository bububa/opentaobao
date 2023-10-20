package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderQueryOrderCount 查询各种状态订单的总数
// alitrip.merchant.galaxy.order.query.order.count
//
// 调用查询接口整合各个订单类型总数
func AlitripMerchantGalaxyOrderQueryOrderCount(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
