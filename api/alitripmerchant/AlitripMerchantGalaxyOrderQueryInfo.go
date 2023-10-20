package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderQueryInfo 订单详情改版
// alitrip.merchant.galaxy.order.query.info
//
// 订单页详情查询
func AlitripMerchantGalaxyOrderQueryInfo(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderQueryInfoAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyOrderQueryInfoAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyOrderQueryInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
