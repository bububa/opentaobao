package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新模型-回填申请单费用 API请求
taobao.alitrip.ie.agent.refund.new.fillconfirmfee

1. 回填退票费用
*/
type TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest struct {
    model.Params
    // 请求
    _paramRefundOrderFillConfirmFeeRq   *RefundOrderFillConfirmFeeRq
}

// 初始化TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest对象
func NewTaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest() *TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest{
    return &TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.new.fillconfirmfee"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRefundOrderFillConfirmFeeRq Setter
// 请求
func (r *TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest) SetParamRefundOrderFillConfirmFeeRq(_paramRefundOrderFillConfirmFeeRq *RefundOrderFillConfirmFeeRq) error {
    r._paramRefundOrderFillConfirmFeeRq = _paramRefundOrderFillConfirmFeeRq
    r.Set("param_refund_order_fill_confirm_fee_rq", _paramRefundOrderFillConfirmFeeRq)
    return nil
}

// ParamRefundOrderFillConfirmFeeRq Getter
func (r TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest) GetParamRefundOrderFillConfirmFeeRq() *RefundOrderFillConfirmFeeRq {
    return r._paramRefundOrderFillConfirmFeeRq
}
