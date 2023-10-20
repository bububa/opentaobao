package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundPrecalAPIRequest 商旅机票分销-退票费预计算 API请求
// alitrip.btrip.flight.distribution.refund.precal
//
// 商旅机票分销-退票费预计算
type AlitripBtripFlightDistributionRefundPrecalAPIRequest struct {
	model.Params
	// 退票费预计算请求入参
	_paramBtripFlightRefundPreCalRq *BtripFlightRefundPreCalRq
}

// NewAlitripBtripFlightDistributionRefundPrecalRequest 初始化AlitripBtripFlightDistributionRefundPrecalAPIRequest对象
func NewAlitripBtripFlightDistributionRefundPrecalRequest() *AlitripBtripFlightDistributionRefundPrecalAPIRequest {
	return &AlitripBtripFlightDistributionRefundPrecalAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionRefundPrecalAPIRequest) Reset() {
	r._paramBtripFlightRefundPreCalRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionRefundPrecalAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.precal"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionRefundPrecalAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionRefundPrecalAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightRefundPreCalRq is ParamBtripFlightRefundPreCalRq Setter
// 退票费预计算请求入参
func (r *AlitripBtripFlightDistributionRefundPrecalAPIRequest) SetParamBtripFlightRefundPreCalRq(_paramBtripFlightRefundPreCalRq *BtripFlightRefundPreCalRq) error {
	r._paramBtripFlightRefundPreCalRq = _paramBtripFlightRefundPreCalRq
	r.Set("param_btrip_flight_refund_pre_cal_rq", _paramBtripFlightRefundPreCalRq)
	return nil
}

// GetParamBtripFlightRefundPreCalRq ParamBtripFlightRefundPreCalRq Getter
func (r AlitripBtripFlightDistributionRefundPrecalAPIRequest) GetParamBtripFlightRefundPreCalRq() *BtripFlightRefundPreCalRq {
	return r._paramBtripFlightRefundPreCalRq
}

var poolAlitripBtripFlightDistributionRefundPrecalAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionRefundPrecalRequest()
	},
}

// GetAlitripBtripFlightDistributionRefundPrecalRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionRefundPrecalAPIRequest
func GetAlitripBtripFlightDistributionRefundPrecalAPIRequest() *AlitripBtripFlightDistributionRefundPrecalAPIRequest {
	return poolAlitripBtripFlightDistributionRefundPrecalAPIRequest.Get().(*AlitripBtripFlightDistributionRefundPrecalAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionRefundPrecalAPIRequest 将 AlitripBtripFlightDistributionRefundPrecalAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionRefundPrecalAPIRequest(v *AlitripBtripFlightDistributionRefundPrecalAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionRefundPrecalAPIRequest.Put(v)
}
