package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionOrderDetailAPIRequest 商旅酒店API分销查询订单详情 API请求
// alitrip.btrip.hotel.distribution.order.detail
//
// 商旅酒店API分销查询订单详情
type AlitripBtripHotelDistributionOrderDetailAPIRequest struct {
	model.Params
	// 订单详情接口入参
	_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq
}

// NewAlitripBtripHotelDistributionOrderDetailRequest 初始化AlitripBtripHotelDistributionOrderDetailAPIRequest对象
func NewAlitripBtripHotelDistributionOrderDetailRequest() *AlitripBtripHotelDistributionOrderDetailAPIRequest {
	return &AlitripBtripHotelDistributionOrderDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripHotelOrderOperateRq is ParamBtripHotelOrderOperateRq Setter
// 订单详情接口入参
func (r *AlitripBtripHotelDistributionOrderDetailAPIRequest) SetParamBtripHotelOrderOperateRq(_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq) error {
	r._paramBtripHotelOrderOperateRq = _paramBtripHotelOrderOperateRq
	r.Set("param_btrip_hotel_order_operate_rq", _paramBtripHotelOrderOperateRq)
	return nil
}

// GetParamBtripHotelOrderOperateRq ParamBtripHotelOrderOperateRq Getter
func (r AlitripBtripHotelDistributionOrderDetailAPIRequest) GetParamBtripHotelOrderOperateRq() *BtripHotelOrderOperateRq {
	return r._paramBtripHotelOrderOperateRq
}
