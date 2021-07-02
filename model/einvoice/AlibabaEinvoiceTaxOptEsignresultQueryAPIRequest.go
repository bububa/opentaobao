package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest 查询用户签约税优结果 API请求
// alibaba.einvoice.tax.opt.esignresult.query
//
// 查询用户是否已经签约
type AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest struct {
	model.Params
	// 业务方编码
	_employerCode string
	// 用户在业务方平台的userid
	_identificationInBelongingEmployer string
}

// NewAlibabaEinvoiceTaxOptEsignresultQueryRequest 初始化AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest对象
func NewAlibabaEinvoiceTaxOptEsignresultQueryRequest() *AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest {
	return &AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.esignresult.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEmployerCode is EmployerCode Setter
// 业务方编码
func (r *AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest) SetEmployerCode(_employerCode string) error {
	r._employerCode = _employerCode
	r.Set("employer_code", _employerCode)
	return nil
}

// GetEmployerCode EmployerCode Getter
func (r AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest) GetEmployerCode() string {
	return r._employerCode
}

// SetIdentificationInBelongingEmployer is IdentificationInBelongingEmployer Setter
// 用户在业务方平台的userid
func (r *AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest) SetIdentificationInBelongingEmployer(_identificationInBelongingEmployer string) error {
	r._identificationInBelongingEmployer = _identificationInBelongingEmployer
	r.Set("identification_in_belonging_employer", _identificationInBelongingEmployer)
	return nil
}

// GetIdentificationInBelongingEmployer IdentificationInBelongingEmployer Getter
func (r AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest) GetIdentificationInBelongingEmployer() string {
	return r._identificationInBelongingEmployer
}
