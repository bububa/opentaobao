package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
冻结角色 
alibaba.campus.acl.new.freezerole

冻结角色
*/
func AlibabaCampusAclNewFreezerole(clt *core.SDKClient, req *campus.AlibabaCampusAclNewFreezeroleRequest, session string) (*campus.AlibabaCampusAclNewFreezeroleAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewFreezeroleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
