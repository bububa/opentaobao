package alitripcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户投诉达成一致后给用户退款 APIRequest
taobao.alitrip.car.order.refund

用户投诉后，供应商客服与客户沟通达成一致后通知飞猪给客户退款。退款金额以接口回调金额为准。
*/
type TaobaoAlitripCarOrderRefundRequest struct {
    model.Params

    // 退款对象
    paramOrderRefund   *OrderRefund 

}

func NewTaobaoAlitripCarOrderRefundRequest() *TaobaoAlitripCarOrderRefundRequest{
    return &TaobaoAlitripCarOrderRefundRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripCarOrderRefundRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.order.refund"
}

func (r TaobaoAlitripCarOrderRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripCarOrderRefundRequest) SetParamOrderRefund(paramOrderRefund *OrderRefund) error {
    r.paramOrderRefund = paramOrderRefund
    r.Set("param_order_refund", paramOrderRefund)
    return nil
}

func (r TaobaoAlitripCarOrderRefundRequest) GetParamOrderRefund() *OrderRefund {
    return r.paramOrderRefund
}

