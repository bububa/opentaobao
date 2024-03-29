package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse 查询发薪结果 API返回值
// alibaba.einvoice.tax.opt.salaryresult.query
//
// 查询发薪结果
type AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponseModel).Reset()
}

// AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponseModel is 查询发薪结果 成功返回结果
type AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryresult_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *TaxOptimizationSalaryPayResultQueryResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse)
	},
}

// GetAlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse
func GetAlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse() *AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse {
	return poolAlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse.Get().(*AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse)
}

// ReleaseAlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse 将 AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse(v *AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse.Put(v)
}
