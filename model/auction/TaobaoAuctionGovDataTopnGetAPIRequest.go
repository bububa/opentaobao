package auction

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAuctionGovDataTopnGetAPIRequest
根据不同维度，获取排行榜列表 API请求
taobao.auction.gov.data.topn.get

根据不同时间维度(周,月,年)，获取(成交额或发拍件数)排行榜列表 */
type TaobaoAuctionGovDataTopnGetAPIRequest struct {
	model.Params
	// 周期类型  （2：周，3：月，4：年）
	_circleType int64
	// 周期区间 周期（周填0、月份 yyyyMM、年份 yyyy)
	_circle int64
	// 业务类型 （1：成交额，2：发拍件数）
	_busiType int64
	// 区域类型（1：全国，2：全省）
	_zoneType int64
	// 法院名称
	_courtName string
}

// New
