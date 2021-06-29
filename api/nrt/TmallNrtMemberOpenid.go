package nrt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nrt"
)

/* 
根据会员手机查询openId 
tmall.nrt.member.openid

根据会员手机查询openId
*/
func TmallNrtMemberOpenid(clt *core.SDKClient, req *nrt.TmallNrtMemberOpenidRequest, session string) (*nrt.TmallNrtMemberOpenidAPIResponse, error) {
    var resp nrt.TmallNrtMemberOpenidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
