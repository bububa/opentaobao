package iotticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商签收客户邮寄设备附件上传 APIResponse
cainiao.iot.ticket.sp.mail.sign.upload

IoT售后服务商签收客户邮寄设备附件上传
*/
type CainiaoIotTicketSpMailSignUploadAPIResponse struct {
    model.CommonResponse
    CainiaoIotTicketSpMailSignUploadResponse
}

type CainiaoIotTicketSpMailSignUploadResponse struct {
    XMLName xml.Name `xml:"cainiao_iot_ticket_sp_mail_sign_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *CainiaoIotTicketSpMailSignUploadResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
