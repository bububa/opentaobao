package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新人员和账号属性 APIRequest
alibaba.mozi.fusion.update.employee.account

更新人员和账号基本属性
*/
type AlibabaMoziFusionUpdateEmployeeAccountRequest struct {
    model.Params

    // 入参
    employeeAccount   *UpdateTenantEmployeeAndAccountRequest 

}

func NewAlibabaMoziFusionUpdateEmployeeAccountRequest() *AlibabaMoziFusionUpdateEmployeeAccountRequest{
    return &AlibabaMoziFusionUpdateEmployeeAccountRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziFusionUpdateEmployeeAccountRequest) GetApiMethodName() string {
    return "alibaba.mozi.fusion.update.employee.account"
}

func (r AlibabaMoziFusionUpdateEmployeeAccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziFusionUpdateEmployeeAccountRequest) SetEmployeeAccount(employeeAccount *UpdateTenantEmployeeAndAccountRequest) error {
    r.employeeAccount = employeeAccount
    r.Set("employee_account", employeeAccount)
    return nil
}

func (r AlibabaMoziFusionUpdateEmployeeAccountRequest) GetEmployeeAccount() *UpdateTenantEmployeeAndAccountRequest {
    return r.employeeAccount
}

