package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
查询用户的菜单 
alibaba.campus.acl.getmenubyempid

查询用户的菜单
*/
func AlibabaCampusAclGetmenubyempid(clt *core.SDKClient, req *campus.AlibabaCampusAclGetmenubyempidRequest, session string) (*campus.AlibabaCampusAclGetmenubyempidAPIResponse, error) {
    var resp campus.AlibabaCampusAclGetmenubyempidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
