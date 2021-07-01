package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest
更新用户发薪资产 API请求
alibaba.einvoice.tax.opt.salaryaccount.update

更新用户的发薪账号 */
type AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramTaxOptimizationEmployeeAssetUpdateDTO *TaxOptimizationEmployeeAssetUpdateDto
}

// New
