package iotticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Iot售后服务商取消工单 APIResponse
cainiao.iot.ticket.sp.cancle

IoT售后服务商取消接单
*/
type CainiaoIotTicketSpCancleAPIResponse struct {
    model.CommonResponse
    CainiaoIotTicketSpCancleResponse
}

type CainiaoIotTicketSpCancleResponse struct {
    XMLName xml.Name `xml:"cainiao_iot_ticket_sp_cancle_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *CainiaoIotTicketSpCancleResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
