package mozi

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mozi"
)

/* 
分页查询租户子管理员 
alibaba.mozi.vds.tenant.api.service.pagesubadmins

分页查询租户子管理员
*/
func AlibabaMoziVdsTenantApiServicePagesubadmins(clt *core.SDKClient, req *mozi.AlibabaMoziVdsTenantApiServicePagesubadminsRequest, session string) (*mozi.AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse, error) {
    var resp mozi.AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
