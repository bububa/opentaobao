package traveltrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪门票核销通知 API返回值 
taobao.travel.ticket.order.verify

系统商通过TOP接口调用通知飞猪门票核销情况
*/
type TaobaoTravelTicketOrderVerifyAPIResponse struct {
    model.CommonResponse
    TaobaoTravelTicketOrderVerifyResponse
}

// 飞猪门票核销通知 成功返回结果
type TaobaoTravelTicketOrderVerifyResponse struct {
    XMLName xml.Name `xml:"travel_ticket_order_verify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 正确返回true
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
