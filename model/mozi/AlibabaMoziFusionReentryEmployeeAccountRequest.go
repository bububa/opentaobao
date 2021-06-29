package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
重新入职并且重新启用账号 APIRequest
alibaba.mozi.fusion.reentry.employee.account

重新入职并且重新启用账号
*/
type AlibabaMoziFusionReentryEmployeeAccountRequest struct {
    model.Params

    // 入参
    reentryEmployeeAccount   *ReEntryTenantEmployeeAndAccountRequest 

}

func NewAlibabaMoziFusionReentryEmployeeAccountRequest() *AlibabaMoziFusionReentryEmployeeAccountRequest{
    return &AlibabaMoziFusionReentryEmployeeAccountRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziFusionReentryEmployeeAccountRequest) GetApiMethodName() string {
    return "alibaba.mozi.fusion.reentry.employee.account"
}

func (r AlibabaMoziFusionReentryEmployeeAccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziFusionReentryEmployeeAccountRequest) SetReentryEmployeeAccount(reentryEmployeeAccount *ReEntryTenantEmployeeAndAccountRequest) error {
    r.reentryEmployeeAccount = reentryEmployeeAccount
    r.Set("reentry.employee.account", reentryEmployeeAccount)
    return nil
}

func (r AlibabaMoziFusionReentryEmployeeAccountRequest) GetReentryEmployeeAccount() *ReEntryTenantEmployeeAndAccountRequest {
    return r.reentryEmployeeAccount
}

