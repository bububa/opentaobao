package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店分销-创建订单 API请求
alitrip.btrip.hotel.distribution.order.create

商旅酒店分销-创建订单
*/
type AlitripBtripHotelDistributionOrderCreateRequest struct {
    model.Params
    // 创建订单请求入参
    paramBtripHotelCreateOrderRq   *BtripHotelCreateOrderRq
}

// 初始化AlitripBtripHotelDistributionOrderCreateRequest对象
func NewAlitripBtripHotelDistributionOrderCreateRequest() *AlitripBtripHotelDistributionOrderCreateRequest{
    return &AlitripBtripHotelDistributionOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderCreateRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBtripHotelCreateOrderRq Setter
// 创建订单请求入参
func (r *AlitripBtripHotelDistributionOrderCreateRequest) SetParamBtripHotelCreateOrderRq(paramBtripHotelCreateOrderRq *BtripHotelCreateOrderRq) error {
    r.paramBtripHotelCreateOrderRq = paramBtripHotelCreateOrderRq
    r.Set("param_btrip_hotel_create_order_rq", paramBtripHotelCreateOrderRq)
    return nil
}

// ParamBtripHotelCreateOrderRq Getter
func (r AlitripBtripHotelDistributionOrderCreateRequest) GetParamBtripHotelCreateOrderRq() *BtripHotelCreateOrderRq {
    return r.paramBtripHotelCreateOrderRq
}
