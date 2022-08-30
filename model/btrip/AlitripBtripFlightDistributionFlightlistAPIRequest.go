package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionFlightlistAPIRequest 商旅机票航班列表接口 API请求
// alitrip.btrip.flight.distribution.flightlist
//
// 商旅机票航班列表接口，用于分销询价
type AlitripBtripFlightDistributionFlightlistAPIRequest struct {
	model.Params
	// 机票搜索入参
	_paramFlightSearchListRQ *BtripFlightSearchListRq
}

// NewAlitripBtripFlightDistributionFlightlistRequest 初始化AlitripBtripFlightDistributionFlightlistAPIRequest对象
func NewAlitripBtripFlightDistributionFlightlistRequest() *AlitripBtripFlightDistributionFlightlistAPIRequest {
	return &AlitripBtripFlightDistributionFlightlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionFlightlistAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.flightlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionFlightlistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamFlightSearchListRQ is ParamFlightSearchListRQ Setter
// 机票搜索入参
func (r *AlitripBtripFlightDistributionFlightlistAPIRequest) SetParamFlightSearchListRQ(_paramFlightSearchListRQ *BtripFlightSearchListRq) error {
	r._paramFlightSearchListRQ = _paramFlightSearchListRQ
	r.Set("param_flight_search_list_r_q", _paramFlightSearchListRQ)
	return nil
}

// GetParamFlightSearchListRQ ParamFlightSearchListRQ Getter
func (r AlitripBtripFlightDistributionFlightlistAPIRequest) GetParamFlightSearchListRQ() *BtripFlightSearchListRq {
	return r._paramFlightSearchListRQ
}
