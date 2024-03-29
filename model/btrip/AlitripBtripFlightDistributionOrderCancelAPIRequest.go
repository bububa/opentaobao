package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionOrderCancelAPIRequest 商旅机票分销-取消订单 API请求
// alitrip.btrip.flight.distribution.order.cancel
//
// 商旅机票分销取消订单
type AlitripBtripFlightDistributionOrderCancelAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramBtripFlightOrderOperateCommonRq *BtripFlightCancelOrderRq
}

// NewAlitripBtripFlightDistributionOrderCancelRequest 初始化AlitripBtripFlightDistributionOrderCancelAPIRequest对象
func NewAlitripBtripFlightDistributionOrderCancelRequest() *AlitripBtripFlightDistributionOrderCancelAPIRequest {
	return &AlitripBtripFlightDistributionOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionOrderCancelAPIRequest) Reset() {
	r._paramBtripFlightOrderOperateCommonRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionOrderCancelAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightOrderOperateCommonRq is ParamBtripFlightOrderOperateCommonRq Setter
// 系统自动生成
func (r *AlitripBtripFlightDistributionOrderCancelAPIRequest) SetParamBtripFlightOrderOperateCommonRq(_paramBtripFlightOrderOperateCommonRq *BtripFlightCancelOrderRq) error {
	r._paramBtripFlightOrderOperateCommonRq = _paramBtripFlightOrderOperateCommonRq
	r.Set("param_btrip_flight_order_operate_common_rq", _paramBtripFlightOrderOperateCommonRq)
	return nil
}

// GetParamBtripFlightOrderOperateCommonRq ParamBtripFlightOrderOperateCommonRq Getter
func (r AlitripBtripFlightDistributionOrderCancelAPIRequest) GetParamBtripFlightOrderOperateCommonRq() *BtripFlightCancelOrderRq {
	return r._paramBtripFlightOrderOperateCommonRq
}

var poolAlitripBtripFlightDistributionOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionOrderCancelRequest()
	},
}

// GetAlitripBtripFlightDistributionOrderCancelRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionOrderCancelAPIRequest
func GetAlitripBtripFlightDistributionOrderCancelAPIRequest() *AlitripBtripFlightDistributionOrderCancelAPIRequest {
	return poolAlitripBtripFlightDistributionOrderCancelAPIRequest.Get().(*AlitripBtripFlightDistributionOrderCancelAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionOrderCancelAPIRequest 将 AlitripBtripFlightDistributionOrderCancelAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionOrderCancelAPIRequest(v *AlitripBtripFlightDistributionOrderCancelAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionOrderCancelAPIRequest.Put(v)
}
