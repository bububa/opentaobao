package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozifusionupdateemployeeaccountAPIRequest 更新人员和账号属性 API请求
// alibaba.mozi.fusion.update.employee.account
//
// 更新人员和账号基本属性
type AlibabamozifusionupdateemployeeaccountAPIRequest struct {
	model.Params
	// 入参
	_employeeAccount *UpdateTenantEmployeeAndAccountRequest
}

// NewAlibabamozifusionupdateemployeeaccountRequest 初始化AlibabamozifusionupdateemployeeaccountAPIRequest对象
func NewAlibabamozifusionupdateemployeeaccountRequest() *AlibabamozifusionupdateemployeeaccountAPIRequest {
	return &AlibabamozifusionupdateemployeeaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamozifusionupdateemployeeaccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.update.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamozifusionupdateemployeeaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamozifusionupdateemployeeaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEmployeeAccount is EmployeeAccount Setter
// 入参
func (r *AlibabamozifusionupdateemployeeaccountAPIRequest) SetEmployeeAccount(_employeeAccount *UpdateTenantEmployeeAndAccountRequest) error {
	r._employeeAccount = _employeeAccount
	r.Set("employee_account", _employeeAccount)
	return nil
}

// GetEmployeeAccount EmployeeAccount Getter
func (r AlibabamozifusionupdateemployeeaccountAPIRequest) GetEmployeeAccount() *UpdateTenantEmployeeAndAccountRequest {
	return r._employeeAccount
}
