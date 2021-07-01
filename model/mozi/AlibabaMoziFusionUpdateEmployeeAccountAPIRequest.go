package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziFusionUpdateEmployeeAccountAPIRequest
更新人员和账号属性 API请求
alibaba.mozi.fusion.update.employee.account

更新人员和账号基本属性 */
type AlibabaMoziFusionUpdateEmployeeAccountAPIRequest struct {
	model.Params
	// 入参
	_employeeAccount *UpdateTenantEmployeeAndAccountRequest
}

// NewAlibabaMoziFusionUpdateEmployeeAccountRequest 初始化AlibabaMoziFusionUpdateEmployeeAccountAPIRequest对象
func NewAlibabaMoziFusionUpdateEmployeeAccountRequest() *AlibabaMoziFusionUpdateEmployeeAccountAPIRequest {
	return &AlibabaMoziFusionUpdateEmployeeAccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionUpdateEmployeeAccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.update.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionUpdateEmployeeAccountAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EmployeeAccount Setter
// 入参
func (r *AlibabaMoziFusionUpdateEmployeeAccountAPIRequest) SetEmployeeAccount(_employeeAccount *UpdateTenantEmployeeAndAccountRequest) error {
	r._employeeAccount = _employeeAccount
	r.Set("employee_account", _employeeAccount)
	return nil
}

// Get EmployeeAccount Getter
func (r AlibabaMoziFusionUpdateEmployeeAccountAPIRequest) GetEmployeeAccount() *UpdateTenantEmployeeAndAccountRequest {
	return r._employeeAccount
}
