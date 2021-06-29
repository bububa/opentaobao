package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取员工租户管理员信息（查询员工是否为租户管理员） API请求
alibaba.mozi.vds.tenant.api.service.getadmin

获取员工租户管理员信息（查询员工是否为租户管理员）
*/
type AlibabaMoziVdsTenantApiServiceGetadminRequest struct {
    model.Params
    // 入参
    par0   *GetEmployeeTenantAdminInfoRequest
}

// 初始化AlibabaMoziVdsTenantApiServiceGetadminRequest对象
func NewAlibabaMoziVdsTenantApiServiceGetadminRequest() *AlibabaMoziVdsTenantApiServiceGetadminRequest{
    return &AlibabaMoziVdsTenantApiServiceGetadminRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceGetadminRequest) GetApiMethodName() string {
    return "alibaba.mozi.vds.tenant.api.service.getadmin"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceGetadminRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Par0 Setter
// 入参
func (r *AlibabaMoziVdsTenantApiServiceGetadminRequest) SetPar0(par0 *GetEmployeeTenantAdminInfoRequest) error {
    r.par0 = par0
    r.Set("par0", par0)
    return nil
}

// Par0 Getter
func (r AlibabaMoziVdsTenantApiServiceGetadminRequest) GetPar0() *GetEmployeeTenantAdminInfoRequest {
    return r.par0
}
