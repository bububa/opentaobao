package iotticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoiotticketspmailsignuploadAPIResponse IoT售后服务商签收客户邮寄设备附件上传 API返回值
// cainiao.iot.ticket.sp.mail.sign.upload
//
// IoT售后服务商签收客户邮寄设备附件上传
type CainiaoiotticketspmailsignuploadAPIResponse struct {
	model.CommonResponse
	CainiaoiotticketspmailsignuploadAPIResponseModel
}

// CainiaoiotticketspmailsignuploadAPIResponseModel is IoT售后服务商签收客户邮寄设备附件上传 成功返回结果
type CainiaoiotticketspmailsignuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_sp_mail_sign_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoiotticketspmailsignuploadResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
