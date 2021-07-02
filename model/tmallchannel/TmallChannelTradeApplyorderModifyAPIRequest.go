package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelTradeApplyorderModifyAPIRequest 供应商修改申请单 API请求
// tmall.channel.trade.applyorder.modify
//
// 上游供应商修改申请单, 目前只允许修改价格+件数且sku数量必须完全一致
type TmallChannelTradeApplyorderModifyAPIRequest struct {
	model.Params
	// 采购申请单号
	_channelPurchaseApplyOrderNo string
	// 修改关联的的宝贝信息
	_applyOrderRelateItemModifyParamList []TopChannelApplyOrderRelateItemModifyParam
}

// NewTmallChannelTradeApplyorderModifyRequest 初始化TmallChannelTradeApplyorderModifyAPIRequest对象
func NewTmallChannelTradeApplyorderModifyRequest() *TmallChannelTradeApplyorderModifyAPIRequest {
	return &TmallChannelTradeApplyorderModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallChannelTradeApplyorderModifyAPIRequest) GetApiMethodName() string {
	return "tmall.channel.trade.applyorder.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallChannelTradeApplyorderModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetChannelPurchaseApplyOrderNo is ChannelPurchaseApplyOrderNo Setter
// 采购申请单号
func (r *TmallChannelTradeApplyorderModifyAPIRequest) SetChannelPurchaseApplyOrderNo(_channelPurchaseApplyOrderNo string) error {
	r._channelPurchaseApplyOrderNo = _channelPurchaseApplyOrderNo
	r.Set("channel_purchase_apply_order_no", _channelPurchaseApplyOrderNo)
	return nil
}

// GetChannelPurchaseApplyOrderNo ChannelPurchaseApplyOrderNo Getter
func (r TmallChannelTradeApplyorderModifyAPIRequest) GetChannelPurchaseApplyOrderNo() string {
	return r._channelPurchaseApplyOrderNo
}

// SetApplyOrderRelateItemModifyParamList is ApplyOrderRelateItemModifyParamList Setter
// 修改关联的的宝贝信息
func (r *TmallChannelTradeApplyorderModifyAPIRequest) SetApplyOrderRelateItemModifyParamList(_applyOrderRelateItemModifyParamList []TopChannelApplyOrderRelateItemModifyParam) error {
	r._applyOrderRelateItemModifyParamList = _applyOrderRelateItemModifyParamList
	r.Set("apply_order_relate_item_modify_param_list", _applyOrderRelateItemModifyParamList)
	return nil
}

// GetApplyOrderRelateItemModifyParamList ApplyOrderRelateItemModifyParamList Getter
func (r TmallChannelTradeApplyorderModifyAPIRequest) GetApplyOrderRelateItemModifyParamList() []TopChannelApplyOrderRelateItemModifyParam {
	return r._applyOrderRelateItemModifyParamList
}
