package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询会员卡信息 APIRequest
alibaba.wdk.member.card.get

根据会员卡查询会员信息
*/
type AlibabaWdkMemberCardGetRequest struct {
    model.Params

    // 系统自动生成
    memberQuery   *MemberQueryRequest 

}

func NewAlibabaWdkMemberCardGetRequest() *AlibabaWdkMemberCardGetRequest{
    return &AlibabaWdkMemberCardGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMemberCardGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.member.card.get"
}

func (r AlibabaWdkMemberCardGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMemberCardGetRequest) SetMemberQuery(memberQuery *MemberQueryRequest) error {
    r.memberQuery = memberQuery
    r.Set("member_query", memberQuery)
    return nil
}

func (r AlibabaWdkMemberCardGetRequest) GetMemberQuery() *MemberQueryRequest {
    return r.memberQuery
}

