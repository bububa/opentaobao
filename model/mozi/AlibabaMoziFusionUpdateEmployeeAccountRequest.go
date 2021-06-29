package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新人员和账号属性 API请求
alibaba.mozi.fusion.update.employee.account

更新人员和账号基本属性
*/
type AlibabaMoziFusionUpdateEmployeeAccountRequest struct {
    model.Params
    // 入参
    _employeeAccount   *UpdateTenantEmployeeAndAccountRequest
}

// 初始化AlibabaMoziFusionUpdateEmployeeAccountRequest对象
func NewAlibabaMoziFusionUpdateEmployeeAccountRequest() *AlibabaMoziFusionUpdateEmployeeAccountRequest{
    return &AlibabaMoziFusionUpdateEmployeeAccountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionUpdateEmployeeAccountRequest) GetApiMethodName() string {
    return "alibaba.mozi.fusion.update.employee.account"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionUpdateEmployeeAccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EmployeeAccount Setter
// 入参
func (r *AlibabaMoziFusionUpdateEmployeeAccountRequest) SetEmployeeAccount(_employeeAccount *UpdateTenantEmployeeAndAccountRequest) error {
    r._employeeAccount = _employeeAccount
    r.Set("employee_account", _employeeAccount)
    return nil
}

// EmployeeAccount Getter
func (r AlibabaMoziFusionUpdateEmployeeAccountRequest) GetEmployeeAccount() *UpdateTenantEmployeeAndAccountRequest {
    return r._employeeAccount
}
