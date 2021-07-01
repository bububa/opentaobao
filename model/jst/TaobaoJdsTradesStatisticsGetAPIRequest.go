package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJdsTradesStatisticsGetAPIRequest
获取订单数量统计结果 API请求
taobao.jds.trades.statistics.get

获取订单数量统计结果 */
type TaobaoJdsTradesStatisticsGetAPIRequest struct {
	model.Params
	// 查询的日期，格式如YYYYMMDD的日期对应的数字
	_date int64
}

// New
