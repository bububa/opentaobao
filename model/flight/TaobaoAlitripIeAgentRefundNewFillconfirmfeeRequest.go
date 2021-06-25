package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新模型-回填申请单费用 APIRequest
taobao.alitrip.ie.agent.refund.new.fillconfirmfee

1. 回填退票费用
*/
type TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest struct {
    model.Params

    // 请求
    paramRefundOrderFillConfirmFeeRq   *RefundOrderFillConfirmFeeRq 

}

func NewTaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest() *TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest{
    return &TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.new.fillconfirmfee"
}

func (r TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest) SetParamRefundOrderFillConfirmFeeRq(paramRefundOrderFillConfirmFeeRq *RefundOrderFillConfirmFeeRq) error {
    r.paramRefundOrderFillConfirmFeeRq = paramRefundOrderFillConfirmFeeRq
    r.Set("param_refund_order_fill_confirm_fee_rq", paramRefundOrderFillConfirmFeeRq)
    return nil
}

func (r TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest) GetParamRefundOrderFillConfirmFeeRq() *RefundOrderFillConfirmFeeRq {
    return r.paramRefundOrderFillConfirmFeeRq
}

