package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】订单回填票号/成功订单 APIResponse
taobao.jipiao.agent.order.ticket

淘宝机票代理商回填票号/成功订单
*/
type TaobaoJipiaoAgentOrderTicketAPIResponse struct {
    model.CommonResponse
    TaobaoJipiaoAgentOrderTicketResponse
}

type TaobaoJipiaoAgentOrderTicketResponse struct {
    XMLName xml.Name `xml:"jipiao_agent_order_ticket_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回操作成功失败信息
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 返回回填票号操作成功失败信息
    
    IsTicketSuccess   bool `json:"is_ticket_success,omitempty" xml:"is_ticket_success,omitempty"`

    
    // 返回接口调用完成后，整个订单是否成功
    
    IsOrderSuccess   bool `json:"is_order_success,omitempty" xml:"is_order_success,omitempty"`

    
}
