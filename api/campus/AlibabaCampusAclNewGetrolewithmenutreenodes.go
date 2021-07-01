package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
根据角色id查询权限 
alibaba.campus.acl.new.getrolewithmenutreenodes

根据角色id查询权限
*/
func AlibabaCampusAclNewGetrolewithmenutreenodes(clt *core.SDKClient, req *campus.AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest, session string) (*campus.AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
