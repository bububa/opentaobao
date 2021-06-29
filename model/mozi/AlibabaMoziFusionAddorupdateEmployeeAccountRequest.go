package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加人员和账号复合接口 APIRequest
alibaba.mozi.fusion.addorupdate.employee.account

添加人员和账号复合接口
*/
type AlibabaMoziFusionAddorupdateEmployeeAccountRequest struct {
    model.Params

    // 人员账号
    employeeAccount   *AddOrUpdateTenantEmployeeAndAccountRequest 

}

func NewAlibabaMoziFusionAddorupdateEmployeeAccountRequest() *AlibabaMoziFusionAddorupdateEmployeeAccountRequest{
    return &AlibabaMoziFusionAddorupdateEmployeeAccountRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziFusionAddorupdateEmployeeAccountRequest) GetApiMethodName() string {
    return "alibaba.mozi.fusion.addorupdate.employee.account"
}

func (r AlibabaMoziFusionAddorupdateEmployeeAccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziFusionAddorupdateEmployeeAccountRequest) SetEmployeeAccount(employeeAccount *AddOrUpdateTenantEmployeeAndAccountRequest) error {
    r.employeeAccount = employeeAccount
    r.Set("employee_account", employeeAccount)
    return nil
}

func (r AlibabaMoziFusionAddorupdateEmployeeAccountRequest) GetEmployeeAccount() *AddOrUpdateTenantEmployeeAndAccountRequest {
    return r.employeeAccount
}

