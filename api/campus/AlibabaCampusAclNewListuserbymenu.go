package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
查询菜单下的人员 
alibaba.campus.acl.new.listuserbymenu

查询拥有菜单权限的用户
*/
func AlibabaCampusAclNewListuserbymenu(clt *core.SDKClient, req *campus.AlibabaCampusAclNewListuserbymenuAPIRequest, session string) (*campus.AlibabaCampusAclNewListuserbymenuAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewListuserbymenuAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
