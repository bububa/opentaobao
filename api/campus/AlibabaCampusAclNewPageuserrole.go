package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
分页查询管理员 
alibaba.campus.acl.new.pageuserrole

新增用户和角色的关系
*/
func AlibabaCampusAclNewPageuserrole(clt *core.SDKClient, req *campus.AlibabaCampusAclNewPageuserroleRequest, session string) (*campus.AlibabaCampusAclNewPageuserroleAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewPageuserroleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
