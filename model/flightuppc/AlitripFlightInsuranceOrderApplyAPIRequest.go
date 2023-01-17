package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightInsuranceOrderApplyAPIRequest 多险种批量投保 API请求
// alitrip.flight.insurance.order.apply
//
// 多险种批量投保
type AlitripFlightInsuranceOrderApplyAPIRequest struct {
	model.Params
	// 请求体
	_insApplyReq *InsApplyReq
}

// NewAlitripFlightInsuranceOrderApplyRequest 初始化AlitripFlightInsuranceOrderApplyAPIRequest对象
func NewAlitripFlightInsuranceOrderApplyRequest() *AlitripFlightInsuranceOrderApplyAPIRequest {
	return &AlitripFlightInsuranceOrderApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripFlightInsuranceOrderApplyAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.insurance.order.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripFlightInsuranceOrderApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripFlightInsuranceOrderApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInsApplyReq is InsApplyReq Setter
// 请求体
func (r *AlitripFlightInsuranceOrderApplyAPIRequest) SetInsApplyReq(_insApplyReq *InsApplyReq) error {
	r._insApplyReq = _insApplyReq
	r.Set("ins_apply_req", _insApplyReq)
	return nil
}

// GetInsApplyReq InsApplyReq Getter
func (r AlitripFlightInsuranceOrderApplyAPIRequest) GetInsApplyReq() *InsApplyReq {
	return r._insApplyReq
}
