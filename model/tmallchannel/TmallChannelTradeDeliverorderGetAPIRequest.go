package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeDeliverorderGetAPIRequest
通过发货单单号获取发货单的详情 API请求
tmall.channel.trade.deliverorder.get

通过发货单单号获取发货单的详情 */
type TmallChannelTradeDeliverorderGetAPIRequest struct {
	model.Params
	// 发货单号
	_mainDeliverOrderNo int64
	// 是否包含子发货单
	_isIncludeSubOrder bool
}

// New
