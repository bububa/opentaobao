package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
异步开方处方查询 APIRequest
alibaba.alihealth.asyncprescribe.prescription.search

异步开方处方查询
*/
type AlibabaAlihealthAsyncprescribePrescriptionSearchRequest struct {
    model.Params

    // 查询入参
    searchRequest   *AsyncPrescribeSearchRequest 

}

func NewAlibabaAlihealthAsyncprescribePrescriptionSearchRequest() *AlibabaAlihealthAsyncprescribePrescriptionSearchRequest{
    return &AlibabaAlihealthAsyncprescribePrescriptionSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthAsyncprescribePrescriptionSearchRequest) GetApiMethodName() string {
    return "alibaba.alihealth.asyncprescribe.prescription.search"
}

func (r AlibabaAlihealthAsyncprescribePrescriptionSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthAsyncprescribePrescriptionSearchRequest) SetSearchRequest(searchRequest *AsyncPrescribeSearchRequest) error {
    r.searchRequest = searchRequest
    r.Set("search_request", searchRequest)
    return nil
}

func (r AlibabaAlihealthAsyncprescribePrescriptionSearchRequest) GetSearchRequest() *AsyncPrescribeSearchRequest {
    return r.searchRequest
}

