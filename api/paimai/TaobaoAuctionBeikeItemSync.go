package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoAuctionBeikeItemSync 贝壳商品同步接口
// taobao.auction.beike.item.sync
//
// 贝壳商品同步接口
func TaobaoAuctionBeikeItemSync(clt *core.SDKClient, req *paimai.TaobaoAuctionBeikeItemSyncAPIRequest, resp *paimai.TaobaoAuctionBeikeItemSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
