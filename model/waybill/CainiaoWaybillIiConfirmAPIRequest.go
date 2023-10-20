package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiConfirmAPIRequest 物流订单确认接口 API请求
// cainiao.waybill.ii.confirm
//
// 物流订单确认
type CainiaoWaybillIiConfirmAPIRequest struct {
	model.Params
	// 订单确认信息
	_paramWaybillOrderConfirmRequest *WaybillOrderConfirmRequest
}

// NewCainiaoWaybillIiConfirmRequest 初始化CainiaoWaybillIiConfirmAPIRequest对象
func NewCainiaoWaybillIiConfirmRequest() *CainiaoWaybillIiConfirmAPIRequest {
	return &CainiaoWaybillIiConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoWaybillIiConfirmAPIRequest) Reset() {
	r._paramWaybillOrderConfirmRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiConfirmAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoWaybillIiConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWaybillOrderConfirmRequest is ParamWaybillOrderConfirmRequest Setter
// 订单确认信息
func (r *CainiaoWaybillIiConfirmAPIRequest) SetParamWaybillOrderConfirmRequest(_paramWaybillOrderConfirmRequest *WaybillOrderConfirmRequest) error {
	r._paramWaybillOrderConfirmRequest = _paramWaybillOrderConfirmRequest
	r.Set("param_waybill_order_confirm_request", _paramWaybillOrderConfirmRequest)
	return nil
}

// GetParamWaybillOrderConfirmRequest ParamWaybillOrderConfirmRequest Getter
func (r CainiaoWaybillIiConfirmAPIRequest) GetParamWaybillOrderConfirmRequest() *WaybillOrderConfirmRequest {
	return r._paramWaybillOrderConfirmRequest
}

var poolCainiaoWaybillIiConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoWaybillIiConfirmRequest()
	},
}

// GetCainiaoWaybillIiConfirmRequest 从 sync.Pool 获取 CainiaoWaybillIiConfirmAPIRequest
func GetCainiaoWaybillIiConfirmAPIRequest() *CainiaoWaybillIiConfirmAPIRequest {
	return poolCainiaoWaybillIiConfirmAPIRequest.Get().(*CainiaoWaybillIiConfirmAPIRequest)
}

// ReleaseCainiaoWaybillIiConfirmAPIRequest 将 CainiaoWaybillIiConfirmAPIRequest 放入 sync.Pool
func ReleaseCainiaoWaybillIiConfirmAPIRequest(v *CainiaoWaybillIiConfirmAPIRequest) {
	v.Reset()
	poolCainiaoWaybillIiConfirmAPIRequest.Put(v)
}
