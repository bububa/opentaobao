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
type AlitripBtripHotelDistributionOrderValidateAPIRequest struct {
    model.Params
    // 下单前校验入参
    _paramBtripHotelValidateOrderRq   *BtripHotelValidateOrderRq
}

// 初始化AlitripBtripHotelDistributionOrderValidateAPIRequest对象
func NewAlitripBtripHotelDistributionOrderValidateRequest() *AlitripBtripHotelDistributionOrderValidateAPIRequest{
    return &AlitripBtripHotelDistributionOrderValidateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderValidateAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.order.validate"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderValidateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBtripHotelValidateOrderRq Setter
// 下单前校验入参
func (r *AlitripBtripHotelDistributionOrderValidateAPIRequest) SetParamBtripHotelValidateOrderRq(_paramBtripHotelValidateOrderRq *BtripHotelValidateOrderRq) error {
    r._paramBtripHotelValidateOrderRq = _paramBtripHotelValidateOrderRq
    r.Set("param_btrip_hotel_validate_order_rq", _paramBtripHotelValidateOrderRq)
    return nil
}

// ParamBtripHotelValidateOrderRq Getter
func (r AlitripBtripHotelDistributionOrderValidateAPIRequest) GetParamBtripHotelValidateOrderRq() *BtripHotelValidateOrderRq {
    return r._paramBtripHotelValidateOrderRq
}
