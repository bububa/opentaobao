package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店API分销下单前校验 API请求
alitrip.btrip.hotel.distribution.order.validate

商旅酒店API分销下单前校验
*/
type AlitripBtripHotelDistributionOrderValidateRequest struct {
    model.Params
    // 下单前校验入参
    paramBtripHotelValidateOrderRq   *BtripHotelValidateOrderRq
}

// 初始化AlitripBtripHotelDistributionOrderValidateRequest对象
func NewAlitripBtripHotelDistributionOrderValidateRequest() *AlitripBtripHotelDistributionOrderValidateRequest{
    return &AlitripBtripHotelDistributionOrderValidateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderValidateRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.order.validate"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderValidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBtripHotelValidateOrderRq Setter
// 下单前校验入参
func (r *AlitripBtripHotelDistributionOrderValidateRequest) SetParamBtripHotelValidateOrderRq(paramBtripHotelValidateOrderRq *BtripHotelValidateOrderRq) error {
    r.paramBtripHotelValidateOrderRq = paramBtripHotelValidateOrderRq
    r.Set("param_btrip_hotel_validate_order_rq", paramBtripHotelValidateOrderRq)
    return nil
}

// ParamBtripHotelValidateOrderRq Getter
func (r AlitripBtripHotelDistributionOrderValidateRequest) GetParamBtripHotelValidateOrderRq() *BtripHotelValidateOrderRq {
    return r.paramBtripHotelValidateOrderRq
}
