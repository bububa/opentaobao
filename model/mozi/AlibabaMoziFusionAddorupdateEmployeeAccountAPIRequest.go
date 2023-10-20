package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozifusionaddorupdateemployeeaccountAPIRequest 添加人员和账号复合接口 API请求
// alibaba.mozi.fusion.addorupdate.employee.account
//
// 添加人员和账号复合接口
type AlibabamozifusionaddorupdateemployeeaccountAPIRequest struct {
	model.Params
	// 人员账号
	_employeeAccount *AddOrUpdateTenantEmployeeAndAccountRequest
}

// NewAlibabamozifusionaddorupdateemployeeaccountRequest 初始化AlibabamozifusionaddorupdateemployeeaccountAPIRequest对象
func NewAlibabamozifusionaddorupdateemployeeaccountRequest() *AlibabamozifusionaddorupdateemployeeaccountAPIRequest {
	return &AlibabamozifusionaddorupdateemployeeaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamozifusionaddorupdateemployeeaccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.addorupdate.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamozifusionaddorupdateemployeeaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamozifusionaddorupdateemployeeaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEmployeeAccount is EmployeeAccount Setter
// 人员账号
func (r *AlibabamozifusionaddorupdateemployeeaccountAPIRequest) SetEmployeeAccount(_employeeAccount *AddOrUpdateTenantEmployeeAndAccountRequest) error {
	r._employeeAccount = _employeeAccount
	r.Set("employee_account", _employeeAccount)
	return nil
}

// GetEmployeeAccount EmployeeAccount Getter
func (r AlibabamozifusionaddorupdateemployeeaccountAPIRequest) GetEmployeeAccount() *AddOrUpdateTenantEmployeeAndAccountRequest {
	return r._employeeAccount
}
