package crm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/crm"
)

/* 
会员等级营销-会员吸纳 
taobao.crm.grademkt.member.add

商家通过该接口吸纳线上店铺会员。
*/
func TaobaoCrmGrademktMemberAdd(clt *core.SDKClient, req *crm.TaobaoCrmGrademktMemberAddAPIRequest, session string) (*crm.TaobaoCrmGrademktMemberAddAPIResponse, error) {
    var resp crm.TaobaoCrmGrademktMemberAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
