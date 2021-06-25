package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
喵街会员是否绑定 
alibaba.mj.member.hasbind

喵街检测用户是否为数字化会员
*/
func AlibabaMjMemberHasbind(clt *core.SDKClient, req *mos.AlibabaMjMemberHasbindRequest, session string) (*mos.AlibabaMjMemberHasbindResponse, error) {
    var resp mos.AlibabaMjMemberHasbindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
