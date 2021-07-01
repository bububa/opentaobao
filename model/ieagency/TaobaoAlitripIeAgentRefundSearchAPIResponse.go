package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家查询退票申请 API返回值 
taobao.alitrip.ie.agent.refund.search

卖家查询退票申请
*/
type TaobaoAlitripIeAgentRefundSearchAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundSearchAPIResponseModel
}

// 卖家查询退票申请 成功返回结果
type TaobaoAlitripIeAgentRefundSearchAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *QueryRefundTicketsRs `json:"result,omitempty" xml:"result,omitempty"`
}
