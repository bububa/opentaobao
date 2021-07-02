package admarket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/admarket"
)

// YunosAdmarketAdBid 广告竞价服务
// yunos.admarket.ad.bid
//
// 广告竞价服务
func YunosAdmarketAdBid(clt *core.SDKClient, req *admarket.YunosAdmarketAdBidAPIRequest, session string) (*admarket.YunosAdmarketAdBidAPIResponse, error) {
	var resp admarket.YunosAdmarketAdBidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
