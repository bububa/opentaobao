package btrip

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionOrderCancelAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionOrderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
