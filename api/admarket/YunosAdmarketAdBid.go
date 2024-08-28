package admarket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/admarket"
)

// YunosAdmarketAdBid 广告竞价服务
// yunos.admarket.ad.bid
//
// 广告竞价服务
func YunosAdmarketAdBid(ctx context.Context, clt *core.SDKClient, req *admarket.YunosAdmarketAdBidAPIRequest, resp *admarket.YunosAdmarketAdBidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
