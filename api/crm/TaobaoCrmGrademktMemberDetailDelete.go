package crm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/crm"
)

/* 
会员等级营销-删除商品等级营销明细 
taobao.crm.grademkt.member.detail.delete

删除商品等级营销明细
*/
func TaobaoCrmGrademktMemberDetailDelete(clt *core.SDKClient, req *crm.TaobaoCrmGrademktMemberDetailDeleteAPIRequest, session string) (*crm.TaobaoCrmGrademktMemberDetailDeleteAPIResponse, error) {
    var resp crm.TaobaoCrmGrademktMemberDetailDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
