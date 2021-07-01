package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest
查询用户的发薪账号 API请求
alibaba.einvoice.tax.opt.salaryaccount.query

查询用户的发薪账号状态 */
type AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest struct {
	model.Params
	// 入参
	_paramTaxOptimizationQueryPaySalaryAccountDTO *TaxOptimizationQueryPaySalaryAccountDto
}

// New
