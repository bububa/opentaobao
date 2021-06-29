package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
人员离职 APIRequest
alibaba.mozi.fusion.dimission.employee.account

人员离职并且回收账号
*/
type AlibabaMoziFusionDimissionEmployeeAccountRequest struct {
    model.Params

    // 入参
    dimissionEmployee   *RemoveTenantEmployeeAndAccountRequest 

}

func NewAlibabaMoziFusionDimissionEmployeeAccountRequest() *AlibabaMoziFusionDimissionEmployeeAccountRequest{
    return &AlibabaMoziFusionDimissionEmployeeAccountRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziFusionDimissionEmployeeAccountRequest) GetApiMethodName() string {
    return "alibaba.mozi.fusion.dimission.employee.account"
}

func (r AlibabaMoziFusionDimissionEmployeeAccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziFusionDimissionEmployeeAccountRequest) SetDimissionEmployee(dimissionEmployee *RemoveTenantEmployeeAndAccountRequest) error {
    r.dimissionEmployee = dimissionEmployee
    r.Set("dimission.employee", dimissionEmployee)
    return nil
}

func (r AlibabaMoziFusionDimissionEmployeeAccountRequest) GetDimissionEmployee() *RemoveTenantEmployeeAndAccountRequest {
    return r.dimissionEmployee
}

