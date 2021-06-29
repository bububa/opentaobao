package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建MOZI自建人员和账号 APIRequest
alibaba.mozi.fusion.create.employee.account

创建MOZI自建人员和账号
*/
type AlibabaMoziFusionCreateEmployeeAccountRequest struct {
    model.Params

    // 入参
    employeeAccount   *CreateTenantEmployeeAndAccountRequest 

}

func NewAlibabaMoziFusionCreateEmployeeAccountRequest() *AlibabaMoziFusionCreateEmployeeAccountRequest{
    return &AlibabaMoziFusionCreateEmployeeAccountRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziFusionCreateEmployeeAccountRequest) GetApiMethodName() string {
    return "alibaba.mozi.fusion.create.employee.account"
}

func (r AlibabaMoziFusionCreateEmployeeAccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziFusionCreateEmployeeAccountRequest) SetEmployeeAccount(employeeAccount *CreateTenantEmployeeAndAccountRequest) error {
    r.employeeAccount = employeeAccount
    r.Set("employee_account", employeeAccount)
    return nil
}

func (r AlibabaMoziFusionCreateEmployeeAccountRequest) GetEmployeeAccount() *CreateTenantEmployeeAndAccountRequest {
    return r.employeeAccount
}

