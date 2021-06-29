package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取员工租户管理员信息（查询员工是否为租户管理员） APIRequest
alibaba.mozi.vds.tenant.api.service.getadmin

获取员工租户管理员信息（查询员工是否为租户管理员）
*/
type AlibabaMoziVdsTenantApiServiceGetadminRequest struct {
    model.Params

    // 入参
    par0   *GetEmployeeTenantAdminInfoRequest 

}

func NewAlibabaMoziVdsTenantApiServiceGetadminRequest() *AlibabaMoziVdsTenantApiServiceGetadminRequest{
    return &AlibabaMoziVdsTenantApiServiceGetadminRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziVdsTenantApiServiceGetadminRequest) GetApiMethodName() string {
    return "alibaba.mozi.vds.tenant.api.service.getadmin"
}

func (r AlibabaMoziVdsTenantApiServiceGetadminRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziVdsTenantApiServiceGetadminRequest) SetPar0(par0 *GetEmployeeTenantAdminInfoRequest) error {
    r.par0 = par0
    r.Set("par0", par0)
    return nil
}

func (r AlibabaMoziVdsTenantApiServiceGetadminRequest) GetPar0() *GetEmployeeTenantAdminInfoRequest {
    return r.par0
}

