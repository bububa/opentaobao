package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
校验用户是否有角色 
alibaba.campus.acl.new.checkuserrole

校验用户是否有角色
*/
func AlibabaCampusAclNewCheckuserrole(clt *core.SDKClient, req *campus.AlibabaCampusAclNewCheckuserroleRequest, session string) (*campus.AlibabaCampusAclNewCheckuserroleAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewCheckuserroleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
