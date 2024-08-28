package auction

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/auction"
)

// TaobaoAuctionGovGetLatestbid 获取司法拍卖最新出价数据
// taobao.auction.gov.get.latestbid
//
// 获取司法拍卖最新出价数据
func TaobaoAuctionGovGetLatestbid(ctx context.Context, clt *core.SDKClient, req *auction.TaobaoAuctionGovGetLatestbidAPIRequest, resp *auction.TaobaoAuctionGovGetLatestbidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
