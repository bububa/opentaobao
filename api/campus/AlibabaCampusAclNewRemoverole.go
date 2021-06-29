package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
删除角色 
alibaba.campus.acl.new.removerole

删除角色
*/
func AlibabaCampusAclNewRemoverole(clt *core.SDKClient, req *campus.AlibabaCampusAclNewRemoveroleRequest, session string) (*campus.AlibabaCampusAclNewRemoveroleAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewRemoveroleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
