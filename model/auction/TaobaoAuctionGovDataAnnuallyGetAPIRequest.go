package auction

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAuctionGovDataAnnuallyGetAPIRequest
按年统计法院拍卖数据 API请求
taobao.auction.gov.data.annually.get

按月统计法院拍卖数据 包含：
标的件数统计：发布标的件数、结束标的件数、开拍标的件数
竞价实况：预计成交金额、出价次数、报名人数
在线标的：在线标的件数、意向用户数、网拍围观人次

最长6年，年起始时间2017年 */
type TaobaoAuctionGovDataAnnuallyGetAPIRequest struct {
	model.Params
	// 法院名称
	_courtName string
	// 统计数据是够包含下属法院
	_isIncludeSub bool
	// 开始年份
	_startYear string
	// 结束年份
	_endYear string
}

// New
