package moziacl

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/moziacl"
)

/* 
分页获取应用的权限套餐 
alibaba.mozi.acl.app.getpermisspkgs

分页查询应用下的权限套餐列表
*/
func AlibabaMoziAclAppGetpermisspkgs(clt *core.SDKClient, req *moziacl.AlibabaMoziAclAppGetpermisspkgsRequest, session string) (*moziacl.AlibabaMoziAclAppGetpermisspkgsAPIResponse, error) {
    var resp moziacl.AlibabaMoziAclAppGetpermisspkgsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
