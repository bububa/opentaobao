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
    TaobaoBusAgentReturnticketConfirmResponse
}

type TaobaoBusAgentReturnticketConfirmResponse struct {
    XMLName xml.Name `xml:"bus_agent_returnticket_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 错误描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 是否确认成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
