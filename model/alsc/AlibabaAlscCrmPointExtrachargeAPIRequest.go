package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmpointextrachargeAPIRequest 积分补录 API请求
// alibaba.alsc.crm.point.extracharge
//
// 积分补录
type AlibabaalsccrmpointextrachargeAPIRequest struct {
	model.Params
	// 入参
	_paramExtraChargePointOpenReq *ExtraChargePointOpenReq
}

// NewAlibabaalsccrmpointextrachargeRequest 初始化AlibabaalsccrmpointextrachargeAPIRequest对象
func NewAlibabaalsccrmpointextrachargeRequest() *AlibabaalsccrmpointextrachargeAPIRequest {
	return &AlibabaalsccrmpointextrachargeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmpointextrachargeAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.extracharge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmpointextrachargeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmpointextrachargeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamExtraChargePointOpenReq is ParamExtraChargePointOpenReq Setter
// 入参
func (r *AlibabaalsccrmpointextrachargeAPIRequest) SetParamExtraChargePointOpenReq(_paramExtraChargePointOpenReq *ExtraChargePointOpenReq) error {
	r._paramExtraChargePointOpenReq = _paramExtraChargePointOpenReq
	r.Set("param_extra_charge_point_open_req", _paramExtraChargePointOpenReq)
	return nil
}

// GetParamExtraChargePointOpenReq ParamExtraChargePointOpenReq Getter
func (r AlibabaalsccrmpointextrachargeAPIRequest) GetParamExtraChargePointOpenReq() *ExtraChargePointOpenReq {
	return r._paramExtraChargePointOpenReq
}
