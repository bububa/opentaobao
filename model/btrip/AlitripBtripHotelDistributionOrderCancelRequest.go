package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店API分销取消订单 API请求
alitrip.btrip.hotel.distribution.order.cancel

商旅酒店API分销取消订单
*/
type AlitripBtripHotelDistributionOrderCancelRequest struct {
    model.Params
    // 取消订单接口入参
    _paramBtripHotelOrderOperateRq   *BtripHotelOrderOperateRq
}

// 初始化AlitripBtripHotelDistributionOrderCancelRequest对象
func NewAlitripBtripHotelDistributionOrderCancelRequest() *AlitripBtripHotelDistributionOrderCancelRequest{
    return &AlitripBtripHotelDistributionOrderCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderCancelRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.order.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBtripHotelOrderOperateRq Setter
// 取消订单接口入参
func (r *AlitripBtripHotelDistributionOrderCancelRequest) SetParamBtripHotelOrderOperateRq(_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq) error {
    r._paramBtripHotelOrderOperateRq = _paramBtripHotelOrderOperateRq
    r.Set("param_btrip_hotel_order_operate_rq", _paramBtripHotelOrderOperateRq)
    return nil
}

// ParamBtripHotelOrderOperateRq Getter
func (r AlitripBtripHotelDistributionOrderCancelRequest) GetParamBtripHotelOrderOperateRq() *BtripHotelOrderOperateRq {
    return r._paramBtripHotelOrderOperateRq
}
