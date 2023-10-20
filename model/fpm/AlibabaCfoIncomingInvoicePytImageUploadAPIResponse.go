package fpm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCfoIncomingInvoicePytImageUploadAPIResponse 票易通发票影像上传 API返回值
// alibaba.cfo.incoming.invoice.pyt.image.upload
//
// 票易通发票影像上传
type AlibabaCfoIncomingInvoicePytImageUploadAPIResponse struct {
	model.CommonResponse
	AlibabaCfoIncomingInvoicePytImageUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCfoIncomingInvoicePytImageUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCfoIncomingInvoicePytImageUploadAPIResponseModel).Reset()
}

// AlibabaCfoIncomingInvoicePytImageUploadAPIResponseModel is 票易通发票影像上传 成功返回结果
type AlibabaCfoIncomingInvoicePytImageUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cfo_incoming_invoice_pyt_image_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 文件resourceId
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 响应码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 相应消息
	ResponseMsg string `json:"response_msg,omitempty" xml:"response_msg,omitempty"`
	// 是否成功
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCfoIncomingInvoicePytImageUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.ResponseCode = ""
	m.ResponseMsg = ""
	m.Succeeded = false
}

var poolAlibabaCfoIncomingInvoicePytImageUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCfoIncomingInvoicePytImageUploadAPIResponse)
	},
}

// GetAlibabaCfoIncomingInvoicePytImageUploadAPIResponse 从 sync.Pool 获取 AlibabaCfoIncomingInvoicePytImageUploadAPIResponse
func GetAlibabaCfoIncomingInvoicePytImageUploadAPIResponse() *AlibabaCfoIncomingInvoicePytImageUploadAPIResponse {
	return poolAlibabaCfoIncomingInvoicePytImageUploadAPIResponse.Get().(*AlibabaCfoIncomingInvoicePytImageUploadAPIResponse)
}

// ReleaseAlibabaCfoIncomingInvoicePytImageUploadAPIResponse 将 AlibabaCfoIncomingInvoicePytImageUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCfoIncomingInvoicePytImageUploadAPIResponse(v *AlibabaCfoIncomingInvoicePytImageUploadAPIResponse) {
	v.Reset()
	poolAlibabaCfoIncomingInvoicePytImageUploadAPIResponse.Put(v)
}
