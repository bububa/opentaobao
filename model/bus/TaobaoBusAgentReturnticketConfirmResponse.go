package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家回调退票 APIResponse
taobao.bus.agent.returnticket.confirm

商家通过TOP接口调用来回传退票状态
*/
type TaobaoBusAgentReturnticketConfirmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_agent_returnticket_confirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"