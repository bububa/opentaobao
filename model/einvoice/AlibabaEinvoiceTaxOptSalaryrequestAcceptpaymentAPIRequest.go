package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest 受理发薪 API请求
// alibaba.einvoice.tax.opt.salaryrequest.acceptpayment
//
// 发薪受理接口
type AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest struct {
	model.Params
	// 请求入参
	_paramTaxOptimizationSalaryPaymentAccessReqDTO *TaxOptimizationSalaryPaymentAccessReqDto
}

// NewAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest 初始化AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest对象
func NewAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest() *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest {
	return &AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest) Reset() {
	r._paramTaxOptimizationSalaryPaymentAccessReqDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.salaryrequest.acceptpayment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTaxOptimizationSalaryPaymentAccessReqDTO is ParamTaxOptimizationSalaryPaymentAccessReqDTO Setter
// 请求入参
func (r *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest) SetParamTaxOptimizationSalaryPaymentAccessReqDTO(_paramTaxOptimizationSalaryPaymentAccessReqDTO *TaxOptimizationSalaryPaymentAccessReqDto) error {
	r._paramTaxOptimizationSalaryPaymentAccessReqDTO = _paramTaxOptimizationSalaryPaymentAccessReqDTO
	r.Set("param_tax_optimization_salary_payment_access_req_d_t_o", _paramTaxOptimizationSalaryPaymentAccessReqDTO)
	return nil
}

// GetParamTaxOptimizationSalaryPaymentAccessReqDTO ParamTaxOptimizationSalaryPaymentAccessReqDTO Getter
func (r AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest) GetParamTaxOptimizationSalaryPaymentAccessReqDTO() *TaxOptimizationSalaryPaymentAccessReqDto {
	return r._paramTaxOptimizationSalaryPaymentAccessReqDTO
}

var poolAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest()
	},
}

// GetAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest 从 sync.Pool 获取 AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest
func GetAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest() *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest {
	return poolAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest.Get().(*AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest)
}

// ReleaseAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest 将 AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest(v *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest.Put(v)
}
