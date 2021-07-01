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
type AlibabaIworkCoreHrsGetpersonAPIRequest struct {
    model.Params
    // 用户ACCOUNT_ID
    _accountId   string
    // 用户ID
    _personId   int64
    // 应用ID
    _appId   string
    // 操作人ID
    _operatorId   string
}

// 初始化AlibabaIworkCoreHrsGetpersonAPIRequest对象
func NewAlibabaIworkCoreHrsGetpersonRequest() *AlibabaIworkCoreHrsGetpersonAPIRequest{
    return &AlibabaIworkCoreHrsGetpersonAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetApiMethodName() string {
    return "alibaba.iwork.core.hrs.getperson"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountId Setter
// 用户ACCOUNT_ID
func (r *AlibabaIworkCoreHrsGetpersonAPIRequest) SetAccountId(_accountId string) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetAccountId() string {
    return r._accountId
}
// PersonId Setter
// 用户ID
func (r *AlibabaIworkCoreHrsGetpersonAPIRequest) SetPersonId(_personId int64) error {
    r._personId = _personId
    r.Set("person_id", _personId)
    return nil
}

// PersonId Getter
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetPersonId() int64 {
    return r._personId
}
// AppId Setter
// 应用ID
func (r *AlibabaIworkCoreHrsGetpersonAPIRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetAppId() string {
    return r._appId
}
// OperatorId Setter
// 操作人ID
func (r *AlibabaIworkCoreHrsGetpersonAPIRequest) SetOperatorId(_operatorId string) error {
    r._operatorId = _operatorId
    r.Set("operator_id", _operatorId)
    return nil
}

// OperatorId Getter
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetOperatorId() string {
    return r._operatorId
}
