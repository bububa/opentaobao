package crm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/crm"
)

/* 
会员等级营销-会员关系查询 
taobao.crm.grademkt.member.query

商家通过该接口查询线上店铺会员。
*/
func TaobaoCrmGrademktMemberQuery(clt *core.SDKClient, req *crm.TaobaoCrmGrademktMemberQueryRequest, session string) (*crm.TaobaoCrmGrademktMemberQueryResponse, error) {
    var resp crm.TaobaoCrmGrademktMemberQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
