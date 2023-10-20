package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse 查询用户的发薪账号 API返回值
// alibaba.einvoice.tax.opt.salaryaccount.query
//
// 查询用户的发薪账号状态
type AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponseModel).Reset()
}

// AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponseModel is 查询用户的发薪账号 成功返回结果
type AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryaccount_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *TaxOptimizationQueryAlipayAccountResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse)
	},
}

// GetAlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse
func GetAlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse() *AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse {
	return poolAlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse.Get().(*AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse)
}

// ReleaseAlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse 将 AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse(v *AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse.Put(v)
}
