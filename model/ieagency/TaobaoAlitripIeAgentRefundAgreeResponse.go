package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同意退票 API返回值 
taobao.alitrip.ie.agent.refund.agree

卖家同意买家退票申请
*/
type TaobaoAlitripIeAgentRefundAgreeAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundAgreeResponse
}

// 同意退票 成功返回结果
type TaobaoAlitripIeAgentRefundAgreeResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_agree_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AgreeRefundTicketRs `json:"result,omitempty" xml:"result,omitempty"`
}
