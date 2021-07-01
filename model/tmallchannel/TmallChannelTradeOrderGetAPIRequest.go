package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeOrderGetAPIRequest
通过主采购单号查询采购单 API请求
tmall.channel.trade.order.get

通过主采购单号查询采购单 */
type TmallChannelTradeOrderGetAPIRequest struct {
	model.Params
	// 主采购单ID
	_mainPurchaseOrderNo int64
	// 是否包含子采购单
	_isIncludeSubOrder bool
	// 是否包含主采购单（针对特殊业务）
	_isIncludeMainOrder bool
	// 是否包含物流信息
	_isIncludeLogistics bool
}

// New
