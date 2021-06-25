package mei

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mei"
)

/* 
会员积分变更 
tmall.crm.member.point.change

品牌CRM项目中，会员积分变更接口。
*/
func TmallCrmMemberPointChange(clt *core.SDKClient, req *mei.TmallCrmMemberPointChangeRequest, session string) (*mei.TmallCrmMemberPointChangeResponse, error) {
    var resp mei.TmallCrmMemberPointChangeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
