package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询会员卡信息 API请求
alibaba.wdk.member.card.get

根据会员卡查询会员信息
*/
type AlibabaWdkMemberCardGetRequest struct {
    model.Params
    // 系统自动生成
    _memberQuery   *MemberQueryRequest
}

// 初始化AlibabaWdkMemberCardGetRequest对象
func NewAlibabaWdkMemberCardGetRequest() *AlibabaWdkMemberCardGetRequest{
    return &AlibabaWdkMemberCardGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMemberCardGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.member.card.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMemberCardGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberQuery Setter
// 系统自动生成
func (r *AlibabaWdkMemberCardGetRequest) SetMemberQuery(_memberQuery *MemberQueryRequest) error {
    r._memberQuery = _memberQuery
    r.Set("member_query", _memberQuery)
    return nil
}

// MemberQuery Getter
func (r AlibabaWdkMemberCardGetRequest) GetMemberQuery() *MemberQueryRequest {
    return r._memberQuery
}
