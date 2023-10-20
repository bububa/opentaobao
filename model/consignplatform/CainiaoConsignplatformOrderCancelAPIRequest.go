package consignplatform

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoConsignplatformOrderCancelAPIRequest 菜鸟发货工作台取消包裹以及订单 API请求
// cainiao.consignplatform.order.cancel
//
// 菜鸟发货工作台，商家或者isv通过api取消包裹、回收单号，如果是裹裹运力会取消小件员上门。最后删除订单信息。
type CainiaoConsignplatformOrderCancelAPIRequest struct {
	model.Params
	// 取消参数
	_cancelRequest *OrderCancelRequest
}

// NewCainiaoConsignplatformOrderCancelRequest 初始化CainiaoConsignplatformOrderCancelAPIRequest对象
func NewCainiaoConsignplatformOrderCancelRequest() *CainiaoConsignplatformOrderCancelAPIRequest {
	return &CainiaoConsignplatformOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoConsignplatformOrderCancelAPIRequest) Reset() {
	r._cancelRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoConsignplatformOrderCancelAPIRequest) GetApiMethodName() string {
	return "cainiao.consignplatform.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoConsignplatformOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoConsignplatformOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelRequest is CancelRequest Setter
// 取消参数
func (r *CainiaoConsignplatformOrderCancelAPIRequest) SetCancelRequest(_cancelRequest *OrderCancelRequest) error {
	r._cancelRequest = _cancelRequest
	r.Set("cancel_request", _cancelRequest)
	return nil
}

// GetCancelRequest CancelRequest Getter
func (r CainiaoConsignplatformOrderCancelAPIRequest) GetCancelRequest() *OrderCancelRequest {
	return r._cancelRequest
}

var poolCainiaoConsignplatformOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoConsignplatformOrderCancelRequest()
	},
}

// GetCainiaoConsignplatformOrderCancelRequest 从 sync.Pool 获取 CainiaoConsignplatformOrderCancelAPIRequest
func GetCainiaoConsignplatformOrderCancelAPIRequest() *CainiaoConsignplatformOrderCancelAPIRequest {
	return poolCainiaoConsignplatformOrderCancelAPIRequest.Get().(*CainiaoConsignplatformOrderCancelAPIRequest)
}

// ReleaseCainiaoConsignplatformOrderCancelAPIRequest 将 CainiaoConsignplatformOrderCancelAPIRequest 放入 sync.Pool
func ReleaseCainiaoConsignplatformOrderCancelAPIRequest(v *CainiaoConsignplatformOrderCancelAPIRequest) {
	v.Reset()
	poolCainiaoConsignplatformOrderCancelAPIRequest.Put(v)
}
