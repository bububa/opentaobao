package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
出票接口 APIResponse
taobao.bus.ticket.set

提供给汽车票商家出票使用
*/
type TaobaoBusTicketSetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusTicketSetResponse `json:"bus_ticket_set_response,omitempty"` 
    TaobaoBusTicketSetResponse
}

/* model for simplify = false
type TaobaoBusTicketSetResponse struct {

    // errorCode
    
    ErrorCode1   string `json:"error_code1,omitempty"`
    

    // errorMsg
    
    ErrorMsg1   string `json:"error_msg1,omitempty"`
    

    // success1
    
    Success1   bool `json:"success1,omitempty"`
    

}
*/

type TaobaoBusTicketSetResponse struct {

    // errorCode
    ErrorCode1   string `json:"error_code1,omitempty"`

    // errorMsg
    ErrorMsg1   string `json:"error_msg1,omitempty"`

    // success1
    Success1   bool `json:"success1,omitempty"`

}
