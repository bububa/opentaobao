package moziacl

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/moziacl"
)

/* 
新增一个角色 
alibaba.mozi.acl.role.add

新增一个角色
*/
func AlibabaMoziAclRoleAdd(clt *core.SDKClient, req *moziacl.AlibabaMoziAclRoleAddRequest, session string) (*moziacl.AlibabaMoziAclRoleAddAPIResponse, error) {
    var resp moziacl.AlibabaMoziAclRoleAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
