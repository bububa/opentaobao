package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
修改用户到角色关系 
alibaba.campus.acl.updategrantroletouser

修改用户到角色关系
*/
func AlibabaCampusAclUpdategrantroletouser(clt *core.SDKClient, req *campus.AlibabaCampusAclUpdategrantroletouserAPIRequest, session string) (*campus.AlibabaCampusAclUpdategrantroletouserAPIResponse, error) {
    var resp campus.AlibabaCampusAclUpdategrantroletouserAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
