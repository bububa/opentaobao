package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionrefunddetailAPIRequest 商旅机票退票详情接口 API请求
// alitrip.btrip.flight.distribution.refund.detail
//
// 商旅机票分销退票详情
type AlitripbtripflightdistributionrefunddetailAPIRequest struct {
	model.Params
	// 获取退票单详情入参
	_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq
}

// NewAlitripbtripflightdistributionrefunddetailRequest 初始化AlitripbtripflightdistributionrefunddetailAPIRequest对象
func NewAlitripbtripflightdistributionrefunddetailRequest() *AlitripbtripflightdistributionrefunddetailAPIRequest {
	return &AlitripbtripflightdistributionrefunddetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionrefunddetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionrefunddetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionrefunddetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightOrderOperateCommonRq is ParamBtripFlightOrderOperateCommonRq Setter
// 获取退票单详情入参
func (r *AlitripbtripflightdistributionrefunddetailAPIRequest) SetParamBtripFlightOrderOperateCommonRq(_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq) error {
	r._paramBtripFlightOrderOperateCommonRq = _paramBtripFlightOrderOperateCommonRq
	r.Set("param_btrip_flight_order_operate_common_rq", _paramBtripFlightOrderOperateCommonRq)
	return nil
}

// GetParamBtripFlightOrderOperateCommonRq ParamBtripFlightOrderOperateCommonRq Getter
func (r AlitripbtripflightdistributionrefunddetailAPIRequest) GetParamBtripFlightOrderOperateCommonRq() *BtripFlightOrderOperateCommonRq {
	return r._paramBtripFlightOrderOperateCommonRq
}
