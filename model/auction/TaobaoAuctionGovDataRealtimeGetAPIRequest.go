package auction

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAuctionGovDataRealtimeGetAPIRequest
获取实时(今日)统计数据 API请求
taobao.auction.gov.data.realtime.get

提供查询当日法院及下属法院的拍卖统计数据 */
type TaobaoAuctionGovDataRealtimeGetAPIRequest struct {
	model.Params
	// 法院名称
	_courtName string
	// 统计数据是否包含下级法院
	_isIncludeSub bool
}

// New
