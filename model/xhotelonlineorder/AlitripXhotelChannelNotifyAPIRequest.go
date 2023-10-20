package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripxhotelchannelnotifyAPIRequest 分销渠道各类通知接口 API请求
// alitrip.xhotel.channel.notify
//
// 分销渠道支付通知
type AlitripxhotelchannelnotifyAPIRequest struct {
	model.Params
	// 通知类型查询条件
	_orderNotifyQuery *OrderNotifyQuery
}

// NewAlitripxhotelchannelnotifyRequest 初始化AlitripxhotelchannelnotifyAPIRequest对象
func NewAlitripxhotelchannelnotifyRequest() *AlitripxhotelchannelnotifyAPIRequest {
	return &AlitripxhotelchannelnotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripxhotelchannelnotifyAPIRequest) GetApiMethodName() string {
	return "alitrip.xhotel.channel.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripxhotelchannelnotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripxhotelchannelnotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNotifyQuery is OrderNotifyQuery Setter
// 通知类型查询条件
func (r *AlitripxhotelchannelnotifyAPIRequest) SetOrderNotifyQuery(_orderNotifyQuery *OrderNotifyQuery) error {
	r._orderNotifyQuery = _orderNotifyQuery
	r.Set("order_notify_query", _orderNotifyQuery)
	return nil
}

// GetOrderNotifyQuery OrderNotifyQuery Getter
func (r AlitripxhotelchannelnotifyAPIRequest) GetOrderNotifyQuery() *OrderNotifyQuery {
	return r._orderNotifyQuery
}
