package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeDeliverorderAgreeAPIRequest
供应商审核通过发货确认 API请求
tmall.channel.trade.deliverorder.agree

供应商通过收货确认单 */
type TmallChannelTradeDeliverorderAgreeAPIRequest struct {
	model.Params
	// 发货单号
	_mainDeliverOrderNo int64
	// 同意理由
	_operateDesc string
}

// NewTmallChannelTradeDeliverorderAgreeRequest 初始化TmallChannelTradeDeliverorderAgreeAPIRequest对象
func NewTmallChannelTradeDeliverorderAgreeRequest() *TmallChannelTradeDeliverorderAgreeAPIRequest {
	return &TmallChannelTradeDeliverorderAgreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallChannelTradeDeliverorderAgreeAPIRequest) GetApiMethodName() string {
	return "tmall.channel.trade.deliverorder.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallChannelTradeDeliverorderAgreeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MainDeliverOrderNo Setter
// 发货单号
func (r *TmallChannelTradeDeliverorderAgreeAPIRequest) SetMainDeliverOrderNo(_mainDeliverOrderNo int64) error {
	r._mainDeliverOrderNo = _mainDeliverOrderNo
	r.Set("main_deliver_order_no", _mainDeliverOrderNo)
	return nil
}

// Get MainDeliverOrderNo Getter
func (r TmallChannelTradeDeliverorderAgreeAPIRequest) GetMainDeliverOrderNo() int64 {
	return r._mainDeliverOrderNo
}

// Set is OperateDesc Setter
// 同意理由
func (r *TmallChannelTradeDeliverorderAgreeAPIRequest) SetOperateDesc(_operateDesc string) error {
	r._operateDesc = _operateDesc
	r.Set("operate_desc", _operateDesc)
	return nil
}

// Get OperateDesc Getter
func (r TmallChannelTradeDeliverorderAgreeAPIRequest) GetOperateDesc() string {
	return r._operateDesc
}
