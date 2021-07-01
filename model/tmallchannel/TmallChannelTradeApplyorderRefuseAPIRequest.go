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

// NewTmallChannelTradeApplyorderRefuseRequest 初始化TmallChannelTradeApplyorderRefuseAPIRequest对象
func NewTmallChannelTradeApplyorderRefuseRequest() *TmallChannelTradeApplyorderRefuseAPIRequest {
	return &TmallChannelTradeApplyorderRefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallChannelTradeApplyorderRefuseAPIRequest) GetApiMethodName() string {
	return "tmall.channel.trade.applyorder.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallChannelTradeApplyorderRefuseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OperateDesc Setter
// 操作描述
func (r *TmallChannelTradeApplyorderRefuseAPIRequest) SetOperateDesc(_operateDesc string) error {
	r._operateDesc = _operateDesc
	r.Set("operate_desc", _operateDesc)
	return nil
}

// Get OperateDesc Getter
func (r TmallChannelTradeApplyorderRefuseAPIRequest) GetOperateDesc() string {
	return r._operateDesc
}

// Set is ChannelPurchaseApplyOrderNo Setter
// 采购申请单号
func (r *TmallChannelTradeApplyorderRefuseAPIRequest) SetChannelPurchaseApplyOrderNo(_channelPurchaseApplyOrderNo string) error {
	r._channelPurchaseApplyOrderNo = _channelPurchaseApplyOrderNo
	r.Set("channel_purchase_apply_order_no", _channelPurchaseApplyOrderNo)
	return nil
}

// Get ChannelPurchaseApplyOrderNo Getter
func (r TmallChannelTradeApplyorderRefuseAPIRequest) GetChannelPurchaseApplyOrderNo() string {
	return r._channelPurchaseApplyOrderNo
}
