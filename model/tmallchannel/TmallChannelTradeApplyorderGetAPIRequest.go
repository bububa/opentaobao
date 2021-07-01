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

// NewTmallChannelTradeApplyorderGetRequest 初始化TmallChannelTradeApplyorderGetAPIRequest对象
func NewTmallChannelTradeApplyorderGetRequest() *TmallChannelTradeApplyorderGetAPIRequest {
	return &TmallChannelTradeApplyorderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallChannelTradeApplyorderGetAPIRequest) GetApiMethodName() string {
	return "tmall.channel.trade.applyorder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallChannelTradeApplyorderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ChannelPurchaseApplyOrderNo Setter
// 采购申请单单号
func (r *TmallChannelTradeApplyorderGetAPIRequest) SetChannelPurchaseApplyOrderNo(_channelPurchaseApplyOrderNo string) error {
	r._channelPurchaseApplyOrderNo = _channelPurchaseApplyOrderNo
	r.Set("channel_purchase_apply_order_no", _channelPurchaseApplyOrderNo)
	return nil
}

// Get ChannelPurchaseApplyOrderNo Getter
func (r TmallChannelTradeApplyorderGetAPIRequest) GetChannelPurchaseApplyOrderNo() string {
	return r._channelPurchaseApplyOrderNo
}
