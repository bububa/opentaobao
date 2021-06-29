package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询申请单详情(新版) APIResponse
taobao.alitrip.ie.agent.refund.new.getdetail

查询申请单详情
*/
type TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundNewGetdetailResponse
}

type TaobaoAlitripIeAgentRefundNewGetdetailResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_getdetail_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *RefundOrderQueryDetailRs `json:"result,omitempty" xml:"result,omitempty"`

    
}
