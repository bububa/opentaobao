package einvoice

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest) Reset() {
	r._employerCode = ""
	r._identificationInBelongingEmployer = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.esignresult.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaEinvoiceTaxOptEsignresultQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceTaxOptEsignresultQueryRequest()
	},
}

// GetAlibabaEinvoiceTaxOptEsignresultQueryRequest 从 sync.Pool 获取 AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest
func GetAlibabaEinvoiceTaxOptEsignresultQueryAPIRequest() *AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest {
	return poolAlibabaEinvoiceTaxOptEsignresultQueryAPIRequest.Get().(*AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest)
}

// ReleaseAlibabaEinvoiceTaxOptEsignresultQueryAPIRequest 将 AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptEsignresultQueryAPIRequest(v *AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptEsignresultQueryAPIRequest.Put(v)
}
