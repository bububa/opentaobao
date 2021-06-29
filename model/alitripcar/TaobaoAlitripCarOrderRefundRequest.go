package alitripcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户投诉达成一致后给用户退款 API请求
taobao.alitrip.car.order.refund

用户投诉后，供应商客服与客户沟通达成一致后通知飞猪给客户退款。退款金额以接口回调金额为准。
*/
type TaobaoAlitripCarOrderRefundRequest struct {
    model.Params
    // 退款对象
    _paramOrderRefund   *OrderRefund
}

// 初始化TaobaoAlitripCarOrderRefundRequest对象
func NewTaobaoAlitripCarOrderRefundRequest() *TaobaoAlitripCarOrderRefundRequest{
    return &TaobaoAlitripCarOrderRefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderRefundRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.order.refund"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOrderRefund Setter
// 退款对象
func (r *TaobaoAlitripCarOrderRefundRequest) SetParamOrderRefund(_paramOrderRefund *OrderRefund) error {
    r._paramOrderRefund = _paramOrderRefund
    r.Set("param_order_refund", _paramOrderRefund)
    return nil
}

// ParamOrderRefund Getter
func (r TaobaoAlitripCarOrderRefundRequest) GetParamOrderRefund() *OrderRefund {
    return r._paramOrderRefund
}
