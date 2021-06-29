package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
补退接口 APIRequest
taobao.alitrip.ie.agent.refund.new.multiplerefunds

1. 补退接口， 可以进行多次退款
*/
type TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest struct {
    model.Params

    // 请求参数
    paramRefundOrderMultipleRefundsRq   *RefundOrderMultipleRefundsRq 

}

func NewTaobaoAlitripIeAgentRefundNewMultiplerefundsRequest() *TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest{
    return &TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.new.multiplerefunds"
}

func (r TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest) SetParamRefundOrderMultipleRefundsRq(paramRefundOrderMultipleRefundsRq *RefundOrderMultipleRefundsRq) error {
    r.paramRefundOrderMultipleRefundsRq = paramRefundOrderMultipleRefundsRq
    r.Set("param_refund_order_multiple_refunds_rq", paramRefundOrderMultipleRefundsRq)
    return nil
}

func (r TaobaoAlitripIeAgentRefundNewMultiplerefundsRequest) GetParamRefundOrderMultipleRefundsRq() *RefundOrderMultipleRefundsRq {
    return r.paramRefundOrderMultipleRefundsRq
}

