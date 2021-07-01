package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeDeliverorderRejectAPIRequest
供应商拒绝收货确认单 API请求
tmall.channel.trade.deliverorder.reject

供应商拒绝收货确认单 */
type TmallChannelTradeDeliverorderRejectAPIRequest struct {
	model.Params
	// 发货单号
	_mainDeliverOrderNo int64
	// 拒绝原因
	_operateDesc string
}

// New
