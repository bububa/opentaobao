package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyorderquery 星河-单个订单详细信息查询
// alitrip.merchant.galaxy.order.query
//
// 为用户提供酒店订单的详细信息查询能力
func Alitripmerchantgalaxyorderquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyorderqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyorderqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
