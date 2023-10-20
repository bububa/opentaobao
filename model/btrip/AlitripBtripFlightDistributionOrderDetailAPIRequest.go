package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionOrderDetailAPIRequest 商旅机票分销订单详情接口 API请求
// alitrip.btrip.flight.distribution.order.detail
//
// 商旅机票分销订单详情接口
type AlitripBtripFlightDistributionOrderDetailAPIRequest struct {
	model.Params
	// 订单详情入参
	_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq
}

// NewAlitripBtripFlightDistributionOrderDetailRequest 初始化AlitripBtripFlightDistributionOrderDetailAPIRequest对象
func NewAlitripBtripFlightDistributionOrderDetailRequest() *AlitripBtripFlightDistributionOrderDetailAPIRequest {
	return &AlitripBtripFlightDistributionOrderDetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionOrderDetailAPIRequest) Reset() {
	r._paramBtripFlightOrderOperateCommonRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionOrderDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionOrderDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionOrderDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightOrderOperateCommonRq is ParamBtripFlightOrderOperateCommonRq Setter
// 订单详情入参
func (r *AlitripBtripFlightDistributionOrderDetailAPIRequest) SetParamBtripFlightOrderOperateCommonRq(_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq) error {
	r._paramBtripFlightOrderOperateCommonRq = _paramBtripFlightOrderOperateCommonRq
	r.Set("param_btrip_flight_order_operate_common_rq", _paramBtripFlightOrderOperateCommonRq)
	return nil
}

// GetParamBtripFlightOrderOperateCommonRq ParamBtripFlightOrderOperateCommonRq Getter
func (r AlitripBtripFlightDistributionOrderDetailAPIRequest) GetParamBtripFlightOrderOperateCommonRq() *BtripFlightOrderOperateCommonRq {
	return r._paramBtripFlightOrderOperateCommonRq
}

var poolAlitripBtripFlightDistributionOrderDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionOrderDetailRequest()
	},
}

// GetAlitripBtripFlightDistributionOrderDetailRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionOrderDetailAPIRequest
func GetAlitripBtripFlightDistributionOrderDetailAPIRequest() *AlitripBtripFlightDistributionOrderDetailAPIRequest {
	return poolAlitripBtripFlightDistributionOrderDetailAPIRequest.Get().(*AlitripBtripFlightDistributionOrderDetailAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionOrderDetailAPIRequest 将 AlitripBtripFlightDistributionOrderDetailAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionOrderDetailAPIRequest(v *AlitripBtripFlightDistributionOrderDetailAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionOrderDetailAPIRequest.Put(v)
}
