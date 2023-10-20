package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliiconfirmAPIRequest 物流订单确认接口 API请求
// cainiao.waybill.ii.confirm
//
// 物流订单确认
type CainiaowaybilliiconfirmAPIRequest struct {
	model.Params
	// 订单确认信息
	_paramWaybillOrderConfirmRequest *WaybillOrderConfirmRequest
}

// NewCainiaowaybilliiconfirmRequest 初始化CainiaowaybilliiconfirmAPIRequest对象
func NewCainiaowaybilliiconfirmRequest() *CainiaowaybilliiconfirmAPIRequest {
	return &CainiaowaybilliiconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybilliiconfirmAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybilliiconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybilliiconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWaybillOrderConfirmRequest is ParamWaybillOrderConfirmRequest Setter
// 订单确认信息
func (r *CainiaowaybilliiconfirmAPIRequest) SetParamWaybillOrderConfirmRequest(_paramWaybillOrderConfirmRequest *WaybillOrderConfirmRequest) error {
	r._paramWaybillOrderConfirmRequest = _paramWaybillOrderConfirmRequest
	r.Set("param_waybill_order_confirm_request", _paramWaybillOrderConfirmRequest)
	return nil
}

// GetParamWaybillOrderConfirmRequest ParamWaybillOrderConfirmRequest Getter
func (r CainiaowaybilliiconfirmAPIRequest) GetParamWaybillOrderConfirmRequest() *WaybillOrderConfirmRequest {
	return r._paramWaybillOrderConfirmRequest
}
