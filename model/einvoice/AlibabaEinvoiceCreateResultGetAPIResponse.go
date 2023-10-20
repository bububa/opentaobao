package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceCreateResultGetAPIResponse ERP开票结果获取 API返回值
// alibaba.einvoice.create.result.get
//
// ERP开票结果获取
type AlibabaEinvoiceCreateResultGetAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceCreateResultGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceCreateResultGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceCreateResultGetAPIResponseModel).Reset()
}

// AlibabaEinvoiceCreateResultGetAPIResponseModel is ERP开票结果获取 成功返回结果
type AlibabaEinvoiceCreateResultGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_create_result_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 开票返回结果数据列表
	InvoiceResultList []InvoiceResult `json:"invoice_result_list,omitempty" xml:"invoice_result_list>invoice_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceCreateResultGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InvoiceResultList = m.InvoiceResultList[:0]
}

var poolAlibabaEinvoiceCreateResultGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceCreateResultGetAPIResponse)
	},
}

// GetAlibabaEinvoiceCreateResultGetAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceCreateResultGetAPIResponse
func GetAlibabaEinvoiceCreateResultGetAPIResponse() *AlibabaEinvoiceCreateResultGetAPIResponse {
	return poolAlibabaEinvoiceCreateResultGetAPIResponse.Get().(*AlibabaEinvoiceCreateResultGetAPIResponse)
}

// ReleaseAlibabaEinvoiceCreateResultGetAPIResponse 将 AlibabaEinvoiceCreateResultGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceCreateResultGetAPIResponse(v *AlibabaEinvoiceCreateResultGetAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceCreateResultGetAPIResponse.Put(v)
}
