package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionOrderCancelAPIRequest 商旅酒店API分销取消订单 API请求
// alitrip.btrip.hotel.distribution.order.cancel
//
// 商旅酒店API分销取消订单
type AlitripBtripHotelDistributionOrderCancelAPIRequest struct {
	model.Params
	// 取消订单接口入参
	_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq
}

// NewAlitripBtripHotelDistributionOrderCancelRequest 初始化AlitripBtripHotelDistributionOrderCancelAPIRequest对象
func NewAlitripBtripHotelDistributionOrderCancelRequest() *AlitripBtripHotelDistributionOrderCancelAPIRequest {
	return &AlitripBtripHotelDistributionOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripHotelDistributionOrderCancelAPIRequest) Reset() {
	r._paramBtripHotelOrderOperateRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderCancelAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripHotelDistributionOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripHotelOrderOperateRq is ParamBtripHotelOrderOperateRq Setter
// 取消订单接口入参
func (r *AlitripBtripHotelDistributionOrderCancelAPIRequest) SetParamBtripHotelOrderOperateRq(_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq) error {
	r._paramBtripHotelOrderOperateRq = _paramBtripHotelOrderOperateRq
	r.Set("param_btrip_hotel_order_operate_rq", _paramBtripHotelOrderOperateRq)
	return nil
}

// GetParamBtripHotelOrderOperateRq ParamBtripHotelOrderOperateRq Getter
func (r AlitripBtripHotelDistributionOrderCancelAPIRequest) GetParamBtripHotelOrderOperateRq() *BtripHotelOrderOperateRq {
	return r._paramBtripHotelOrderOperateRq
}

var poolAlitripBtripHotelDistributionOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripHotelDistributionOrderCancelRequest()
	},
}

// GetAlitripBtripHotelDistributionOrderCancelRequest 从 sync.Pool 获取 AlitripBtripHotelDistributionOrderCancelAPIRequest
func GetAlitripBtripHotelDistributionOrderCancelAPIRequest() *AlitripBtripHotelDistributionOrderCancelAPIRequest {
	return poolAlitripBtripHotelDistributionOrderCancelAPIRequest.Get().(*AlitripBtripHotelDistributionOrderCancelAPIRequest)
}

// ReleaseAlitripBtripHotelDistributionOrderCancelAPIRequest 将 AlitripBtripHotelDistributionOrderCancelAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripHotelDistributionOrderCancelAPIRequest(v *AlitripBtripHotelDistributionOrderCancelAPIRequest) {
	v.Reset()
	poolAlitripBtripHotelDistributionOrderCancelAPIRequest.Put(v)
}
