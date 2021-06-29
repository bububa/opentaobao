package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
校验用户是否有菜单权限 
alibaba.campus.acl.new.checkusermenu

校验用户是否有菜单权限
*/
func AlibabaCampusAclNewCheckusermenu(clt *core.SDKClient, req *campus.AlibabaCampusAclNewCheckusermenuRequest, session string) (*campus.AlibabaCampusAclNewCheckusermenuAPIResponse, error) {
    var resp campus.AlibabaCampusAclNewCheckusermenuAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
