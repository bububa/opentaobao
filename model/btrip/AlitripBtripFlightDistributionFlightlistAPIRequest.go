package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionFlightlistAPIRequest) Reset() {
	r._paramFlightSearchListRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionFlightlistAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.flightlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionFlightlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionFlightlistAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripBtripFlightDistributionFlightlistAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionFlightlistRequest()
	},
}

// GetAlitripBtripFlightDistributionFlightlistRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionFlightlistAPIRequest
func GetAlitripBtripFlightDistributionFlightlistAPIRequest() *AlitripBtripFlightDistributionFlightlistAPIRequest {
	return poolAlitripBtripFlightDistributionFlightlistAPIRequest.Get().(*AlitripBtripFlightDistributionFlightlistAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionFlightlistAPIRequest 将 AlitripBtripFlightDistributionFlightlistAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionFlightlistAPIRequest(v *AlitripBtripFlightDistributionFlightlistAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionFlightlistAPIRequest.Put(v)
}
