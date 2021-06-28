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
	RequestId     string         `json:"request_id,omitempty" xml:"bus_agent_bookticket_confirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否确认成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"