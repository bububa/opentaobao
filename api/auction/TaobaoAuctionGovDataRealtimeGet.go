package auction

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/auction"
)

// TaobaoAuctionGovDataRealtimeGet 获取实时(今日)统计数据
// taobao.auction.gov.data.realtime.get
//
// 提供查询当日法院及下属法院的拍卖统计数据
func TaobaoAuctionGovDataRealtimeGet(ctx context.Context, clt *core.SDKClient, req *auction.TaobaoAuctionGovDataRealtimeGetAPIRequest, resp *auction.TaobaoAuctionGovDataRealtimeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
