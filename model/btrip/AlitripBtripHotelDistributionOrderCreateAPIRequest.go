package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionOrderCreateAPIRequest 商旅酒店分销-创建订单 API请求
// alitrip.btrip.hotel.distribution.order.create
//
// 商旅酒店分销-创建订单
type AlitripBtripHotelDistributionOrderCreateAPIRequest struct {
	model.Params
	// 创建订单请求入参
	_paramBtripHotelCreateOrderRq *BtripHotelCreateOrderRq
}

// NewAlitripBtripHotelDistributionOrderCreateRequest 初始化AlitripBtripHotelDistributionOrderCreateAPIRequest对象
func NewAlitripBtripHotelDistributionOrderCreateRequest() *AlitripBtripHotelDistributionOrderCreateAPIRequest {
	return &AlitripBtripHotelDistributionOrderCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripHotelDistributionOrderCreateAPIRequest) Reset() {
	r._paramBtripHotelCreateOrderRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderCreateAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripHotelDistributionOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripHotelCreateOrderRq is ParamBtripHotelCreateOrderRq Setter
// 创建订单请求入参
func (r *AlitripBtripHotelDistributionOrderCreateAPIRequest) SetParamBtripHotelCreateOrderRq(_paramBtripHotelCreateOrderRq *BtripHotelCreateOrderRq) error {
	r._paramBtripHotelCreateOrderRq = _paramBtripHotelCreateOrderRq
	r.Set("param_btrip_hotel_create_order_rq", _paramBtripHotelCreateOrderRq)
	return nil
}

// GetParamBtripHotelCreateOrderRq ParamBtripHotelCreateOrderRq Getter
func (r AlitripBtripHotelDistributionOrderCreateAPIRequest) GetParamBtripHotelCreateOrderRq() *BtripHotelCreateOrderRq {
	return r._paramBtripHotelCreateOrderRq
}

var poolAlitripBtripHotelDistributionOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripHotelDistributionOrderCreateRequest()
	},
}

// GetAlitripBtripHotelDistributionOrderCreateRequest 从 sync.Pool 获取 AlitripBtripHotelDistributionOrderCreateAPIRequest
func GetAlitripBtripHotelDistributionOrderCreateAPIRequest() *AlitripBtripHotelDistributionOrderCreateAPIRequest {
	return poolAlitripBtripHotelDistributionOrderCreateAPIRequest.Get().(*AlitripBtripHotelDistributionOrderCreateAPIRequest)
}

// ReleaseAlitripBtripHotelDistributionOrderCreateAPIRequest 将 AlitripBtripHotelDistributionOrderCreateAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripHotelDistributionOrderCreateAPIRequest(v *AlitripBtripHotelDistributionOrderCreateAPIRequest) {
	v.Reset()
	poolAlitripBtripHotelDistributionOrderCreateAPIRequest.Put(v)
}
