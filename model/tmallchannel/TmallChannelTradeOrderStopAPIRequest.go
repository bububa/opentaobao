package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeOrderStopAPIRequest
供应商停止发货 API请求
tmall.channel.trade.order.stop

供应商停止发货 */
type TmallChannelTradeOrderStopAPIRequest struct {
	model.Params
	// 主采购单号
	_mainPurchaseOrderNo int64
	// 幂等单号
	_requestNo string
}

// New
