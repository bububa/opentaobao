package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionflightlistAPIRequest 商旅机票航班列表接口 API请求
// alitrip.btrip.flight.distribution.flightlist
//
// 商旅机票航班列表接口，用于分销询价
type AlitripbtripflightdistributionflightlistAPIRequest struct {
	model.Params
	// 机票搜索入参
	_paramFlightSearchListRQ *BtripFlightSearchListRq
}

// NewAlitripbtripflightdistributionflightlistRequest 初始化AlitripbtripflightdistributionflightlistAPIRequest对象
func NewAlitripbtripflightdistributionflightlistRequest() *AlitripbtripflightdistributionflightlistAPIRequest {
	return &AlitripbtripflightdistributionflightlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionflightlistAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.flightlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionflightlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionflightlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamFlightSearchListRQ is ParamFlightSearchListRQ Setter
// 机票搜索入参
func (r *AlitripbtripflightdistributionflightlistAPIRequest) SetParamFlightSearchListRQ(_paramFlightSearchListRQ *BtripFlightSearchListRq) error {
	r._paramFlightSearchListRQ = _paramFlightSearchListRQ
	r.Set("param_flight_search_list_r_q", _paramFlightSearchListRQ)
	return nil
}

// GetParamFlightSearchListRQ ParamFlightSearchListRQ Getter
func (r AlitripbtripflightdistributionflightlistAPIRequest) GetParamFlightSearchListRQ() *BtripFlightSearchListRq {
	return r._paramFlightSearchListRQ
}
