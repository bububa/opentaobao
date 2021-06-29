package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
MOZI解除组织主管服务 APIRequest
alibaba.mozi.vds.tenant.api.service.dismiss

解除组织主管
*/
type AlibabaMoziVdsTenantApiServiceDismissRequest struct {
    model.Params

    // 第一个入参
    par0   *DismissOrganizationSupervisorRequest 

}

func NewAlibabaMoziVdsTenantApiServiceDismissRequest() *AlibabaMoziVdsTenantApiServiceDismissRequest{
    return &AlibabaMoziVdsTenantApiServiceDismissRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziVdsTenantApiServiceDismissRequest) GetApiMethodName() string {
    return "alibaba.mozi.vds.tenant.api.service.dismiss"
}

func (r AlibabaMoziVdsTenantApiServiceDismissRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziVdsTenantApiServiceDismissRequest) SetPar0(par0 *DismissOrganizationSupervisorRequest) error {
    r.par0 = par0
    r.Set("par0", par0)
    return nil
}

func (r AlibabaMoziVdsTenantApiServiceDismissRequest) GetPar0() *DismissOrganizationSupervisorRequest {
    return r.par0
}

