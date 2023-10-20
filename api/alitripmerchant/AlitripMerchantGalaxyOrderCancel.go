package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyordercancel 星河-取消预订
// alitrip.merchant.galaxy.order.cancel
//
// 雅高酒店用户使用该接口，取消酒店预订
func Alitripmerchantgalaxyordercancel(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyordercancelAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyordercancelAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
