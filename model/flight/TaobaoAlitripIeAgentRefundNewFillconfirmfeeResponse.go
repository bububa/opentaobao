package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新模型-回填申请单费用 APIResponse
taobao.alitrip.ie.agent.refund.new.fillconfirmfee

1. 回填退票费用
*/
type TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripIeAgentRefundNewFillconfirmfeeResponse `json:"taobao_alitrip_ie_agent_refund_new_fillconfirmfee_response,omitempty"`
}

type TaobaoAlitripIeAgentRefundNewFillconfirmfeeResponse struct {

    // result
    Result   *RefundOrderFillConfirmFeeRs `json:"result,omitempty"`

}
