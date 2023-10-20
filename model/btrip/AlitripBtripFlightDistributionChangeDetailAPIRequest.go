package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeDetailAPIRequest 商旅机票改签详情接口 API请求
// alitrip.btrip.flight.distribution.change.detail
//
// 商旅机票改签详情接口
type AlitripBtripFlightDistributionChangeDetailAPIRequest struct {
	model.Params
	// 获取改签单详情入参
	_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq
}

// NewAlitripBtripFlightDistributionChangeDetailRequest 初始化AlitripBtripFlightDistributionChangeDetailAPIRequest对象
func NewAlitripBtripFlightDistributionChangeDetailRequest() *AlitripBtripFlightDistributionChangeDetailAPIRequest {
	return &AlitripBtripFlightDistributionChangeDetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionChangeDetailAPIRequest) Reset() {
	r._paramBtripFlightOrderOperateCommonRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangeDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangeDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionChangeDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightOrderOperateCommonRq is ParamBtripFlightOrderOperateCommonRq Setter
// 获取改签单详情入参
func (r *AlitripBtripFlightDistributionChangeDetailAPIRequest) SetParamBtripFlightOrderOperateCommonRq(_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq) error {
	r._paramBtripFlightOrderOperateCommonRq = _paramBtripFlightOrderOperateCommonRq
	r.Set("param_btrip_flight_order_operate_common_rq", _paramBtripFlightOrderOperateCommonRq)
	return nil
}

// GetParamBtripFlightOrderOperateCommonRq ParamBtripFlightOrderOperateCommonRq Getter
func (r AlitripBtripFlightDistributionChangeDetailAPIRequest) GetParamBtripFlightOrderOperateCommonRq() *BtripFlightOrderOperateCommonRq {
	return r._paramBtripFlightOrderOperateCommonRq
}

var poolAlitripBtripFlightDistributionChangeDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionChangeDetailRequest()
	},
}

// GetAlitripBtripFlightDistributionChangeDetailRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionChangeDetailAPIRequest
func GetAlitripBtripFlightDistributionChangeDetailAPIRequest() *AlitripBtripFlightDistributionChangeDetailAPIRequest {
	return poolAlitripBtripFlightDistributionChangeDetailAPIRequest.Get().(*AlitripBtripFlightDistributionChangeDetailAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionChangeDetailAPIRequest 将 AlitripBtripFlightDistributionChangeDetailAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangeDetailAPIRequest(v *AlitripBtripFlightDistributionChangeDetailAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangeDetailAPIRequest.Put(v)
}
