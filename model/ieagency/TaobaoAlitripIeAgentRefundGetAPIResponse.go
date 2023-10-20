package ieagency

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentrefundgetAPIResponse 获取退票申请详情 API返回值
// taobao.alitrip.ie.agent.refund.get
//
// 获取退票申请详情
type TaobaoalitripieagentrefundgetAPIResponse struct {
	model.CommonResponse
	TaobaoalitripieagentrefundgetAPIResponseModel
}

// TaobaoalitripieagentrefundgetAPIResponseModel is 获取退票申请详情 成功返回结果
type TaobaoalitripieagentrefundgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *QueryRefundTicketDetailRs `json:"result,omitempty" xml:"result,omitempty"`
}
