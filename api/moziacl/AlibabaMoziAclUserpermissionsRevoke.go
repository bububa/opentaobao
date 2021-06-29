package moziacl

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/moziacl"
)

/* 
回收账户权限 
alibaba.mozi.acl.userpermissions.revoke

调用此接口，会根据入参回收该账户下的该批权限点
*/
func AlibabaMoziAclUserpermissionsRevoke(clt *core.SDKClient, req *moziacl.AlibabaMoziAclUserpermissionsRevokeRequest, session string) (*moziacl.AlibabaMoziAclUserpermissionsRevokeAPIResponse, error) {
    var resp moziacl.AlibabaMoziAclUserpermissionsRevokeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
