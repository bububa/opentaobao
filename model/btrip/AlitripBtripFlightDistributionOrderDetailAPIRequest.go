package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionorderdetailAPIRequest 商旅机票分销订单详情接口 API请求
// alitrip.btrip.flight.distribution.order.detail
//
// 商旅机票分销订单详情接口
type AlitripbtripflightdistributionorderdetailAPIRequest struct {
	model.Params
	// 订单详情入参
	_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq
}

// NewAlitripbtripflightdistributionorderdetailRequest 初始化AlitripbtripflightdistributionorderdetailAPIRequest对象
func NewAlitripbtripflightdistributionorderdetailRequest() *AlitripbtripflightdistributionorderdetailAPIRequest {
	return &AlitripbtripflightdistributionorderdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionorderdetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionorderdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionorderdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightOrderOperateCommonRq is ParamBtripFlightOrderOperateCommonRq Setter
// 订单详情入参
func (r *AlitripbtripflightdistributionorderdetailAPIRequest) SetParamBtripFlightOrderOperateCommonRq(_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq) error {
	r._paramBtripFlightOrderOperateCommonRq = _paramBtripFlightOrderOperateCommonRq
	r.Set("param_btrip_flight_order_operate_common_rq", _paramBtripFlightOrderOperateCommonRq)
	return nil
}

// GetParamBtripFlightOrderOperateCommonRq ParamBtripFlightOrderOperateCommonRq Getter
func (r AlitripbtripflightdistributionorderdetailAPIRequest) GetParamBtripFlightOrderOperateCommonRq() *BtripFlightOrderOperateCommonRq {
	return r._paramBtripFlightOrderOperateCommonRq
}
