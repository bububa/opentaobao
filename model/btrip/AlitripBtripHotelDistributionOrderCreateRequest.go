package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店分销-创建订单 APIRequest
alitrip.btrip.hotel.distribution.order.create

商旅酒店分销-创建订单
*/
type AlitripBtripHotelDistributionOrderCreateRequest struct {
    model.Params

    // 创建订单请求入参
    paramBtripHotelCreateOrderRq   *BtripHotelCreateOrderRq 

}

func NewAlitripBtripHotelDistributionOrderCreateRequest() *AlitripBtripHotelDistributionOrderCreateRequest{
    return &AlitripBtripHotelDistributionOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripHotelDistributionOrderCreateRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.order.create"
}

func (r AlitripBtripHotelDistributionOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripHotelDistributionOrderCreateRequest) SetParamBtripHotelCreateOrderRq(paramBtripHotelCreateOrderRq *BtripHotelCreateOrderRq) error {
    r.paramBtripHotelCreateOrderRq = paramBtripHotelCreateOrderRq
    r.Set("param_btrip_hotel_create_order_rq", paramBtripHotelCreateOrderRq)
    return nil
}

func (r AlitripBtripHotelDistributionOrderCreateRequest) GetParamBtripHotelCreateOrderRq() *BtripHotelCreateOrderRq {
    return r.paramBtripHotelCreateOrderRq
}

