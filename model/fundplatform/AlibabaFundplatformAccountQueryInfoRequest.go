package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账户信息 API请求
alibaba.fundplatform.account.query.info

外部查询资金平台用户账户信息
*/
type AlibabaFundplatformAccountQueryInfoRequest struct {
    model.Params
    // 账户ID
    accountId   int64
}

// 初始化AlibabaFundplatformAccountQueryInfoRequest对象
func NewAlibabaFundplatformAccountQueryInfoRequest() *AlibabaFundplatformAccountQueryInfoRequest{
    return &AlibabaFundplatformAccountQueryInfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformAccountQueryInfoRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.account.query.info"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformAccountQueryInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountId Setter
// 账户ID
func (r *AlibabaFundplatformAccountQueryInfoRequest) SetAccountId(accountId int64) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r AlibabaFundplatformAccountQueryInfoRequest) GetAccountId() int64 {
    return r.accountId
}
