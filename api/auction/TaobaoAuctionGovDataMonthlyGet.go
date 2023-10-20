package auction

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/auction"
)

// Taobaoauctiongovdatamonthlyget 按月统计法院拍卖数据
// taobao.auction.gov.data.monthly.get
//
// 按月统计法院拍卖数据
// 包含：
// 标的件数统计：发布标的件数、结束标的件数、开拍标的件数
// 竞价实况：预计成交金额、出价次数、报名人数
// 在线标的：在线标的件数、意向用户数、网拍围观人次
//
// 最长12个月，月的起始时间不能早于2017年3月
func Taobaoauctiongovdatamonthlyget(clt *core.SDKClient, req *auction.TaobaoauctiongovdatamonthlygetAPIRequest, session string) (*auction.TaobaoauctiongovdatamonthlygetAPIResponse, error) {
	var resp auction.TaobaoauctiongovdatamonthlygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
