package mozivds

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mozivds"
)

/* 
新建租户管理员 
alibaba.mozi.vds.tenant.api.service.addadmin

新建租户管理员
alibaba.mozi.vds.tenant.api.service.addadmin
*/
func AlibabaMoziVdsTenantApiServiceAddadmin(clt *core.SDKClient, req *mozivds.AlibabaMoziVdsTenantApiServiceAddadminRequest, session string) (*mozivds.AlibabaMoziVdsTenantApiServiceAddadminAPIResponse, error) {
    var resp mozivds.AlibabaMoziVdsTenantApiServiceAddadminAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
