package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家退票受理申请(对外) API返回值 
taobao.alitrip.ie.agent.refund.new.receive

允许代理商通过top接口受理退票申请
*/
type TaobaoAlitripIeAgentRefundNewReceiveAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundNewReceiveResponse
}

// 商家退票受理申请(对外) 成功返回结果
type TaobaoAlitripIeAgentRefundNewReceiveResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_receive_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ReceiveRefundTicketRs `json:"result,omitempty" xml:"result,omitempty"`
}
