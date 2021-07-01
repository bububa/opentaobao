package jym

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jym"
)

/* 
用户实名认证 
taobao.jym.member.realname.verify.withoutuid

开放用户实名认证接口使用
*/
func TaobaoJymMemberRealnameVerifyWithoutuid(clt *core.SDKClient, req *jym.TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest, session string) (*jym.TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse, error) {
    var resp jym.TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
