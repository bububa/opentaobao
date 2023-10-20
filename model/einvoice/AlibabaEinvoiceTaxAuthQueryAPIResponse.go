package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxAuthQueryAPIResponse 发票中台授权信息获取 API返回值
// alibaba.einvoice.tax.auth.query
//
// 发票中台授权信息获取
type AlibabaEinvoiceTaxAuthQueryAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceTaxAuthQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxAuthQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceTaxAuthQueryAPIResponseModel).Reset()
}

// AlibabaEinvoiceTaxAuthQueryAPIResponseModel is 发票中台授权信息获取 成功返回结果
type AlibabaEinvoiceTaxAuthQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_auth_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxAuthQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceTaxAuthQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceTaxAuthQueryAPIResponse)
	},
}

// GetAlibabaEinvoiceTaxAuthQueryAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceTaxAuthQueryAPIResponse
func GetAlibabaEinvoiceTaxAuthQueryAPIResponse() *AlibabaEinvoiceTaxAuthQueryAPIResponse {
	return poolAlibabaEinvoiceTaxAuthQueryAPIResponse.Get().(*AlibabaEinvoiceTaxAuthQueryAPIResponse)
}

// ReleaseAlibabaEinvoiceTaxAuthQueryAPIResponse 将 AlibabaEinvoiceTaxAuthQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceTaxAuthQueryAPIResponse(v *AlibabaEinvoiceTaxAuthQueryAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceTaxAuthQueryAPIResponse.Put(v)
}
