package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新模型-回填申请单费用 APIResponse
taobao.alitrip.ie.agent.refund.new.fillconfirmfee

1. 回填退票费用
*/
type TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_ie_agent_refund_new_fillconfirmfee_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *RefundOrderFillConfirmFeeRs `json:"result,omitempty" xml:"