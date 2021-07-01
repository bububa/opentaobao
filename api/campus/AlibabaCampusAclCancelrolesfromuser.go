package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
撤销用户授予的角色 
alibaba.campus.acl.cancelrolesfromuser

撤销用户授予的角色
*/
func AlibabaCampusAclCancelrolesfromuser(clt *core.SDKClient, req *campus.AlibabaCampusAclCancelrolesfromuserAPIRequest, session string) (*campus.AlibabaCampusAclCancelrolesfromuserAPIResponse, error) {
    var resp campus.AlibabaCampusAclCancelrolesfromuserAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
