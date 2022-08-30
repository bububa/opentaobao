package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeNewqueryAPIRequest 改签航班询价V2 API请求
// alitrip.btrip.flight.distribution.change.newquery
//
// 商旅机票改签航班询价
type AlitripBtripFlightDistributionChangeNewqueryAPIRequest struct {
	model.Params
	// 改签询价入参
	_paramBtripFlightModifyFlightInfoRq *BtripFlightModifyFlightInfoRq
}

// NewAlitripBtripFlightDistributionChangeNewqueryRequest 初始化AlitripBtripFlightDistributionChangeNewqueryAPIRequest对象
func NewAlitripBtripFlightDistributionChangeNewqueryRequest() *AlitripBtripFlightDistributionChangeNewqueryAPIRequest {
	return &AlitripBtripFlightDistributionChangeNewqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangeNewqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.newquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangeNewqueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripFlightModifyFlightInfoRq is ParamBtripFlightModifyFlightInfoRq Setter
// 改签询价入参
func (r *AlitripBtripFlightDistributionChangeNewqueryAPIRequest) SetParamBtripFlightModifyFlightInfoRq(_paramBtripFlightModifyFlightInfoRq *BtripFlightModifyFlightInfoRq) error {
	r._paramBtripFlightModifyFlightInfoRq = _paramBtripFlightModifyFlightInfoRq
	r.Set("param_btrip_flight_modify_flight_info_rq", _paramBtripFlightModifyFlightInfoRq)
	return nil
}

// GetParamBtripFlightModifyFlightInfoRq ParamBtripFlightModifyFlightInfoRq Getter
func (r AlitripBtripFlightDistributionChangeNewqueryAPIRequest) GetParamBtripFlightModifyFlightInfoRq() *BtripFlightModifyFlightInfoRq {
	return r._paramBtripFlightModifyFlightInfoRq
}
