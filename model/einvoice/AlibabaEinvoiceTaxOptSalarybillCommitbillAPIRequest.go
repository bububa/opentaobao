package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest
提交发薪账单 API请求
alibaba.einvoice.tax.opt.salarybill.commitbill

提交发薪账单 */
type AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest struct {
	model.Params
	// 入参
	_paramTaxOptimizationSalaryBillCommitReqDTO *TaxOptimizationSalaryBillCommitReqDto
}

// New
