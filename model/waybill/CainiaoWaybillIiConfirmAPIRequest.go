package waybill

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiConfirmAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
