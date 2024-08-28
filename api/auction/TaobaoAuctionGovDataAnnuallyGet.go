package auction

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/auction"
)

// TaobaoAuctionGovDataAnnuallyGet 按年统计法院拍卖数据
// taobao.auction.gov.data.annually.get
//
// 按月统计法院拍卖数据 包含：
// 标的件数统计：发布标的件数、结束标的件数、开拍标的件数
// 竞价实况：预计成交金额、出价次数、报名人数
// 在线标的：在线标的件数、意向用户数、网拍围观人次
//
// 最长6年，年起始时间2017年
func TaobaoAuctionGovDataAnnuallyGet(ctx context.Context, clt *core.SDKClient, req *auction.TaobaoAuctionGovDataAnnuallyGetAPIRequest, resp *auction.TaobaoAuctionGovDataAnnuallyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
