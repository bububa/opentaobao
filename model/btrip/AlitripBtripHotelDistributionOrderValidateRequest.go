package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店API分销下单前校验 APIRequest
alitrip.btrip.hotel.distribution.order.validate

商旅酒店API分销下单前校验
*/
type AlitripBtripHotelDistributionOrderValidateRequest struct {
    model.Params

    // 下单前校验入参
    paramBtripHotelValidateOrderRq   *BtripHotelValidateOrderRq 

}

func NewAlitripBtripHotelDistributionOrderValidateRequest() *AlitripBtripHotelDistributionOrderValidateRequest{
    return &AlitripBtripHotelDistributionOrderValidateRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripHotelDistributionOrderValidateRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.order.validate"
}

func (r AlitripBtripHotelDistributionOrderValidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripHotelDistributionOrderValidateRequest) SetParamBtripHotelValidateOrderRq(paramBtripHotelValidateOrderRq *BtripHotelValidateOrderRq) error {
    r.paramBtripHotelValidateOrderRq = paramBtripHotelValidateOrderRq
    r.Set("param_btrip_hotel_validate_order_rq", paramBtripHotelValidateOrderRq)
    return nil
}

func (r AlitripBtripHotelDistributionOrderValidateRequest) GetParamBtripHotelValidateOrderRq() *BtripHotelValidateOrderRq {
    return r.paramBtripHotelValidateOrderRq
}

