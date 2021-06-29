package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
新增角色 
alibaba.campus.acl.insertrole

新增角色
*/
func AlibabaCampusAclInsertrole(clt *core.SDKClient, req *campus.AlibabaCampusAclInsertroleRequest, session string) (*campus.AlibabaCampusAclInsertroleAPIResponse, error) {
    var resp campus.AlibabaCampusAclInsertroleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
