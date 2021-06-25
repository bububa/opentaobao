package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
委托回调接口 APIRequest
alibaba.legal.case.entrust.callback

委托回调接口
*/
type AlibabaLegalCaseEntrustCallbackRequest struct {
    model.Params

    // 委托id
    entrustId   int64 

    // 案件id
    caseId   int64 

}

func NewAlibabaLegalCaseEntrustCallbackRequest() *AlibabaLegalCaseEntrustCallbackRequest{
    return &AlibabaLegalCaseEntrustCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalCaseEntrustCallbackRequest) GetApiMethodName() string {
    return "alibaba.legal.case.entrust.callback"
}

func (r AlibabaLegalCaseEntrustCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalCaseEntrustCallbackRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

func (r AlibabaLegalCaseEntrustCallbackRequest) GetEntrustId() int64 {
    return r.entrustId
}

func (r *AlibabaLegalCaseEntrustCallbackRequest) SetCaseId(caseId int64) error {
    r.caseId = caseId
    r.Set("case_id", caseId)
    return nil
}

func (r AlibabaLegalCaseEntrustCallbackRequest) GetCaseId() int64 {
    return r.caseId
}

