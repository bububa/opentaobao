package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
校验用户是否有权限 
alibaba.campus.acl.new.checkuserpermission

校验用户是否有权限
*/
func AlibabaCampusAclNewCheckuserpermission(clt *core.SDKClient, req *campus.AlibabaCampusAclNewCheckuserpermissionAPIRequest, session string) (*campus.AlibabaCampusAclNewCheckuserpermissionAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewCheckuserpermissionAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
