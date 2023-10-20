package auction

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/auction"
)

// Taobaoauctiongovgetlatestbid 获取司法拍卖最新出价数据
// taobao.auction.gov.get.latestbid
//
// 获取司法拍卖最新出价数据
func Taobaoauctiongovgetlatestbid(clt *core.SDKClient, req *auction.TaobaoauctiongovgetlatestbidAPIRequest, session string) (*auction.TaobaoauctiongovgetlatestbidAPIResponse, error) {
	var resp auction.TaobaoauctiongovgetlatestbidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
