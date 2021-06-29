package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
MOZI解除组织主管服务 API请求
alibaba.mozi.vds.tenant.api.service.dismiss

解除组织主管
*/
type AlibabaMoziVdsTenantApiServiceDismissRequest struct {
    model.Params
    // 第一个入参
    par0   *DismissOrganizationSupervisorRequest
}

// 初始化AlibabaMoziVdsTenantApiServiceDismissRequest对象
func NewAlibabaMoziVdsTenantApiServiceDismissRequest() *AlibabaMoziVdsTenantApiServiceDismissRequest{
    return &AlibabaMoziVdsTenantApiServiceDismissRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceDismissRequest) GetApiMethodName() string {
    return "alibaba.mozi.vds.tenant.api.service.dismiss"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceDismissRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Par0 Setter
// 第一个入参
func (r *AlibabaMoziVdsTenantApiServiceDismissRequest) SetPar0(par0 *DismissOrganizationSupervisorRequest) error {
    r.par0 = par0
    r.Set("par0", par0)
    return nil
}

// Par0 Getter
func (r AlibabaMoziVdsTenantApiServiceDismissRequest) GetPar0() *DismissOrganizationSupervisorRequest {
    return r.par0
}
