package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeApplyAPIRequest 商旅机票改签申请 API请求
// alitrip.btrip.flight.distribution.change.apply
//
// 商旅机票改签申请
type AlitripBtripFlightDistributionChangeApplyAPIRequest struct {
	model.Params
	// 改签申请入参
	_paramBtripFlightModifyApplyRq *BtripFlightModifyApplyRq
}

// NewAlitripBtripFlightDistributionChangeApplyRequest 初始化AlitripBtripFlightDistributionChangeApplyAPIRequest对象
func NewAlitripBtripFlightDistributionChangeApplyRequest() *AlitripBtripFlightDistributionChangeApplyAPIRequest {
	return &AlitripBtripFlightDistributionChangeApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangeApplyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangeApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripFlightModifyApplyRq is ParamBtripFlightModifyApplyRq Setter
// 改签申请入参
func (r *AlitripBtripFlightDistributionChangeApplyAPIRequest) SetParamBtripFlightModifyApplyRq(_paramBtripFlightModifyApplyRq *BtripFlightModifyApplyRq) error {
	r._paramBtripFlightModifyApplyRq = _paramBtripFlightModifyApplyRq
	r.Set("param_btrip_flight_modify_apply_rq", _paramBtripFlightModifyApplyRq)
	return nil
}

// GetParamBtripFlightModifyApplyRq ParamBtripFlightModifyApplyRq Getter
func (r AlitripBtripFlightDistributionChangeApplyAPIRequest) GetParamBtripFlightModifyApplyRq() *BtripFlightModifyApplyRq {
	return r._paramBtripFlightModifyApplyRq
}
