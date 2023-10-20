package iotticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpMailVoucherUploadAPIResponse 服务商寄出维修件上传凭证信息 API返回值
// cainiao.iot.ticket.sp.mail.voucher.upload
//
// IoT售后服务商寄出维修件上传凭证信息
type CainiaoIotTicketSpMailVoucherUploadAPIResponse struct {
	model.CommonResponse
	CainiaoIotTicketSpMailVoucherUploadAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoIotTicketSpMailVoucherUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoIotTicketSpMailVoucherUploadAPIResponseModel).Reset()
}

// CainiaoIotTicketSpMailVoucherUploadAPIResponseModel is 服务商寄出维修件上传凭证信息 成功返回结果
type CainiaoIotTicketSpMailVoucherUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_sp_mail_voucher_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoIotTicketSpMailVoucherUploadResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoIotTicketSpMailVoucherUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoIotTicketSpMailVoucherUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketSpMailVoucherUploadAPIResponse)
	},
}

// GetCainiaoIotTicketSpMailVoucherUploadAPIResponse 从 sync.Pool 获取 CainiaoIotTicketSpMailVoucherUploadAPIResponse
func GetCainiaoIotTicketSpMailVoucherUploadAPIResponse() *CainiaoIotTicketSpMailVoucherUploadAPIResponse {
	return poolCainiaoIotTicketSpMailVoucherUploadAPIResponse.Get().(*CainiaoIotTicketSpMailVoucherUploadAPIResponse)
}

// ReleaseCainiaoIotTicketSpMailVoucherUploadAPIResponse 将 CainiaoIotTicketSpMailVoucherUploadAPIResponse 保存到 sync.Pool
func ReleaseCainiaoIotTicketSpMailVoucherUploadAPIResponse(v *CainiaoIotTicketSpMailVoucherUploadAPIResponse) {
	v.Reset()
	poolCainiaoIotTicketSpMailVoucherUploadAPIResponse.Put(v)
}
