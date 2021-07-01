package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
查询并标记用户选择的角色 
alibaba.campus.acl.new.listuserroles

查询并标记用户选择的角色
*/
func AlibabaCampusAclNewListuserroles(clt *core.SDKClient, req *campus.AlibabaCampusAclNewListuserrolesAPIRequest, session string) (*campus.AlibabaCampusAclNewListuserrolesAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewListuserrolesAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
