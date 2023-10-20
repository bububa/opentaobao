package auction

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/auction"
)

// TaobaoAuctionGovGetLatestbid 获取司法拍卖最新出价数据
// taobao.auction.gov.get.latestbid
//
// 获取司法拍卖最新出价数据
func TaobaoAuctionGovGetLatestbid(clt *core.SDKClient, req *auction.TaobaoAuctionGovGetLatestbidAPIRequest, session string) (*auction.TaobaoAuctionGovGetLatestbidAPIResponse, error) {
	var resp auction.TaobaoAuctionGovGetLatestbidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
