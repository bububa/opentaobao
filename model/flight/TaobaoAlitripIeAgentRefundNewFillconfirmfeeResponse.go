package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新模型-回填申请单费用 API返回值 
taobao.alitrip.ie.agent.refund.new.fillconfirmfee

1. 回填退票费用
*/
type TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundNewFillconfirmfeeResponse
}

// 新模型-回填申请单费用 成功返回结果
type TaobaoAlitripIeAgentRefundNewFillconfirmfeeResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_fillconfirmfee_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *RefundOrderFillConfirmFeeRs `json:"result,omitempty" xml:"result,omitempty"`
}
