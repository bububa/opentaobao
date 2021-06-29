package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店分销订单支付 API请求
alitrip.btrip.hotel.distribution.order.pay

商旅酒店分销订单支付
*/
type AlitripBtripHotelDistributionOrderPayRequest struct {
    model.Params
    // 通知商旅支付成功接口参数
    paramBtripHotelOrderOperateRq   *BtripHotelOrderOperateRq
}

// 初始化AlitripBtripHotelDistributionOrderPayRequest对象
func NewAlitripBtripHotelDistributionOrderPayRequest() *AlitripBtripHotelDistributionOrderPayRequest{
    return &AlitripBtripHotelDistributionOrderPayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderPayRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.order.pay"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderPayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBtripHotelOrderOperateRq Setter
// 通知商旅支付成功接口参数
func (r *AlitripBtripHotelDistributionOrderPayRequest) SetParamBtripHotelOrderOperateRq(paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq) error {
    r.paramBtripHotelOrderOperateRq = paramBtripHotelOrderOperateRq
    r.Set("param_btrip_hotel_order_operate_rq", paramBtripHotelOrderOperateRq)
    return nil
}

// ParamBtripHotelOrderOperateRq Getter
func (r AlitripBtripHotelDistributionOrderPayRequest) GetParamBtripHotelOrderOperateRq() *BtripHotelOrderOperateRq {
    return r.paramBtripHotelOrderOperateRq
}
