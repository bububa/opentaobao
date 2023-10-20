package fpm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCfoIncomingInvoiceRegisterAPIResponse 发票登记接口 API返回值
// alibaba.cfo.incoming.invoice.register
//
// 发票登记接口
type AlibabaCfoIncomingInvoiceRegisterAPIResponse struct {
	model.CommonResponse
	AlibabaCfoIncomingInvoiceRegisterAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCfoIncomingInvoiceRegisterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCfoIncomingInvoiceRegisterAPIResponseModel).Reset()
}

// AlibabaCfoIncomingInvoiceRegisterAPIResponseModel is 发票登记接口 成功返回结果
type AlibabaCfoIncomingInvoiceRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cfo_incoming_invoice_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 响应消息
	ResponseMsg string `json:"response_msg,omitempty" xml:"response_msg,omitempty"`
	// 发票Id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCfoIncomingInvoiceRegisterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResponseCode = ""
	m.ResponseMsg = ""
	m.Data = 0
	m.Succeeded = false
}

var poolAlibabaCfoIncomingInvoiceRegisterAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCfoIncomingInvoiceRegisterAPIResponse)
	},
}

// GetAlibabaCfoIncomingInvoiceRegisterAPIResponse 从 sync.Pool 获取 AlibabaCfoIncomingInvoiceRegisterAPIResponse
func GetAlibabaCfoIncomingInvoiceRegisterAPIResponse() *AlibabaCfoIncomingInvoiceRegisterAPIResponse {
	return poolAlibabaCfoIncomingInvoiceRegisterAPIResponse.Get().(*AlibabaCfoIncomingInvoiceRegisterAPIResponse)
}

// ReleaseAlibabaCfoIncomingInvoiceRegisterAPIResponse 将 AlibabaCfoIncomingInvoiceRegisterAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCfoIncomingInvoiceRegisterAPIResponse(v *AlibabaCfoIncomingInvoiceRegisterAPIResponse) {
	v.Reset()
	poolAlibabaCfoIncomingInvoiceRegisterAPIResponse.Put(v)
}
