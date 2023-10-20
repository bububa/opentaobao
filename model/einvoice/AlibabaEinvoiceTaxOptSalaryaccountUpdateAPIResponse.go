package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicetaxoptsalaryaccountupdateAPIResponse 更新用户发薪资产 API返回值
// alibaba.einvoice.tax.opt.salaryaccount.update
//
// 更新用户的发薪账号
type AlibabaeinvoicetaxoptsalaryaccountupdateAPIResponse struct {
	model.CommonResponse
	AlibabaeinvoicetaxoptsalaryaccountupdateAPIResponseModel
}

// AlibabaeinvoicetaxoptsalaryaccountupdateAPIResponseModel is 更新用户发薪资产 成功返回结果
type AlibabaeinvoicetaxoptsalaryaccountupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryaccount_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *TaxOptimizationEmployeeAssetUpdateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
