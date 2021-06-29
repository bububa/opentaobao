package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按租户ID查询租户信息 APIRequest
alibaba.mozi.vds.tenant.api.service.tenantbyid

按租户ID查询租户信息
*/
type AlibabaMoziVdsTenantApiServiceTenantbyidRequest struct {
    model.Params

    // 入参
    par0   *GetTenantByIdRequest 

}

func NewAlibabaMoziVdsTenantApiServiceTenantbyidRequest() *AlibabaMoziVdsTenantApiServiceTenantbyidRequest{
    return &AlibabaMoziVdsTenantApiServiceTenantbyidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziVdsTenantApiServiceTenantbyidRequest) GetApiMethodName() string {
    return "alibaba.mozi.vds.tenant.api.service.tenantbyid"
}

func (r AlibabaMoziVdsTenantApiServiceTenantbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziVdsTenantApiServiceTenantbyidRequest) SetPar0(par0 *GetTenantByIdRequest) error {
    r.par0 = par0
    r.Set("par0", par0)
    return nil
}

func (r AlibabaMoziVdsTenantApiServiceTenantbyidRequest) GetPar0() *GetTenantByIdRequest {
    return r.par0
}

