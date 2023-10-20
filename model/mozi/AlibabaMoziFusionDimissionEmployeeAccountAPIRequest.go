package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozifusiondimissionemployeeaccountAPIRequest 人员离职 API请求
// alibaba.mozi.fusion.dimission.employee.account
//
// 人员离职并且回收账号
type AlibabamozifusiondimissionemployeeaccountAPIRequest struct {
	model.Params
	// 入参
	_dimissionEmployee *RemoveTenantEmployeeAndAccountRequest
}

// NewAlibabamozifusiondimissionemployeeaccountRequest 初始化AlibabamozifusiondimissionemployeeaccountAPIRequest对象
func NewAlibabamozifusiondimissionemployeeaccountRequest() *AlibabamozifusiondimissionemployeeaccountAPIRequest {
	return &AlibabamozifusiondimissionemployeeaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamozifusiondimissionemployeeaccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.dimission.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamozifusiondimissionemployeeaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamozifusiondimissionemployeeaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDimissionEmployee is DimissionEmployee Setter
// 入参
func (r *AlibabamozifusiondimissionemployeeaccountAPIRequest) SetDimissionEmployee(_dimissionEmployee *RemoveTenantEmployeeAndAccountRequest) error {
	r._dimissionEmployee = _dimissionEmployee
	r.Set("dimission.employee", _dimissionEmployee)
	return nil
}

// GetDimissionEmployee DimissionEmployee Getter
func (r AlibabamozifusiondimissionemployeeaccountAPIRequest) GetDimissionEmployee() *RemoveTenantEmployeeAndAccountRequest {
	return r._dimissionEmployee
}
