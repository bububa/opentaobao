package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellModifyBackfillAPIResponse
销售改签回填 API返回值
alitrip.agent.flight.sell.modify.backfill

销售改签回填 */
type AlitripAgentFlightSellModifyBackfillAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightSellModifyBackfillAPIResponseModel
}

// AlitripAgentFlightSellModifyBackfillAPIResponseModel is 销售改签回填 成功返回结果
type AlitripAgentFlightSellModifyBackfillAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_modify_backfill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripAgentFlightSellModifyBackfillResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
