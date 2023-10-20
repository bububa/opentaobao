package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtriphoteldistributionorderdetailAPIRequest 商旅酒店API分销查询订单详情 API请求
// alitrip.btrip.hotel.distribution.order.detail
//
// 商旅酒店API分销查询订单详情
type AlitripbtriphoteldistributionorderdetailAPIRequest struct {
	model.Params
	// 订单详情接口入参
	_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq
}

// NewAlitripbtriphoteldistributionorderdetailRequest 初始化AlitripbtriphoteldistributionorderdetailAPIRequest对象
func NewAlitripbtriphoteldistributionorderdetailRequest() *AlitripbtriphoteldistributionorderdetailAPIRequest {
	return &AlitripbtriphoteldistributionorderdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtriphoteldistributionorderdetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtriphoteldistributionorderdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtriphoteldistributionorderdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripHotelOrderOperateRq is ParamBtripHotelOrderOperateRq Setter
// 订单详情接口入参
func (r *AlitripbtriphoteldistributionorderdetailAPIRequest) SetParamBtripHotelOrderOperateRq(_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq) error {
	r._paramBtripHotelOrderOperateRq = _paramBtripHotelOrderOperateRq
	r.Set("param_btrip_hotel_order_operate_rq", _paramBtripHotelOrderOperateRq)
	return nil
}

// GetParamBtripHotelOrderOperateRq ParamBtripHotelOrderOperateRq Getter
func (r AlitripbtriphoteldistributionorderdetailAPIRequest) GetParamBtripHotelOrderOperateRq() *BtripHotelOrderOperateRq {
	return r._paramBtripHotelOrderOperateRq
}
