package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
取消角色和权限之间的关系 
alibaba.campus.acl.cancelpermiitemfromrole

取消角色和权限之间的关系
*/
func AlibabaCampusAclCancelpermiitemfromrole(clt *core.SDKClient, req *campus.AlibabaCampusAclCancelpermiitemfromroleAPIRequest, session string) (*campus.AlibabaCampusAclCancelpermiitemfromroleAPIResponse, error) {
    var resp campus.AlibabaCampusAclCancelpermiitemfromroleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
