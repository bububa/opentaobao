package auction

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/auction"
)

// Taobaoauctiongovdatatopnget 根据不同维度，获取排行榜列表
// taobao.auction.gov.data.topn.get
//
// 根据不同时间维度(周,月,年)，获取(成交额或发拍件数)排行榜列表
func Taobaoauctiongovdatatopnget(clt *core.SDKClient, req *auction.TaobaoauctiongovdatatopngetAPIRequest, session string) (*auction.TaobaoauctiongovdatatopngetAPIResponse, error) {
	var resp auction.TaobaoauctiongovdatatopngetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
