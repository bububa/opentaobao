package auction

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/auction"
)

// TaobaoAuctionGovGetLatestbid 获取司法拍卖最新出价数据
// taobao.auction.gov.get.latestbid
//
// 获取司法拍卖最新出价数据
func TaobaoAuctionGovGetLatestbid(clt *core.SDKClient, req *auction.TaobaoAuctionGovGetLatestbidAPIRequest, resp *auction.TaobaoAuctionGovGetLatestbidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
