package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
根据用户查询角色 
alibaba.campus.acl.getrolebyempid

根据用户查询角色
*/
func AlibabaCampusAclGetrolebyempid(clt *core.SDKClient, req *campus.AlibabaCampusAclGetrolebyempidRequest, session string) (*campus.AlibabaCampusAclGetrolebyempidAPIResponse, error) {
    var resp campus.AlibabaCampusAclGetrolebyempidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
