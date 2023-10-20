package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest 查询发薪结果 API请求
// alibaba.einvoice.tax.opt.salaryresult.query
//
// 查询发薪结果
type AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest struct {
	model.Params
	// 发薪流水号
	_detailIdList []string
	// 业务方编码
	_employerCode string
}

// NewAlibabaeinvoicetaxoptsalaryresultqueryRequest 初始化AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest对象
func NewAlibabaeinvoicetaxoptsalaryresultqueryRequest() *AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest {
	return &AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.salaryresult.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetailIdList is DetailIdList Setter
// 发薪流水号
func (r *AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest) SetDetailIdList(_detailIdList []string) error {
	r._detailIdList = _detailIdList
	r.Set("detail_id_list", _detailIdList)
	return nil
}

// GetDetailIdList DetailIdList Getter
func (r AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest) GetDetailIdList() []string {
	return r._detailIdList
}

// SetEmployerCode is EmployerCode Setter
// 业务方编码
func (r *AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest) SetEmployerCode(_employerCode string) error {
	r._employerCode = _employerCode
	r.Set("employer_code", _employerCode)
	return nil
}

// GetEmployerCode EmployerCode Getter
func (r AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest) GetEmployerCode() string {
	return r._employerCode
}
