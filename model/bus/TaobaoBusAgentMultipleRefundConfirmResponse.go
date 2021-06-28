package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
综合交通多次退款接口 APIResponse
taobao.bus.agent.multiple.refund.confirm

此接口支持多次按照单客进行多次退款操作，只进行退款操作。
*/
type TaobaoBusAgentMultipleRefundConfirmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_agent_multiple_refund_confirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"