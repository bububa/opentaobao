package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
补退接口 API请求
taobao.alitrip.ie.agent.refund.new.multiplerefunds

1. 补退接口， 可以进行多次退款
*/
type TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest struct {
    model.Params
    // 请求参数
    paramRefundOrderMultipleRefundsRq   *RefundOrderMultipleRefundsRq
}

// 初始化TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest对象
func NewTaobaoAlitripIeAgentRefundNewMultiplerefundsRequest() *TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest{
    return &TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.new.multiplerefunds"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRefundOrderMultipleRefundsRq Setter
// 请求参数
func (r *TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest) SetParamRefundOrderMultipleRefundsRq(paramRefundOrderMultipleRefundsRq *RefundOrderMultipleRefundsRq) error {
    r.paramRefundOrderMultipleRefundsRq = paramRefundOrderMultipleRefundsRq
    r.Set("param_refund_order_multiple_refunds_rq", paramRefundOrderMultipleRefundsRq)
    return nil
}

// ParamRefundOrderMultipleRefundsRq Getter
func (r TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest) GetParamRefundOrderMultipleRefundsRq() *RefundOrderMultipleRefundsRq {
    return r.paramRefundOrderMultipleRefundsRq
}
