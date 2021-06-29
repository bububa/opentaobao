package mozivds

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除租户管理员服务 APIRequest
alibaba.mozi.vds.tenant.api.service.removeadmin

删除租户管理员top服务
*/
type AlibabaMoziVdsTenantApiServiceRemoveadminRequest struct {
    model.Params

    // 请求入参
    param   *RemoveTenantAdminsRequest 

}

func NewAlibabaMoziVdsTenantApiServiceRemoveadminRequest() *AlibabaMoziVdsTenantApiServiceRemoveadminRequest{
    return &AlibabaMoziVdsTenantApiServiceRemoveadminRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziVdsTenantApiServiceRemoveadminRequest) GetApiMethodName() string {
    return "alibaba.mozi.vds.tenant.api.service.removeadmin"
}

func (r AlibabaMoziVdsTenantApiServiceRemoveadminRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziVdsTenantApiServiceRemoveadminRequest) SetParam(param *RemoveTenantAdminsRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaMoziVdsTenantApiServiceRemoveadminRequest) GetParam() *RemoveTenantAdminsRequest {
    return r.param
}

