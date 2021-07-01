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
type TaobaoAlitripSellerOrderConfirmAPIRequest struct {
    model.Params
    // 请求参数
    _tripConfirmOrderParam   *TripConfirmOrderParam
}

// 初始化TaobaoAlitripSellerOrderConfirmAPIRequest对象
func NewTaobaoAlitripSellerOrderConfirmRequest() *TaobaoAlitripSellerOrderConfirmAPIRequest{
    return &TaobaoAlitripSellerOrderConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerOrderConfirmAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.order.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerOrderConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TripConfirmOrderParam Setter
// 请求参数
func (r *TaobaoAlitripSellerOrderConfirmAPIRequest) SetTripConfirmOrderParam(_tripConfirmOrderParam *TripConfirmOrderParam) error {
    r._tripConfirmOrderParam = _tripConfirmOrderParam
    r.Set("trip_confirm_order_param", _tripConfirmOrderParam)
    return nil
}

// TripConfirmOrderParam Getter
func (r TaobaoAlitripSellerOrderConfirmAPIRequest) GetTripConfirmOrderParam() *TripConfirmOrderParam {
    return r._tripConfirmOrderParam
}
