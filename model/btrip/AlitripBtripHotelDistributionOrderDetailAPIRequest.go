package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripHotelDistributionOrderDetailAPIRequest) Reset() {
	r._paramBtripHotelOrderOperateRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripHotelDistributionOrderDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripBtripHotelDistributionOrderDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripHotelDistributionOrderDetailRequest()
	},
}

// GetAlitripBtripHotelDistributionOrderDetailRequest 从 sync.Pool 获取 AlitripBtripHotelDistributionOrderDetailAPIRequest
func GetAlitripBtripHotelDistributionOrderDetailAPIRequest() *AlitripBtripHotelDistributionOrderDetailAPIRequest {
	return poolAlitripBtripHotelDistributionOrderDetailAPIRequest.Get().(*AlitripBtripHotelDistributionOrderDetailAPIRequest)
}

// ReleaseAlitripBtripHotelDistributionOrderDetailAPIRequest 将 AlitripBtripHotelDistributionOrderDetailAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripHotelDistributionOrderDetailAPIRequest(v *AlitripBtripHotelDistributionOrderDetailAPIRequest) {
	v.Reset()
	poolAlitripBtripHotelDistributionOrderDetailAPIRequest.Put(v)
}
