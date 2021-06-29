package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取案件信息接口v2版本 APIRequest
alibaba.legal.suit.case.get

获取案件信息
*/
type AlibabaLegalSuitCaseGetRequest struct {
    model.Params

    // 案件id
    id   int64 

}

func NewAlibabaLegalSuitCaseGetRequest() *AlibabaLegalSuitCaseGetRequest{
    return &AlibabaLegalSuitCaseGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalSuitCaseGetRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.case.get"
}

func (r AlibabaLegalSuitCaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalSuitCaseGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaLegalSuitCaseGetRequest) GetId() int64 {
    return r.id
}

