package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
保存角色级联保存角色和权限的关系 
alibaba.campus.acl.new.saverolewithmenu

保存角色级联保存角色和权限的关系
*/
func AlibabaCampusAclNewSaverolewithmenu(clt *core.SDKClient, req *campus.AlibabaCampusAclNewSaverolewithmenuAPIRequest, session string) (*campus.AlibabaCampusAclNewSaverolewithmenuAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewSaverolewithmenuAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
