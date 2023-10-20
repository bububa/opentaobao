package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionrefundnewdetailAPIRequest 商旅机票退票详情接口V2 API请求
// alitrip.btrip.flight.distribution.refund.newdetail
//
// 商旅机票退票详情接口V2
type AlitripbtripflightdistributionrefundnewdetailAPIRequest struct {
	model.Params
	// 获取退票单详情入参
	_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq
}

// NewAlitripbtripflightdistributionrefundnewdetailRequest 初始化AlitripbtripflightdistributionrefundnewdetailAPIRequest对象
func NewAlitripbtripflightdistributionrefundnewdetailRequest() *AlitripbtripflightdistributionrefundnewdetailAPIRequest {
	return &AlitripbtripflightdistributionrefundnewdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionrefundnewdetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.newdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionrefundnewdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionrefundnewdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightOrderOperateCommonRq is ParamBtripFlightOrderOperateCommonRq Setter
// 获取退票单详情入参
func (r *AlitripbtripflightdistributionrefundnewdetailAPIRequest) SetParamBtripFlightOrderOperateCommonRq(_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq) error {
	r._paramBtripFlightOrderOperateCommonRq = _paramBtripFlightOrderOperateCommonRq
	r.Set("param_btrip_flight_order_operate_common_rq", _paramBtripFlightOrderOperateCommonRq)
	return nil
}

// GetParamBtripFlightOrderOperateCommonRq ParamBtripFlightOrderOperateCommonRq Getter
func (r AlitripbtripflightdistributionrefundnewdetailAPIRequest) GetParamBtripFlightOrderOperateCommonRq() *BtripFlightOrderOperateCommonRq {
	return r._paramBtripFlightOrderOperateCommonRq
}
