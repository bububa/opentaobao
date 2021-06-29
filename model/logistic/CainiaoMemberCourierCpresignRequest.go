package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
cp清理离职用户信息 API请求
cainiao.member.courier.cpresign

CP清理内部离职的用户信息
*/
type CainiaoMemberCourierCpresignRequest struct {
    model.Params
    // 菜鸟用户id
    accountId   int64
}

// 初始化CainiaoMemberCourierCpresignRequest对象
func NewCainiaoMemberCourierCpresignRequest() *CainiaoMemberCourierCpresignRequest{
    return &CainiaoMemberCourierCpresignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoMemberCourierCpresignRequest) GetApiMethodName() string {
    return "cainiao.member.courier.cpresign"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoMemberCourierCpresignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountId Setter
// 菜鸟用户id
func (r *CainiaoMemberCourierCpresignRequest) SetAccountId(accountId int64) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r CainiaoMemberCourierCpresignRequest) GetAccountId() int64 {
    return r.accountId
}
