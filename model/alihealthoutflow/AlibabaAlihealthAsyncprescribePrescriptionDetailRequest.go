package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
异步开方处方详情 APIRequest
alibaba.alihealth.asyncprescribe.prescription.detail

异步开方处方查询
*/
type AlibabaAlihealthAsyncprescribePrescriptionDetailRequest struct {
    model.Params

    // 入参
    detailRequest   *AsyncPrescribeDetailRequest 

}

func NewAlibabaAlihealthAsyncprescribePrescriptionDetailRequest() *AlibabaAlihealthAsyncprescribePrescriptionDetailRequest{
    return &AlibabaAlihealthAsyncprescribePrescriptionDetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthAsyncprescribePrescriptionDetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.asyncprescribe.prescription.detail"
}

func (r AlibabaAlihealthAsyncprescribePrescriptionDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthAsyncprescribePrescriptionDetailRequest) SetDetailRequest(detailRequest *AsyncPrescribeDetailRequest) error {
    r.detailRequest = detailRequest
    r.Set("detail_request", detailRequest)
    return nil
}

func (r AlibabaAlihealthAsyncprescribePrescriptionDetailRequest) GetDetailRequest() *AsyncPrescribeDetailRequest {
    return r.detailRequest
}

