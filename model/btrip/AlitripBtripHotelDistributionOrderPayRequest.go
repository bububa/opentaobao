package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店分销订单支付 APIRequest
alitrip.btrip.hotel.distribution.order.pay

商旅酒店分销订单支付
*/
type AlitripBtripHotelDistributionOrderPayRequest struct {
    model.Params

    // 通知商旅支付成功接口参数
    paramBtripHotelOrderOperateRq   *BtripHotelOrderOperateRq 

}

func NewAlitripBtripHotelDistributionOrderPayRequest() *AlitripBtripHotelDistributionOrderPayRequest{
    return &AlitripBtripHotelDistributionOrderPayRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripHotelDistributionOrderPayRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.order.pay"
}

func (r AlitripBtripHotelDistributionOrderPayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripHotelDistributionOrderPayRequest) SetParamBtripHotelOrderOperateRq(paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq) error {
    r.paramBtripHotelOrderOperateRq = paramBtripHotelOrderOperateRq
    r.Set("param_btrip_hotel_order_operate_rq", paramBtripHotelOrderOperateRq)
    return nil
}

func (r AlitripBtripHotelDistributionOrderPayRequest) GetParamBtripHotelOrderOperateRq() *BtripHotelOrderOperateRq {
    return r.paramBtripHotelOrderOperateRq
}

