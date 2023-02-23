package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeCancelAPIRequest 商旅机票改签取消 API请求
// alitrip.btrip.flight.distribution.change.cancel
//
// 商旅机票改签取消
type AlitripBtripFlightDistributionChangeCancelAPIRequest struct {
	model.Params
	// 改签取消输入参数
	_paramBtripFlightModifyCancelRq *BtripFlightModifyCancelRq
}

// NewAlitripBtripFlightDistributionChangeCancelRequest 初始化AlitripBtripFlightDistributionChangeCancelAPIRequest对象
func NewAlitripBtripFlightDistributionChangeCancelRequest() *AlitripBtripFlightDistributionChangeCancelAPIRequest {
	return &AlitripBtripFlightDistributionChangeCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangeCancelAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangeCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionChangeCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifyCancelRq is ParamBtripFlightModifyCancelRq Setter
// 改签取消输入参数
func (r *AlitripBtripFlightDistributionChangeCancelAPIRequest) SetParamBtripFlightModifyCancelRq(_paramBtripFlightModifyCancelRq *BtripFlightModifyCancelRq) error {
	r._paramBtripFlightModifyCancelRq = _paramBtripFlightModifyCancelRq
	r.Set("param_btrip_flight_modify_cancel_rq", _paramBtripFlightModifyCancelRq)
	return nil
}

// GetParamBtripFlightModifyCancelRq ParamBtripFlightModifyCancelRq Getter
func (r AlitripBtripFlightDistributionChangeCancelAPIRequest) GetParamBtripFlightModifyCancelRq() *BtripFlightModifyCancelRq {
	return r._paramBtripFlightModifyCancelRq
}
