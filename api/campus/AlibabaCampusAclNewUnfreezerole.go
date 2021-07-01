package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
解冻角色 
alibaba.campus.acl.new.unfreezerole

解冻角色
*/
func AlibabaCampusAclNewUnfreezerole(clt *core.SDKClient, req *campus.AlibabaCampusAclNewUnfreezeroleAPIRequest, session string) (*campus.AlibabaCampusAclNewUnfreezeroleAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewUnfreezeroleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
