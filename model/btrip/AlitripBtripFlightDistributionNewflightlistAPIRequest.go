package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionnewflightlistAPIRequest 商旅机票航班列表接口，用于分销询价V2 API请求
// alitrip.btrip.flight.distribution.newflightlist
//
// 商旅机票航班列表接口，用于分销询价V2
type AlitripbtripflightdistributionnewflightlistAPIRequest struct {
	model.Params
	// 机票搜索入参
	_paramFlightSearchListRQ *BtripFlightSearchListRq
}

// NewAlitripbtripflightdistributionnewflightlistRequest 初始化AlitripbtripflightdistributionnewflightlistAPIRequest对象
func NewAlitripbtripflightdistributionnewflightlistRequest() *AlitripbtripflightdistributionnewflightlistAPIRequest {
	return &AlitripbtripflightdistributionnewflightlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionnewflightlistAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.newflightlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionnewflightlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionnewflightlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamFlightSearchListRQ is ParamFlightSearchListRQ Setter
// 机票搜索入参
func (r *AlitripbtripflightdistributionnewflightlistAPIRequest) SetParamFlightSearchListRQ(_paramFlightSearchListRQ *BtripFlightSearchListRq) error {
	r._paramFlightSearchListRQ = _paramFlightSearchListRQ
	r.Set("param_flight_search_list_r_q", _paramFlightSearchListRQ)
	return nil
}

// GetParamFlightSearchListRQ ParamFlightSearchListRQ Getter
func (r AlitripbtripflightdistributionnewflightlistAPIRequest) GetParamFlightSearchListRQ() *BtripFlightSearchListRq {
	return r._paramFlightSearchListRQ
}
