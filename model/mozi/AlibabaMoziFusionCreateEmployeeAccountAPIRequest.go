package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziFusionCreateEmployeeAccountAPIRequest 创建MOZI自建人员和账号 API请求
// alibaba.mozi.fusion.create.employee.account
//
// 创建MOZI自建人员和账号
type AlibabaMoziFusionCreateEmployeeAccountAPIRequest struct {
	model.Params
	// 入参
	_employeeAccount *CreateTenantEmployeeAndAccountRequest
}

// NewAlibabaMoziFusionCreateEmployeeAccountRequest 初始化AlibabaMoziFusionCreateEmployeeAccountAPIRequest对象
func NewAlibabaMoziFusionCreateEmployeeAccountRequest() *AlibabaMoziFusionCreateEmployeeAccountAPIRequest {
	return &AlibabaMoziFusionCreateEmployeeAccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionCreateEmployeeAccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.create.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionCreateEmployeeAccountAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEmployeeAccount is EmployeeAccount Setter
// 入参
func (r *AlibabaMoziFusionCreateEmployeeAccountAPIRequest) SetEmployeeAccount(_employeeAccount *CreateTenantEmployeeAndAccountRequest) error {
	r._employeeAccount = _employeeAccount
	r.Set("employee_account", _employeeAccount)
	return nil
}

// GetEmployeeAccount EmployeeAccount Getter
func (r AlibabaMoziFusionCreateEmployeeAccountAPIRequest) GetEmployeeAccount() *CreateTenantEmployeeAndAccountRequest {
	return r._employeeAccount
}
