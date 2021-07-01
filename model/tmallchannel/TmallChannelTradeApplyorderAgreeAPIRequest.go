package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeApplyorderAgreeAPIRequest
供应商审核同意采购申请单 API请求
tmall.channel.trade.applyorder.agree

供应商审核同意采购申请单 */
type TmallChannelTradeApplyorderAgreeAPIRequest struct {
	model.Params
	// 操作描述
	_operateDesc string
	// 采购申请单号
	_channelPurchaseApplyOrderNo string
}

// New
