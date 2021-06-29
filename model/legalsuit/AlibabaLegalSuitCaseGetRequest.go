package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取案件信息接口v2版本 API请求
alibaba.legal.suit.case.get

获取案件信息
*/
type AlibabaLegalSuitCaseGetRequest struct {
    model.Params
    // 案件id
    id   int64
}

// 初始化AlibabaLegalSuitCaseGetRequest对象
func NewAlibabaLegalSuitCaseGetRequest() *AlibabaLegalSuitCaseGetRequest{
    return &AlibabaLegalSuitCaseGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCaseGetRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.case.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 案件id
func (r *AlibabaLegalSuitCaseGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r AlibabaLegalSuitCaseGetRequest) GetId() int64 {
    return r.id
}
