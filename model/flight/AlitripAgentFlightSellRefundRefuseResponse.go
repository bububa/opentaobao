package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退票拒绝 APIResponse
alitrip.agent.flight.sell.refund.refuse

销售退票拒绝
*/
type AlitripAgentFlightSellRefundRefuseAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_agent_flight_sell_refund_refuse_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异步获取历史数据接口返回结果
    
    Result   *AlitripAgentFlightSellRefundRefuseResultDto `json:"result,omitempty" xml:"