package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellmodifydetailAPIResponse 销售改签详情 API返回值
// alitrip.agent.flight.sell.modify.detail
//
// 销售改签详情
type AlitripagentflightsellmodifydetailAPIResponse struct {
	model.CommonResponse
	AlitripagentflightsellmodifydetailAPIResponseModel
}

// AlitripagentflightsellmodifydetailAPIResponseModel is 销售改签详情 成功返回结果
type AlitripagentflightsellmodifydetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_modify_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripagentflightsellmodifydetailResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
