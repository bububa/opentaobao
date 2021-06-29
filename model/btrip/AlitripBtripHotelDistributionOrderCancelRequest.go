package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店API分销取消订单 APIRequest
alitrip.btrip.hotel.distribution.order.cancel

商旅酒店API分销取消订单
*/
type AlitripBtripHotelDistributionOrderCancelRequest struct {
    model.Params

    // 取消订单接口入参
    paramBtripHotelOrderOperateRq   *BtripHotelOrderOperateRq 

}

func NewAlitripBtripHotelDistributionOrderCancelRequest() *AlitripBtripHotelDistributionOrderCancelRequest{
    return &AlitripBtripHotelDistributionOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripHotelDistributionOrderCancelRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.order.cancel"
}

func (r AlitripBtripHotelDistributionOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripHotelDistributionOrderCancelRequest) SetParamBtripHotelOrderOperateRq(paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq) error {
    r.paramBtripHotelOrderOperateRq = paramBtripHotelOrderOperateRq
    r.Set("param_btrip_hotel_order_operate_rq", paramBtripHotelOrderOperateRq)
    return nil
}

func (r AlitripBtripHotelDistributionOrderCancelRequest) GetParamBtripHotelOrderOperateRq() *BtripHotelOrderOperateRq {
    return r.paramBtripHotelOrderOperateRq
}

