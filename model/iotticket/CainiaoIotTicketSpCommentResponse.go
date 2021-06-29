package iotticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商工单备注 APIResponse
cainiao.iot.ticket.sp.comment

IoT售后服务商工单备注
*/
type CainiaoIotTicketSpCommentAPIResponse struct {
    model.CommonResponse
    CainiaoIotTicketSpCommentResponse
}

type CainiaoIotTicketSpCommentResponse struct {
    XMLName xml.Name `xml:"cainiao_iot_ticket_sp_comment_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *CainiaoIotTicketSpCommentResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
