package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeApplyorderRefuseAPIRequest
供应商审核拒绝采购申请单 API请求
tmall.channel.trade.applyorder.refuse

供应商审核拒绝采购申请单 */
type TmallChannelTradeApplyorderRefuseAPIRequest struct {
	model.Params
	// 操作描述
	_operateDesc string
	// 采购申请单号
	_channelPurchaseApplyOrderNo string
}

// New
