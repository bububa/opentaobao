package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicetaxoptesignresultqueryAPIRequest 查询用户签约税优结果 API请求
// alibaba.einvoice.tax.opt.esignresult.query
//
// 查询用户是否已经签约
type AlibabaeinvoicetaxoptesignresultqueryAPIRequest struct {
	model.Params
	// 业务方编码
	_employerCode string
	// 用户在业务方平台的userid
	_identificationInBelongingEmployer string
}

// NewAlibabaeinvoicetaxoptesignresultqueryRequest 初始化AlibabaeinvoicetaxoptesignresultqueryAPIRequest对象
func NewAlibabaeinvoicetaxoptesignresultqueryRequest() *AlibabaeinvoicetaxoptesignresultqueryAPIRequest {
	return &AlibabaeinvoicetaxoptesignresultqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicetaxoptesignresultqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.esignresult.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicetaxoptesignresultqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicetaxoptesignresultqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEmployerCode is EmployerCode Setter
// 业务方编码
func (r *AlibabaeinvoicetaxoptesignresultqueryAPIRequest) SetEmployerCode(_employerCode string) error {
	r._employerCode = _employerCode
	r.Set("employer_code", _employerCode)
	return nil
}

// GetEmployerCode EmployerCode Getter
func (r AlibabaeinvoicetaxoptesignresultqueryAPIRequest) GetEmployerCode() string {
	return r._employerCode
}

// SetIdentificationInBelongingEmployer is IdentificationInBelongingEmployer Setter
// 用户在业务方平台的userid
func (r *AlibabaeinvoicetaxoptesignresultqueryAPIRequest) SetIdentificationInBelongingEmployer(_identificationInBelongingEmployer string) error {
	r._identificationInBelongingEmployer = _identificationInBelongingEmployer
	r.Set("identification_in_belonging_employer", _identificationInBelongingEmployer)
	return nil
}

// GetIdentificationInBelongingEmployer IdentificationInBelongingEmployer Getter
func (r AlibabaeinvoicetaxoptesignresultqueryAPIRequest) GetIdentificationInBelongingEmployer() string {
	return r._identificationInBelongingEmployer
}
