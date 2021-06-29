package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商确认机票订单接口 API请求
taobao.alitrip.seller.order.confirm

此接口用于代理商确认机票订单。
*/
type TaobaoAlitripSellerOrderConfirmRequest struct {
    model.Params
    // 请求参数
    tripConfirmOrderParam   *TripConfirmOrderParam
}

// 初始化TaobaoAlitripSellerOrderConfirmRequest对象
func NewTaobaoAlitripSellerOrderConfirmRequest() *TaobaoAlitripSellerOrderConfirmRequest{
    return &TaobaoAlitripSellerOrderConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerOrderConfirmRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.order.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerOrderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TripConfirmOrderParam Setter
// 请求参数
func (r *TaobaoAlitripSellerOrderConfirmRequest) SetTripConfirmOrderParam(tripConfirmOrderParam *TripConfirmOrderParam) error {
    r.tripConfirmOrderParam = tripConfirmOrderParam
    r.Set("trip_confirm_order_param", tripConfirmOrderParam)
    return nil
}

// TripConfirmOrderParam Getter
func (r TaobaoAlitripSellerOrderConfirmRequest) GetTripConfirmOrderParam() *TripConfirmOrderParam {
    return r.tripConfirmOrderParam
}
