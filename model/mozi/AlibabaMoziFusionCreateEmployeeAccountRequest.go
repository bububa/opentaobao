package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建MOZI自建人员和账号 API请求
alibaba.mozi.fusion.create.employee.account

创建MOZI自建人员和账号
*/
type AlibabaMoziFusionCreateEmployeeAccountRequest struct {
    model.Params
    // 入参
    _employeeAccount   *CreateTenantEmployeeAndAccountRequest
}

// 初始化AlibabaMoziFusionCreateEmployeeAccountRequest对象
func NewAlibabaMoziFusionCreateEmployeeAccountRequest() *AlibabaMoziFusionCreateEmployeeAccountRequest{
    return &AlibabaMoziFusionCreateEmployeeAccountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionCreateEmployeeAccountRequest) GetApiMethodName() string {
    return "alibaba.mozi.fusion.create.employee.account"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionCreateEmployeeAccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EmployeeAccount Setter
// 入参
func (r *AlibabaMoziFusionCreateEmployeeAccountRequest) SetEmployeeAccount(_employeeAccount *CreateTenantEmployeeAndAccountRequest) error {
    r._employeeAccount = _employeeAccount
    r.Set("employee_account", _employeeAccount)
    return nil
}

// EmployeeAccount Getter
func (r AlibabaMoziFusionCreateEmployeeAccountRequest) GetEmployeeAccount() *CreateTenantEmployeeAndAccountRequest {
    return r._employeeAccount
}
