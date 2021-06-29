package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-药品同步接口 APIRequest
alibaba.alihealth.outflow.drug.saveorupdate

处方外流-药品同步接口
*/
type AlibabaAlihealthOutflowDrugSaveorupdateRequest struct {
    model.Params

    // 结果集
    drugRequest   *DrugRequest 

}

func NewAlibabaAlihealthOutflowDrugSaveorupdateRequest() *AlibabaAlihealthOutflowDrugSaveorupdateRequest{
    return &AlibabaAlihealthOutflowDrugSaveorupdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowDrugSaveorupdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.drug.saveorupdate"
}

func (r AlibabaAlihealthOutflowDrugSaveorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowDrugSaveorupdateRequest) SetDrugRequest(drugRequest *DrugRequest) error {
    r.drugRequest = drugRequest
    r.Set("drug_request", drugRequest)
    return nil
}

func (r AlibabaAlihealthOutflowDrugSaveorupdateRequest) GetDrugRequest() *DrugRequest {
    return r.drugRequest
}

