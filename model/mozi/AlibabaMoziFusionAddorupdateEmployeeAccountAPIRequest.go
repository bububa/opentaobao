package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest
添加人员和账号复合接口 API请求
alibaba.mozi.fusion.addorupdate.employee.account

添加人员和账号复合接口 */
type AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest struct {
	model.Params
	// 人员账号
	_employeeAccount *AddOrUpdateTenantEmployeeAndAccountRequest
}

// NewAlibabaMoziFusionAddorupdateEmployeeAccountRequest 初始化AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest对象
func NewAlibabaMoziFusionAddorupdateEmployeeAccountRequest() *AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest {
	return &AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.addorupdate.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EmployeeAccount Setter
// 人员账号
func (r *AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest) SetEmployeeAccount(_employeeAccount *AddOrUpdateTenantEmployeeAndAccountRequest) error {
	r._employeeAccount = _employeeAccount
	r.Set("employee_account", _employeeAccount)
	return nil
}

// Get EmployeeAccount Getter
func (r AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest) GetEmployeeAccount() *AddOrUpdateTenantEmployeeAndAccountRequest {
	return r._employeeAccount
}
