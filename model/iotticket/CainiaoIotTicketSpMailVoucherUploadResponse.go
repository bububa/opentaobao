package iotticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商寄出维修件上传凭证信息 API返回值 
cainiao.iot.ticket.sp.mail.voucher.upload

IoT售后服务商寄出维修件上传凭证信息
*/
type CainiaoIotTicketSpMailVoucherUploadAPIResponse struct {
    model.CommonResponse
    CainiaoIotTicketSpMailVoucherUploadResponse
}

// 服务商寄出维修件上传凭证信息 成功返回结果
type CainiaoIotTicketSpMailVoucherUploadResponse struct {
    XMLName xml.Name `xml:"cainiao_iot_ticket_sp_mail_voucher_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *ResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
