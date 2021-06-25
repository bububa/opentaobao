package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
cp清理离职用户信息 APIRequest
cainiao.member.courier.cpresign

CP清理内部离职的用户信息
*/
type CainiaoMemberCourierCpresignRequest struct {
    model.Params

    // 菜鸟用户id
    accountId   int64 

}

func NewCainiaoMemberCourierCpresignRequest() *CainiaoMemberCourierCpresignRequest{
    return &CainiaoMemberCourierCpresignRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoMemberCourierCpresignRequest) GetApiMethodName() string {
    return "cainiao.member.courier.cpresign"
}

func (r CainiaoMemberCourierCpresignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoMemberCourierCpresignRequest) SetAccountId(accountId int64) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r CainiaoMemberCourierCpresignRequest) GetAccountId() int64 {
    return r.accountId
}

