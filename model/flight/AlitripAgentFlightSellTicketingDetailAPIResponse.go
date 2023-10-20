package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellticketingdetailAPIResponse 销售出票详情 API返回值
// alitrip.agent.flight.sell.ticketing.detail
//
// 销售出票详情
type AlitripagentflightsellticketingdetailAPIResponse struct {
	model.CommonResponse
	AlitripagentflightsellticketingdetailAPIResponseModel
}

// AlitripagentflightsellticketingdetailAPIResponseModel is 销售出票详情 成功返回结果
type AlitripagentflightsellticketingdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_ticketing_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *AlitripagentflightsellticketingdetailResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
