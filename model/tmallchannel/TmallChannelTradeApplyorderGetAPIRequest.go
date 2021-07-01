package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeApplyorderGetAPIRequest
查询采购申请单详情 API请求
tmall.channel.trade.applyorder.get

通过采购申请单ID获取单据详情 */
type TmallChannelTradeApplyorderGetAPIRequest struct {
	model.Params
	// 采购申请单单号
	_channelPurchaseApplyOrderNo string
}

// New
