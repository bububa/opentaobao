package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderBook 星河-订单预订接口
// alitrip.merchant.galaxy.order.book
//
// 为星河酒店解决方案的C端用户提供酒店预订能力
func AlitripMerchantGalaxyOrderBook(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderBookAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyOrderBookAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
