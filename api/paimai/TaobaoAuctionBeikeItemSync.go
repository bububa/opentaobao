package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// Taobaoauctionbeikeitemsync 贝壳商品同步接口
// taobao.auction.beike.item.sync
//
// 贝壳商品同步接口
func Taobaoauctionbeikeitemsync(clt *core.SDKClient, req *paimai.TaobaoauctionbeikeitemsyncAPIRequest, session string) (*paimai.TaobaoauctionbeikeitemsyncAPIResponse, error) {
	var resp paimai.TaobaoauctionbeikeitemsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
