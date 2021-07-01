package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallExchangeReturngoodsAgreeAPIRequest
卖家确认收货 API请求
tmall.exchange.returngoods.agree

卖家确认收货 */
type TmallExchangeReturngoodsAgreeAPIRequest struct {
	model.Params
	// 换货单号ID
	_disputeId int64
	// 返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态）
	_fields []string
}

// New
