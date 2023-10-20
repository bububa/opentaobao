package auction

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/auction"
)

// TaobaoAuctionGovDataTopnGet 根据不同维度，获取排行榜列表
// taobao.auction.gov.data.topn.get
//
// 根据不同时间维度(周,月,年)，获取(成交额或发拍件数)排行榜列表
func TaobaoAuctionGovDataTopnGet(clt *core.SDKClient, req *auction.TaobaoAuctionGovDataTopnGetAPIRequest, resp *auction.TaobaoAuctionGovDataTopnGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
