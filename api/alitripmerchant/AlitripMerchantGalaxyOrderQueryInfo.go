package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyorderqueryinfo 订单详情改版
// alitrip.merchant.galaxy.order.query.info
//
// 订单页详情查询
func Alitripmerchantgalaxyorderqueryinfo(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyorderqueryinfoAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyorderqueryinfoAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyorderqueryinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
