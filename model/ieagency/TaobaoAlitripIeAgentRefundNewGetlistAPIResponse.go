package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新查询退票申请单列表 API返回值 
taobao.alitrip.ie.agent.refund.new.getlist

查询商家国际机票退票申请单列表
*/
type TaobaoAlitripIeAgentRefundNewGetlistAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundNewGetlistAPIResponseModel
}

// 新查询退票申请单列表 成功返回结果
type TaobaoAlitripIeAgentRefundNewGetlistAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_getlist_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *RefundOrderQueryListRs `json:"result,omitempty" xml:"result,omitempty"`
}
