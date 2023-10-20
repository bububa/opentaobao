package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozifusioncreateemployeeaccountAPIRequest 创建MOZI自建人员和账号 API请求
// alibaba.mozi.fusion.create.employee.account
//
// 创建MOZI自建人员和账号
type AlibabamozifusioncreateemployeeaccountAPIRequest struct {
	model.Params
	// 入参
	_employeeAccount *CreateTenantEmployeeAndAccountRequest
}

// NewAlibabamozifusioncreateemployeeaccountRequest 初始化AlibabamozifusioncreateemployeeaccountAPIRequest对象
func NewAlibabamozifusioncreateemployeeaccountRequest() *AlibabamozifusioncreateemployeeaccountAPIRequest {
	return &AlibabamozifusioncreateemployeeaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamozifusioncreateemployeeaccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.create.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamozifusioncreateemployeeaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamozifusioncreateemployeeaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEmployeeAccount is EmployeeAccount Setter
// 入参
func (r *AlibabamozifusioncreateemployeeaccountAPIRequest) SetEmployeeAccount(_employeeAccount *CreateTenantEmployeeAndAccountRequest) error {
	r._employeeAccount = _employeeAccount
	r.Set("employee_account", _employeeAccount)
	return nil
}

// GetEmployeeAccount EmployeeAccount Getter
func (r AlibabamozifusioncreateemployeeaccountAPIRequest) GetEmployeeAccount() *CreateTenantEmployeeAndAccountRequest {
	return r._employeeAccount
}
