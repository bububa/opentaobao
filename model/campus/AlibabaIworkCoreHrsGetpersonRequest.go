package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取神鲸用户基本信息 API请求
alibaba.iwork.core.hrs.getperson

神鲸用户的基本信息查询，根据PERSON_ID或者用户ACCOUNT_ID查询
*/
type AlibabaIworkCoreHrsGetpersonRequest struct {
    model.Params
    // 用户ACCOUNT_ID
    accountId   string
    // 用户ID
    personId   int64
    // 应用ID
    appId   string
    // 操作人ID
    operatorId   string
}

// 初始化AlibabaIworkCoreHrsGetpersonRequest对象
func NewAlibabaIworkCoreHrsGetpersonRequest() *AlibabaIworkCoreHrsGetpersonRequest{
    return &AlibabaIworkCoreHrsGetpersonRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIworkCoreHrsGetpersonRequest) GetApiMethodName() string {
    return "alibaba.iwork.core.hrs.getperson"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIworkCoreHrsGetpersonRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountId Setter
// 用户ACCOUNT_ID
func (r *AlibabaIworkCoreHrsGetpersonRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r AlibabaIworkCoreHrsGetpersonRequest) GetAccountId() string {
    return r.accountId
}
// PersonId Setter
// 用户ID
func (r *AlibabaIworkCoreHrsGetpersonRequest) SetPersonId(personId int64) error {
    r.personId = personId
    r.Set("person_id", personId)
    return nil
}

// PersonId Getter
func (r AlibabaIworkCoreHrsGetpersonRequest) GetPersonId() int64 {
    return r.personId
}
// AppId Setter
// 应用ID
func (r *AlibabaIworkCoreHrsGetpersonRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r AlibabaIworkCoreHrsGetpersonRequest) GetAppId() string {
    return r.appId
}
// OperatorId Setter
// 操作人ID
func (r *AlibabaIworkCoreHrsGetpersonRequest) SetOperatorId(operatorId string) error {
    r.operatorId = operatorId
    r.Set("operator_id", operatorId)
    return nil
}

// OperatorId Getter
func (r AlibabaIworkCoreHrsGetpersonRequest) GetOperatorId() string {
    return r.operatorId
}
