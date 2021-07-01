package c2m

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/c2m"
)

/* 
淘小铺机构上下级关系 
taobao.sebp.organization.getinviteinfo

机构人员获取机构上下级关系信息
*/
func TaobaoSebpOrganizationGetinviteinfo(clt *core.SDKClient, req *c2m.TaobaoSebpOrganizationGetinviteinfoAPIRequest, session string) (*c2m.TaobaoSebpOrganizationGetinviteinfoAPIResponse, error) {
    var resp c2m.TaobaoSebpOrganizationGetinviteinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
