package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票代理商接口—确认出票是否成功 APIResponse
taobao.bus.agent.bookticket.confirm

代理商通过该接口通知汽车票系统订单出票结果。
*/
type TaobaoBusAgentBookticketConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoBusAgentBookticketConfirmResponse
}

type TaobaoBusAgentBookticketConfirmResponse struct {
    XMLName xml.Name `xml:"bus_agent_bookticket_confirm_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否确认成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 错误描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
