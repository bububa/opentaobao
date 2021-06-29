package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商确认机票订单接口 APIRequest
taobao.alitrip.seller.order.confirm

此接口用于代理商确认机票订单。
*/
type TaobaoAlitripSellerOrderConfirmRequest struct {
    model.Params

    // 请求参数
    tripConfirmOrderParam   *TripConfirmOrderParam 

}

func NewTaobaoAlitripSellerOrderConfirmRequest() *TaobaoAlitripSellerOrderConfirmRequest{
    return &TaobaoAlitripSellerOrderConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripSellerOrderConfirmRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.order.confirm"
}

func (r TaobaoAlitripSellerOrderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripSellerOrderConfirmRequest) SetTripConfirmOrderParam(tripConfirmOrderParam *TripConfirmOrderParam) error {
    r.tripConfirmOrderParam = tripConfirmOrderParam
    r.Set("trip_confirm_order_param", tripConfirmOrderParam)
    return nil
}

func (r TaobaoAlitripSellerOrderConfirmRequest) GetTripConfirmOrderParam() *TripConfirmOrderParam {
    return r.tripConfirmOrderParam
}

