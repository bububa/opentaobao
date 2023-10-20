package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyorderbook 星河-订单预订接口
// alitrip.merchant.galaxy.order.book
//
// 为星河酒店解决方案的C端用户提供酒店预订能力
func Alitripmerchantgalaxyorderbook(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyorderbookAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyorderbookAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyorderbookAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
