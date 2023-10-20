package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest 查询发薪结果 API请求
// alibaba.einvoice.tax.opt.salaryresult.query
//
// 查询发薪结果
type AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest struct {
	model.Params
	// 发薪流水号
	_detailIdList []string
	// 业务方编码
	_employerCode string
}

// NewAlibabaEinvoiceTaxOptSalaryresultQueryRequest 初始化AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest对象
func NewAlibabaEinvoiceTaxOptSalaryresultQueryRequest() *AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest {
	return &AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) Reset() {
	r._detailIdList = r._detailIdList[:0]
	r._employerCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.salaryresult.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetailIdList is DetailIdList Setter
// 发薪流水号
func (r *AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) SetDetailIdList(_detailIdList []string) error {
	r._detailIdList = _detailIdList
	r.Set("detail_id_list", _detailIdList)
	return nil
}

// GetDetailIdList DetailIdList Getter
func (r AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) GetDetailIdList() []string {
	return r._detailIdList
}

// SetEmployerCode is EmployerCode Setter
// 业务方编码
func (r *AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) SetEmployerCode(_employerCode string) error {
	r._employerCode = _employerCode
	r.Set("employer_code", _employerCode)
	return nil
}

// GetEmployerCode EmployerCode Getter
func (r AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) GetEmployerCode() string {
	return r._employerCode
}

var poolAlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceTaxOptSalaryresultQueryRequest()
	},
}

// GetAlibabaEinvoiceTaxOptSalaryresultQueryRequest 从 sync.Pool 获取 AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest
func GetAlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest() *AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest {
	return poolAlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest.Get().(*AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest)
}

// ReleaseAlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest 将 AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest(v *AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest.Put(v)
}
