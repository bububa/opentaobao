package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验组-员工是否匹配 API请求
alibaba.mozi.vds.tenant.api.service.matchempcodes

校验组-员工是否匹配
*/
type AlibabaMoziVdsTenantApiServiceMatchempcodesRequest struct {
    model.Params
    // 入参
    par0   *MatchWithEmployeeRequest
}

// 初始化AlibabaMoziVdsTenantApiServiceMatchempcodesRequest对象
func NewAlibabaMoziVdsTenantApiServiceMatchempcodesRequest() *AlibabaMoziVdsTenantApiServiceMatchempcodesRequest{
    return &AlibabaMoziVdsTenantApiServiceMatchempcodesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceMatchempcodesRequest) GetApiMethodName() string {
    return "alibaba.mozi.vds.tenant.api.service.matchempcodes"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceMatchempcodesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Par0 Setter
// 入参
func (r *AlibabaMoziVdsTenantApiServiceMatchempcodesRequest) SetPar0(par0 *MatchWithEmployeeRequest) error {
    r.par0 = par0
    r.Set("par0", par0)
    return nil
}

// Par0 Getter
func (r AlibabaMoziVdsTenantApiServiceMatchempcodesRequest) GetPar0() *MatchWithEmployeeRequest {
    return r.par0
}
