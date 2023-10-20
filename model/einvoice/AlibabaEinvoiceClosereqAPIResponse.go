package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceClosereqAPIResponse 关闭开票失败请求（失败列表可重试） API返回值
// alibaba.einvoice.closereq
//
// 关闭失败开票请求，避免造成重复开票
type AlibabaEinvoiceClosereqAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceClosereqAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceClosereqAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceClosereqAPIResponseModel).Reset()
}

// AlibabaEinvoiceClosereqAPIResponseModel is 关闭开票失败请求（失败列表可重试） 成功返回结果
type AlibabaEinvoiceClosereqAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_closereq_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关闭是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceClosereqAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaEinvoiceClosereqAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceClosereqAPIResponse)
	},
}

// GetAlibabaEinvoiceClosereqAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceClosereqAPIResponse
func GetAlibabaEinvoiceClosereqAPIResponse() *AlibabaEinvoiceClosereqAPIResponse {
	return poolAlibabaEinvoiceClosereqAPIResponse.Get().(*AlibabaEinvoiceClosereqAPIResponse)
}

// ReleaseAlibabaEinvoiceClosereqAPIResponse 将 AlibabaEinvoiceClosereqAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceClosereqAPIResponse(v *AlibabaEinvoiceClosereqAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceClosereqAPIResponse.Put(v)
}
