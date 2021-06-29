package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按租户ID查询租户信息 API请求
alibaba.mozi.vds.tenant.api.service.tenantbyid

按租户ID查询租户信息
*/
type AlibabaMoziVdsTenantApiServiceTenantbyidRequest struct {
    model.Params
    // 入参
    par0   *GetTenantByIdRequest
}

// 初始化AlibabaMoziVdsTenantApiServiceTenantbyidRequest对象
func NewAlibabaMoziVdsTenantApiServiceTenantbyidRequest() *AlibabaMoziVdsTenantApiServiceTenantbyidRequest{
    return &AlibabaMoziVdsTenantApiServiceTenantbyidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceTenantbyidRequest) GetApiMethodName() string {
    return "alibaba.mozi.vds.tenant.api.service.tenantbyid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceTenantbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Par0 Setter
// 入参
func (r *AlibabaMoziVdsTenantApiServiceTenantbyidRequest) SetPar0(par0 *GetTenantByIdRequest) error {
    r.par0 = par0
    r.Set("par0", par0)
    return nil
}

// Par0 Getter
func (r AlibabaMoziVdsTenantApiServiceTenantbyidRequest) GetPar0() *GetTenantByIdRequest {
    return r.par0
}
