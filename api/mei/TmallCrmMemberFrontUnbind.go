package mei

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mei"
)

/* 
品牌会员解绑 
tmall.crm.member.front.unbind

品牌会员解绑功能
*/
func TmallCrmMemberFrontUnbind(clt *core.SDKClient, req *mei.TmallCrmMemberFrontUnbindAPIRequest, session string) (*mei.TmallCrmMemberFrontUnbindAPIResponse, error) {
    var resp mei.TmallCrmMemberFrontUnbindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
