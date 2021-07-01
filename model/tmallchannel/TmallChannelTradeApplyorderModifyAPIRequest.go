package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeApplyorderModifyAPIRequest
供应商修改申请单 API请求
tmall.channel.trade.applyorder.modify

上游供应商修改申请单, 目前只允许修改价格+件数且sku数量必须完全一致 */
type TmallChannelTradeApplyorderModifyAPIRequest struct {
	model.Params
	// 采购申请单号
	_channelPurchaseApplyOrderNo string
	// 修改关联的的宝贝信息
	_applyOrderRelateItemModifyParamList []TopChannelApplyOrderRelateItemModifyParam
}

// New
