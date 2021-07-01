package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest
单明细发薪受理 API请求
alibaba.einvoice.tax.opt.salaryrequest.singleaccept

单明细发薪受理 */
type AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest struct {
	model.Params
	// 入参
	_paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO *TaxOptimizationSingleDetailSalaryPaymentAccessDto
}

// New
