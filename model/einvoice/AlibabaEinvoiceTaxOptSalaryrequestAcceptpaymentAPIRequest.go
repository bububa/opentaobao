package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest
受理发薪 API请求
alibaba.einvoice.tax.opt.salaryrequest.acceptpayment

发薪受理接口 */
type AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest struct {
	model.Params
	// 请求入参
	_paramTaxOptimizationSalaryPaymentAccessReqDTO *TaxOptimizationSalaryPaymentAccessReqDto
}

// New
