package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
出票接口 APIResponse
taobao.bus.ticket.set

提供给汽车票商家出票使用
*/
type TaobaoBusTicketSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_ticket_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode
    
    ErrorCode1   string `json:"error_code1,omitempty" xml:"