package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceSerialnoGenerateAPIResponse 获取统一开票流水号 API返回值
// alibaba.einvoice.serialno.generate
//
// erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突
type AlibabaEinvoiceSerialnoGenerateAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceSerialnoGenerateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceSerialnoGenerateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceSerialnoGenerateAPIResponseModel).Reset()
}

// AlibabaEinvoiceSerialnoGenerateAPIResponseModel is 获取统一开票流水号 成功返回结果
type AlibabaEinvoiceSerialnoGenerateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_serialno_generate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	SerialNo string `json:"serial_no,omitempty" xml:"serial_no,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceSerialnoGenerateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SerialNo = ""
}

var poolAlibabaEinvoiceSerialnoGenerateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceSerialnoGenerateAPIResponse)
	},
}

// GetAlibabaEinvoiceSerialnoGenerateAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceSerialnoGenerateAPIResponse
func GetAlibabaEinvoiceSerialnoGenerateAPIResponse() *AlibabaEinvoiceSerialnoGenerateAPIResponse {
	return poolAlibabaEinvoiceSerialnoGenerateAPIResponse.Get().(*AlibabaEinvoiceSerialnoGenerateAPIResponse)
}

// ReleaseAlibabaEinvoiceSerialnoGenerateAPIResponse 将 AlibabaEinvoiceSerialnoGenerateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceSerialnoGenerateAPIResponse(v *AlibabaEinvoiceSerialnoGenerateAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceSerialnoGenerateAPIResponse.Put(v)
}
