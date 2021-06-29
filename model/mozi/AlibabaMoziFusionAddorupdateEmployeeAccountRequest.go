package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加人员和账号复合接口 API请求
alibaba.mozi.fusion.addorupdate.employee.account

添加人员和账号复合接口
*/
type AlibabaMoziFusionAddorupdateEmployeeAccountRequest struct {
    model.Params
    // 人员账号
    employeeAccount   *AddOrUpdateTenantEmployeeAndAccountRequest
}

// 初始化AlibabaMoziFusionAddorupdateEmployeeAccountRequest对象
func NewAlibabaMoziFusionAddorupdateEmployeeAccountRequest() *AlibabaMoziFusionAddorupdateEmployeeAccountRequest{
    return &AlibabaMoziFusionAddorupdateEmployeeAccountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionAddorupdateEmployeeAccountRequest) GetApiMethodName() string {
    return "alibaba.mozi.fusion.addorupdate.employee.account"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionAddorupdateEmployeeAccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EmployeeAccount Setter
// 人员账号
func (r *AlibabaMoziFusionAddorupdateEmployeeAccountRequest) SetEmployeeAccount(employeeAccount *AddOrUpdateTenantEmployeeAndAccountRequest) error {
    r.employeeAccount = employeeAccount
    r.Set("employee_account", employeeAccount)
    return nil
}

// EmployeeAccount Getter
func (r AlibabaMoziFusionAddorupdateEmployeeAccountRequest) GetEmployeeAccount() *AddOrUpdateTenantEmployeeAndAccountRequest {
    return r.employeeAccount
}
