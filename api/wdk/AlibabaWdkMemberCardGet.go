package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
查询会员卡信息 
alibaba.wdk.member.card.get

根据会员卡查询会员信息
*/
func AlibabaWdkMemberCardGet(clt *core.SDKClient, req *wdk.AlibabaWdkMemberCardGetRequest, session string) (*wdk.AlibabaWdkMemberCardGetResponse, error) {
    var resp wdk.AlibabaWdkMemberCardGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
