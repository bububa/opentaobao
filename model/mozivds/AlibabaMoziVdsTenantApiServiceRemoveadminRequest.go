package mozivds

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除租户管理员服务 API请求
alibaba.mozi.vds.tenant.api.service.removeadmin

删除租户管理员top服务
*/
type AlibabaMoziVdsTenantApiServiceRemoveadminRequest struct {
    model.Params
    // 请求入参
    _param   *RemoveTenantAdminsRequest
}

// 初始化AlibabaMoziVdsTenantApiServiceRemoveadminRequest对象
func NewAlibabaMoziVdsTenantApiServiceRemoveadminRequest() *AlibabaMoziVdsTenantApiServiceRemoveadminRequest{
    return &AlibabaMoziVdsTenantApiServiceRemoveadminRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceRemoveadminRequest) GetApiMethodName() string {
    return "alibaba.mozi.vds.tenant.api.service.removeadmin"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceRemoveadminRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求入参
func (r *AlibabaMoziVdsTenantApiServiceRemoveadminRequest) SetParam(_param *RemoveTenantAdminsRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaMoziVdsTenantApiServiceRemoveadminRequest) GetParam() *RemoveTenantAdminsRequest {
    return r._param
}
