package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
查询全部角色 
alibaba.campus.acl.new.listroles

查询全部角色
*/
func AlibabaCampusAclNewListroles(clt *core.SDKClient, req *campus.AlibabaCampusAclNewListrolesAPIRequest, session string) (*campus.AlibabaCampusAclNewListrolesAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewListrolesAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
