package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelTradeDeliverorderRejectAPIRequest 供应商拒绝收货确认单 API请求
// tmall.channel.trade.deliverorder.reject
//
// 供应商拒绝收货确认单
type TmallChannelTradeDeliverorderRejectAPIRequest struct {
	model.Params
	// 发货单号
	_mainDeliverOrderNo int64
	// 拒绝原因
	_operateDesc string
}

// NewTmallChannelTradeDeliverorderRejectRequest 初始化TmallChannelTradeDeliverorderRejectAPIRequest对象
func NewTmallChannelTradeDeliverorderRejectRequest() *TmallChannelTradeDeliverorderRejectAPIRequest {
	return &TmallChannelTradeDeliverorderRejectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallChannelTradeDeliverorderRejectAPIRequest) GetApiMethodName() string {
	return "tmall.channel.trade.deliverorder.reject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallChannelTradeDeliverorderRejectAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMainDeliverOrderNo is MainDeliverOrderNo Setter
// 发货单号
func (r *TmallChannelTradeDeliverorderRejectAPIRequest) SetMainDeliverOrderNo(_mainDeliverOrderNo int64) error {
	r._mainDeliverOrderNo = _mainDeliverOrderNo
	r.Set("main_deliver_order_no", _mainDeliverOrderNo)
	return nil
}

// GetMainDeliverOrderNo MainDeliverOrderNo Getter
func (r TmallChannelTradeDeliverorderRejectAPIRequest) GetMainDeliverOrderNo() int64 {
	return r._mainDeliverOrderNo
}

// SetOperateDesc is OperateDesc Setter
// 拒绝原因
func (r *TmallChannelTradeDeliverorderRejectAPIRequest) SetOperateDesc(_operateDesc string) error {
	r._operateDesc = _operateDesc
	r.Set("operate_desc", _operateDesc)
	return nil
}

// GetOperateDesc OperateDesc Getter
func (r TmallChannelTradeDeliverorderRejectAPIRequest) GetOperateDesc() string {
	return r._operateDesc
}
