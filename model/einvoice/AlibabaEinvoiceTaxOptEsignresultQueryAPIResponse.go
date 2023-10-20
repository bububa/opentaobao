package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse 查询用户签约税优结果 API返回值
// alibaba.einvoice.tax.opt.esignresult.query
//
// 查询用户是否已经签约
type AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceTaxOptEsignresultQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceTaxOptEsignresultQueryAPIResponseModel).Reset()
}

// AlibabaEinvoiceTaxOptEsignresultQueryAPIResponseModel is 查询用户签约税优结果 成功返回结果
type AlibabaEinvoiceTaxOptEsignresultQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_esignresult_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Results []AgreementInfoDto `json:"results,omitempty" xml:"results>agreement_info_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptEsignresultQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolAlibabaEinvoiceTaxOptEsignresultQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse)
	},
}

// GetAlibabaEinvoiceTaxOptEsignresultQueryAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse
func GetAlibabaEinvoiceTaxOptEsignresultQueryAPIResponse() *AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse {
	return poolAlibabaEinvoiceTaxOptEsignresultQueryAPIResponse.Get().(*AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse)
}

// ReleaseAlibabaEinvoiceTaxOptEsignresultQueryAPIResponse 将 AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptEsignresultQueryAPIResponse(v *AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptEsignresultQueryAPIResponse.Put(v)
}
