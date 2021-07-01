package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
删除管理员 
alibaba.campus.acl.new.deleteuserrole

删除管理员
*/
func AlibabaCampusAclNewDeleteuserrole(clt *core.SDKClient, req *campus.AlibabaCampusAclNewDeleteuserroleAPIRequest, session string) (*campus.AlibabaCampusAclNewDeleteuserroleAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewDeleteuserroleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
