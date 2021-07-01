package mozivds

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新建租户管理员 API请求
alibaba.mozi.vds.tenant.api.service.addadmin

新建租户管理员
alibaba.mozi.vds.tenant.api.service.addadmin
*/
type AlibabaMoziVdsTenantApiServiceAddadminAPIRequest struct {
    model.Params
    // 请求参数
    _param0   *AddTenantAdminsRequest
}

// 初始化AlibabaMoziVdsTenantApiServiceAddadminAPIRequest对象
func NewAlibabaMoziVdsTenantApiServiceAddadminRequest() *AlibabaMoziVdsTenantApiServiceAddadminAPIRequest{
    return &AlibabaMoziVdsTenantApiServiceAddadminAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceAddadminAPIRequest) GetApiMethodName() string {
    return "alibaba.mozi.vds.tenant.api.service.addadmin"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceAddadminAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 请求参数
func (r *AlibabaMoziVdsTenantApiServiceAddadminAPIRequest) SetParam0(_param0 *AddTenantAdminsRequest) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaMoziVdsTenantApiServiceAddadminAPIRequest) GetParam0() *AddTenantAdminsRequest {
    return r._param0
}
