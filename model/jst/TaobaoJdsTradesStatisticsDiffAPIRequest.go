package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJdsTradesStatisticsDiffAPIRequest
订单全链路状态统计差异比较 API请求
taobao.jds.trades.statistics.diff

订单全链路状态统计差异比较 */
type TaobaoJdsTradesStatisticsDiffAPIRequest struct {
	model.Params
	// 查询的日期，格式如YYYYMMDD的日期对应的数字
	_date int64
	// 需要比较的状态，将会和post_status做比较
	_preStatus string
	// 需要比较的状态
	_postStatus string
	// 页数
	_pageNo int64
}

// New
