package ieagency

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentrefundrefuseAPIResponse 拒绝退票申请 API返回值
// taobao.alitrip.ie.agent.refund.refuse
//
// 卖家拒绝退票退票申请
type TaobaoalitripieagentrefundrefuseAPIResponse struct {
	model.CommonResponse
	TaobaoalitripieagentrefundrefuseAPIResponseModel
}

// TaobaoalitripieagentrefundrefuseAPIResponseModel is 拒绝退票申请 成功返回结果
type TaobaoalitripieagentrefundrefuseAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_refuse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RefuseRefundTicketRs `json:"result,omitempty" xml:"result,omitempty"`
}
