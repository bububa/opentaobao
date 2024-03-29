package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPhoneOrderExternalStatusAPIRequest 话费外放订单状态接口 API请求
// taobao.phone.order.external.status
//
// 话费外放订单状态接口
type TaobaoPhoneOrderExternalStatusAPIRequest struct {
	model.Params
	// 状态查询请求
	_queryOrderReq *QueryOrderReq
}

// NewTaobaoPhoneOrderExternalStatusRequest 初始化TaobaoPhoneOrderExternalStatusAPIRequest对象
func NewTaobaoPhoneOrderExternalStatusRequest() *TaobaoPhoneOrderExternalStatusAPIRequest {
	return &TaobaoPhoneOrderExternalStatusAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPhoneOrderExternalStatusAPIRequest) Reset() {
	r._queryOrderReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPhoneOrderExternalStatusAPIRequest) GetApiMethodName() string {
	return "taobao.phone.order.external.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPhoneOrderExternalStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPhoneOrderExternalStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryOrderReq is QueryOrderReq Setter
// 状态查询请求
func (r *TaobaoPhoneOrderExternalStatusAPIRequest) SetQueryOrderReq(_queryOrderReq *QueryOrderReq) error {
	r._queryOrderReq = _queryOrderReq
	r.Set("query_order_req", _queryOrderReq)
	return nil
}

// GetQueryOrderReq QueryOrderReq Getter
func (r TaobaoPhoneOrderExternalStatusAPIRequest) GetQueryOrderReq() *QueryOrderReq {
	return r._queryOrderReq
}

var poolTaobaoPhoneOrderExternalStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPhoneOrderExternalStatusRequest()
	},
}

// GetTaobaoPhoneOrderExternalStatusRequest 从 sync.Pool 获取 TaobaoPhoneOrderExternalStatusAPIRequest
func GetTaobaoPhoneOrderExternalStatusAPIRequest() *TaobaoPhoneOrderExternalStatusAPIRequest {
	return poolTaobaoPhoneOrderExternalStatusAPIRequest.Get().(*TaobaoPhoneOrderExternalStatusAPIRequest)
}

// ReleaseTaobaoPhoneOrderExternalStatusAPIRequest 将 TaobaoPhoneOrderExternalStatusAPIRequest 放入 sync.Pool
func ReleaseTaobaoPhoneOrderExternalStatusAPIRequest(v *TaobaoPhoneOrderExternalStatusAPIRequest) {
	v.Reset()
	poolTaobaoPhoneOrderExternalStatusAPIRequest.Put(v)
}
