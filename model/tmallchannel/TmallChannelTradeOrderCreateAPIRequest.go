package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeOrderCreateAPIRequest
创建渠道分销单 API请求
tmall.channel.trade.order.create

创建渠道分销单 */
type TmallChannelTradeOrderCreateAPIRequest struct {
	model.Params
	// 入参
	_param0 *TopChannelPurchaseOrderCreateParam
}

// New
