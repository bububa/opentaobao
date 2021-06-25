package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
委托 APIRequest
alibaba.legal.case.entrust.get

获取委托案件的基本信息
*/
type AlibabaLegalCaseEntrustGetRequest struct {
    model.Params

    // 委托id
    entrustId   int64 

}

func NewAlibabaLegalCaseEntrustGetRequest() *AlibabaLegalCaseEntrustGetRequest{
    return &AlibabaLegalCaseEntrustGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalCaseEntrustGetRequest) GetApiMethodName() string {
    return "alibaba.legal.case.entrust.get"
}

func (r AlibabaLegalCaseEntrustGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalCaseEntrustGetRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

func (r AlibabaLegalCaseEntrustGetRequest) GetEntrustId() int64 {
    return r.entrustId
}

