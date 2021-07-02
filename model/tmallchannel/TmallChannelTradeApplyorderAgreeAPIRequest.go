package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelTradeApplyorderAgreeAPIRequest 供应商审核同意采购申请单 API请求
// tmall.channel.trade.applyorder.agree
//
// 供应商审核同意采购申请单
type TmallChannelTradeApplyorderAgreeAPIRequest struct {
	model.Params
	// 操作描述
	_operateDesc string
	// 采购申请单号
	_channelPurchaseApplyOrderNo string
}

// NewTmallChannelTradeApplyorderAgreeRequest 初始化TmallChannelTradeApplyorderAgreeAPIRequest对象
func NewTmallChannelTradeApplyorderAgreeRequest() *TmallChannelTradeApplyorderAgreeAPIRequest {
	return &TmallChannelTradeApplyorderAgreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallChannelTradeApplyorderAgreeAPIRequest) GetApiMethodName() string {
	return "tmall.channel.trade.applyorder.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallChannelTradeApplyorderAgreeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOperateDesc is OperateDesc Setter
// 操作描述
func (r *TmallChannelTradeApplyorderAgreeAPIRequest) SetOperateDesc(_operateDesc string) error {
	r._operateDesc = _operateDesc
	r.Set("operate_desc", _operateDesc)
	return nil
}

// GetOperateDesc OperateDesc Getter
func (r TmallChannelTradeApplyorderAgreeAPIRequest) GetOperateDesc() string {
	return r._operateDesc
}

// SetChannelPurchaseApplyOrderNo is ChannelPurchaseApplyOrderNo Setter
// 采购申请单号
func (r *TmallChannelTradeApplyorderAgreeAPIRequest) SetChannelPurchaseApplyOrderNo(_channelPurchaseApplyOrderNo string) error {
	r._channelPurchaseApplyOrderNo = _channelPurchaseApplyOrderNo
	r.Set("channel_purchase_apply_order_no", _channelPurchaseApplyOrderNo)
	return nil
}

// GetChannelPurchaseApplyOrderNo ChannelPurchaseApplyOrderNo Getter
func (r TmallChannelTradeApplyorderAgreeAPIRequest) GetChannelPurchaseApplyOrderNo() string {
	return r._channelPurchaseApplyOrderNo
}
