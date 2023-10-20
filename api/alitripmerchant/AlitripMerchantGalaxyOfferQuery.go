package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyofferquery 星河-offer查询
// alitrip.merchant.galaxy.offer.query
//
// 根据offer的ID查询offer信息
func Alitripmerchantgalaxyofferquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyofferqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyofferqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyofferqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
