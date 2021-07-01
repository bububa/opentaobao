package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

/* AlitripMerchantGalaxyOrderBook
星河-订单预订接口
alitrip.merchant.galaxy.order.book

为星河酒店解决方案的C端用户提供酒店预订能力 */
func AlitripMerchantGalaxyOrderBook(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderBookAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyOrderBookAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyOrderBookAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
