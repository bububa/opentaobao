package auction

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAuctionGovDataMonthlyGetAPIRequest
按月统计法院拍卖数据 API请求
taobao.auction.gov.data.monthly.get

按月统计法院拍卖数据
包含：
标的件数统计：发布标的件数、结束标的件数、开拍标的件数
竞价实况：预计成交金额、出价次数、报名人数
在线标的：在线标的件数、意向用户数、网拍围观人次

最长12个月，月的起始时间不能早于2017年3月 */
type TaobaoAuctionGovDataMonthlyGetAPIRequest struct {
	model.Params
	// 法院名称
	_courtName string
	// 统计数据是够包含下属法院
	_isIncludeSub bool
	// 开始月份
	_startMonth string
	// 截止月份(统计数据包含这个月)
	_endMonth string
}

// New
