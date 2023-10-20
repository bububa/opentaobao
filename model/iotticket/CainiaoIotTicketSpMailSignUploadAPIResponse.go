package iotticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpMailSignUploadAPIResponse IoT售后服务商签收客户邮寄设备附件上传 API返回值
// cainiao.iot.ticket.sp.mail.sign.upload
//
// IoT售后服务商签收客户邮寄设备附件上传
type CainiaoIotTicketSpMailSignUploadAPIResponse struct {
	model.CommonResponse
	CainiaoIotTicketSpMailSignUploadAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoIotTicketSpMailSignUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoIotTicketSpMailSignUploadAPIResponseModel).Reset()
}

// CainiaoIotTicketSpMailSignUploadAPIResponseModel is IoT售后服务商签收客户邮寄设备附件上传 成功返回结果
type CainiaoIotTicketSpMailSignUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_sp_mail_sign_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoIotTicketSpMailSignUploadResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoIotTicketSpMailSignUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoIotTicketSpMailSignUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketSpMailSignUploadAPIResponse)
	},
}

// GetCainiaoIotTicketSpMailSignUploadAPIResponse 从 sync.Pool 获取 CainiaoIotTicketSpMailSignUploadAPIResponse
func GetCainiaoIotTicketSpMailSignUploadAPIResponse() *CainiaoIotTicketSpMailSignUploadAPIResponse {
	return poolCainiaoIotTicketSpMailSignUploadAPIResponse.Get().(*CainiaoIotTicketSpMailSignUploadAPIResponse)
}

// ReleaseCainiaoIotTicketSpMailSignUploadAPIResponse 将 CainiaoIotTicketSpMailSignUploadAPIResponse 保存到 sync.Pool
func ReleaseCainiaoIotTicketSpMailSignUploadAPIResponse(v *CainiaoIotTicketSpMailSignUploadAPIResponse) {
	v.Reset()
	poolCainiaoIotTicketSpMailSignUploadAPIResponse.Put(v)
}
