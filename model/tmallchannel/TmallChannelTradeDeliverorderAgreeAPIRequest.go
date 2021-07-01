package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeDeliverorderAgreeAPIRequest
供应商审核通过发货确认 API请求
tmall.channel.trade.deliverorder.agree

供应商通过收货确认单 */
type TmallChannelTradeDeliverorderAgreeAPIRequest struct {
	model.Params
	// 发货单号
	_mainDeliverOrderNo int64
	// 同意理由
	_operateDesc string
}

// New
