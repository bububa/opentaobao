package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTraderateFeedsGetAPIRequest
查询子订单对应的评价、追评以及语义标签 API请求
tmall.traderate.feeds.get

通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签 */
type TmallTraderateFeedsGetAPIRequest struct {
	model.Params
	// 交易子订单ID
	_childTradeId int64
}

// New
