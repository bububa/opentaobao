package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeNewcancelAPIRequest 商旅机票改签取消 API请求
// alitrip.btrip.flight.distribution.change.newcancel
//
// 商旅机票改签取消
type AlitripBtripFlightDistributionChangeNewcancelAPIRequest struct {
	model.Params
	// 改签取消输入参数
	_paramBtripFlightModifyCancelRq *BtripFlightModifyCancelRq
}

// NewAlitripBtripFlightDistributionChangeNewcancelRequest 初始化AlitripBtripFlightDistributionChangeNewcancelAPIRequest对象
func NewAlitripBtripFlightDistributionChangeNewcancelRequest() *AlitripBtripFlightDistributionChangeNewcancelAPIRequest {
	return &AlitripBtripFlightDistributionChangeNewcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangeNewcancelAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.newcancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangeNewcancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripFlightModifyCancelRq is ParamBtripFlightModifyCancelRq Setter
// 改签取消输入参数
func (r *AlitripBtripFlightDistributionChangeNewcancelAPIRequest) SetParamBtripFlightModifyCancelRq(_paramBtripFlightModifyCancelRq *BtripFlightModifyCancelRq) error {
	r._paramBtripFlightModifyCancelRq = _paramBtripFlightModifyCancelRq
	r.Set("param_btrip_flight_modify_cancel_rq", _paramBtripFlightModifyCancelRq)
	return nil
}

// GetParamBtripFlightModifyCancelRq ParamBtripFlightModifyCancelRq Getter
func (r AlitripBtripFlightDistributionChangeNewcancelAPIRequest) GetParamBtripFlightModifyCancelRq() *BtripFlightModifyCancelRq {
	return r._paramBtripFlightModifyCancelRq
}
