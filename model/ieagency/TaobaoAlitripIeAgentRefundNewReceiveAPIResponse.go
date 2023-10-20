package ieagency

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentrefundnewreceiveAPIResponse 商家退票受理申请(对外) API返回值
// taobao.alitrip.ie.agent.refund.new.receive
//
// 允许代理商通过top接口受理退票申请
type TaobaoalitripieagentrefundnewreceiveAPIResponse struct {
	model.CommonResponse
	TaobaoalitripieagentrefundnewreceiveAPIResponseModel
}

// TaobaoalitripieagentrefundnewreceiveAPIResponseModel is 商家退票受理申请(对外) 成功返回结果
type TaobaoalitripieagentrefundnewreceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ReceiveRefundTicketRs `json:"result,omitempty" xml:"result,omitempty"`
}
