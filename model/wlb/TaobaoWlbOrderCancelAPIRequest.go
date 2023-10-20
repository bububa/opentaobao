package wlb

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderCancelAPIRequest 取消物流宝订单 API请求
// taobao.wlb.order.cancel
//
// 取消物流宝订单
type TaobaoWlbOrderCancelAPIRequest struct {
	model.Params
	// 物流宝订单编号
	_wlbOrderCode string
}

// NewTaobaoWlbOrderCancelRequest 初始化TaobaoWlbOrderCancelAPIRequest对象
func NewTaobaoWlbOrderCancelRequest() *TaobaoWlbOrderCancelAPIRequest {
	return &TaobaoWlbOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbOrderCancelAPIRequest) Reset() {
	r._wlbOrderCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderCancelAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWlbOrderCode is WlbOrderCode Setter
// 物流宝订单编号
func (r *TaobaoWlbOrderCancelAPIRequest) SetWlbOrderCode(_wlbOrderCode string) error {
	r._wlbOrderCode = _wlbOrderCode
	r.Set("wlb_order_code", _wlbOrderCode)
	return nil
}

// GetWlbOrderCode WlbOrderCode Getter
func (r TaobaoWlbOrderCancelAPIRequest) GetWlbOrderCode() string {
	return r._wlbOrderCode
}

var poolTaobaoWlbOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbOrderCancelRequest()
	},
}

// GetTaobaoWlbOrderCancelRequest 从 sync.Pool 获取 TaobaoWlbOrderCancelAPIRequest
func GetTaobaoWlbOrderCancelAPIRequest() *TaobaoWlbOrderCancelAPIRequest {
	return poolTaobaoWlbOrderCancelAPIRequest.Get().(*TaobaoWlbOrderCancelAPIRequest)
}

// ReleaseTaobaoWlbOrderCancelAPIRequest 将 TaobaoWlbOrderCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbOrderCancelAPIRequest(v *TaobaoWlbOrderCancelAPIRequest) {
	v.Reset()
	poolTaobaoWlbOrderCancelAPIRequest.Put(v)
}
