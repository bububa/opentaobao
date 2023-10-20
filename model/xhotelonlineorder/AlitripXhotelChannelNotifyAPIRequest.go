package xhotelonlineorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripXhotelChannelNotifyAPIRequest 分销渠道各类通知接口 API请求
// alitrip.xhotel.channel.notify
//
// 分销渠道支付通知
type AlitripXhotelChannelNotifyAPIRequest struct {
	model.Params
	// 通知类型查询条件
	_orderNotifyQuery *OrderNotifyQuery
}

// NewAlitripXhotelChannelNotifyRequest 初始化AlitripXhotelChannelNotifyAPIRequest对象
func NewAlitripXhotelChannelNotifyRequest() *AlitripXhotelChannelNotifyAPIRequest {
	return &AlitripXhotelChannelNotifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripXhotelChannelNotifyAPIRequest) Reset() {
	r._orderNotifyQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripXhotelChannelNotifyAPIRequest) GetApiMethodName() string {
	return "alitrip.xhotel.channel.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripXhotelChannelNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripXhotelChannelNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNotifyQuery is OrderNotifyQuery Setter
// 通知类型查询条件
func (r *AlitripXhotelChannelNotifyAPIRequest) SetOrderNotifyQuery(_orderNotifyQuery *OrderNotifyQuery) error {
	r._orderNotifyQuery = _orderNotifyQuery
	r.Set("order_notify_query", _orderNotifyQuery)
	return nil
}

// GetOrderNotifyQuery OrderNotifyQuery Getter
func (r AlitripXhotelChannelNotifyAPIRequest) GetOrderNotifyQuery() *OrderNotifyQuery {
	return r._orderNotifyQuery
}

var poolAlitripXhotelChannelNotifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripXhotelChannelNotifyRequest()
	},
}

// GetAlitripXhotelChannelNotifyRequest 从 sync.Pool 获取 AlitripXhotelChannelNotifyAPIRequest
func GetAlitripXhotelChannelNotifyAPIRequest() *AlitripXhotelChannelNotifyAPIRequest {
	return poolAlitripXhotelChannelNotifyAPIRequest.Get().(*AlitripXhotelChannelNotifyAPIRequest)
}

// ReleaseAlitripXhotelChannelNotifyAPIRequest 将 AlitripXhotelChannelNotifyAPIRequest 放入 sync.Pool
func ReleaseAlitripXhotelChannelNotifyAPIRequest(v *AlitripXhotelChannelNotifyAPIRequest) {
	v.Reset()
	poolAlitripXhotelChannelNotifyAPIRequest.Put(v)
}
