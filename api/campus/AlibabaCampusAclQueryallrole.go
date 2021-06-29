package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
查询全部角色 
alibaba.campus.acl.queryallrole

查询全部园区
*/
func AlibabaCampusAclQueryallrole(clt *core.SDKClient, req *campus.AlibabaCampusAclQueryallroleRequest, session string) (*campus.AlibabaCampusAclQueryallroleAPIResponse, error) {
    var resp campus.AlibabaCampusAclQueryallroleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
