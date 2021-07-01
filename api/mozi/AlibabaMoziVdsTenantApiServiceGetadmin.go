package mozi

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mozi"
)

/* 
获取员工租户管理员信息（查询员工是否为租户管理员） 
alibaba.mozi.vds.tenant.api.service.getadmin

获取员工租户管理员信息（查询员工是否为租户管理员）
*/
func AlibabaMoziVdsTenantApiServiceGetadmin(clt *core.SDKClient, req *mozi.AlibabaMoziVdsTenantApiServiceGetadminAPIRequest, session string) (*mozi.AlibabaMoziVdsTenantApiServiceGetadminAPIResponse, error) {
    var resp mozi.AlibabaMoziVdsTenantApiServiceGetadminAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
