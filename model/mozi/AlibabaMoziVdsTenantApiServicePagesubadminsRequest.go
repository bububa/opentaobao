package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询租户子管理员 APIRequest
alibaba.mozi.vds.tenant.api.service.pagesubadmins

分页查询租户子管理员
*/
type AlibabaMoziVdsTenantApiServicePagesubadminsRequest struct {
    model.Params

    // 入参
    par0   *PageTenantSubAdminsRequest 

}

func NewAlibabaMoziVdsTenantApiServicePagesubadminsRequest() *AlibabaMoziVdsTenantApiServicePagesubadminsRequest{
    return &AlibabaMoziVdsTenantApiServicePagesubadminsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziVdsTenantApiServicePagesubadminsRequest) GetApiMethodName() string {
    return "alibaba.mozi.vds.tenant.api.service.pagesubadmins"
}

func (r AlibabaMoziVdsTenantApiServicePagesubadminsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziVdsTenantApiServicePagesubadminsRequest) SetPar0(par0 *PageTenantSubAdminsRequest) error {
    r.par0 = par0
    r.Set("par0", par0)
    return nil
}

func (r AlibabaMoziVdsTenantApiServicePagesubadminsRequest) GetPar0() *PageTenantSubAdminsRequest {
    return r.par0
}

