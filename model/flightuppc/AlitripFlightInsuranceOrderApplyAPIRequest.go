package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightinsuranceorderapplyAPIRequest 多险种批量投保 API请求
// alitrip.flight.insurance.order.apply
//
// 多险种批量投保
type AlitripflightinsuranceorderapplyAPIRequest struct {
	model.Params
	// 请求体
	_insApplyReq *InsApplyReq
}

// NewAlitripflightinsuranceorderapplyRequest 初始化AlitripflightinsuranceorderapplyAPIRequest对象
func NewAlitripflightinsuranceorderapplyRequest() *AlitripflightinsuranceorderapplyAPIRequest {
	return &AlitripflightinsuranceorderapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripflightinsuranceorderapplyAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.insurance.order.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripflightinsuranceorderapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripflightinsuranceorderapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInsApplyReq is InsApplyReq Setter
// 请求体
func (r *AlitripflightinsuranceorderapplyAPIRequest) SetInsApplyReq(_insApplyReq *InsApplyReq) error {
	r._insApplyReq = _insApplyReq
	r.Set("ins_apply_req", _insApplyReq)
	return nil
}

// GetInsApplyReq InsApplyReq Getter
func (r AlitripflightinsuranceorderapplyAPIRequest) GetInsApplyReq() *InsApplyReq {
	return r._insApplyReq
}
