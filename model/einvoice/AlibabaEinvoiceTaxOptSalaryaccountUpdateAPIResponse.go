package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse 更新用户发薪资产 API返回值
// alibaba.einvoice.tax.opt.salaryaccount.update
//
// 更新用户的发薪账号
type AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponseModel).Reset()
}

// AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponseModel is 更新用户发薪资产 成功返回结果
type AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryaccount_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *TaxOptimizationEmployeeAssetUpdateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse)
	},
}

// GetAlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse
func GetAlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse() *AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse {
	return poolAlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse.Get().(*AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse)
}

// ReleaseAlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse 将 AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse(v *AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse.Put(v)
}
