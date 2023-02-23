package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeQueryAPIRequest 改签航班询价 API请求
// alitrip.btrip.flight.distribution.change.query
//
// 商旅机票改签航班询价
type AlitripBtripFlightDistributionChangeQueryAPIRequest struct {
	model.Params
	// 改签询价入参
	_paramBtripFlightModifyFlightInfoRq *BtripFlightModifyFlightInfoRq
}

// NewAlitripBtripFlightDistributionChangeQueryRequest 初始化AlitripBtripFlightDistributionChangeQueryAPIRequest对象
func NewAlitripBtripFlightDistributionChangeQueryRequest() *AlitripBtripFlightDistributionChangeQueryAPIRequest {
	return &AlitripBtripFlightDistributionChangeQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangeQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionChangeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifyFlightInfoRq is ParamBtripFlightModifyFlightInfoRq Setter
// 改签询价入参
func (r *AlitripBtripFlightDistributionChangeQueryAPIRequest) SetParamBtripFlightModifyFlightInfoRq(_paramBtripFlightModifyFlightInfoRq *BtripFlightModifyFlightInfoRq) error {
	r._paramBtripFlightModifyFlightInfoRq = _paramBtripFlightModifyFlightInfoRq
	r.Set("param_btrip_flight_modify_flight_info_rq", _paramBtripFlightModifyFlightInfoRq)
	return nil
}

// GetParamBtripFlightModifyFlightInfoRq ParamBtripFlightModifyFlightInfoRq Getter
func (r AlitripBtripFlightDistributionChangeQueryAPIRequest) GetParamBtripFlightModifyFlightInfoRq() *BtripFlightModifyFlightInfoRq {
	return r._paramBtripFlightModifyFlightInfoRq
}
