package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-药品同步接口 API请求
alibaba.alihealth.outflow.drug.saveorupdate

处方外流-药品同步接口
*/
type AlibabaAlihealthOutflowDrugSaveorupdateRequest struct {
    model.Params
    // 结果集
    drugRequest   *DrugRequest
}

// 初始化AlibabaAlihealthOutflowDrugSaveorupdateRequest对象
func NewAlibabaAlihealthOutflowDrugSaveorupdateRequest() *AlibabaAlihealthOutflowDrugSaveorupdateRequest{
    return &AlibabaAlihealthOutflowDrugSaveorupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowDrugSaveorupdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.drug.saveorupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowDrugSaveorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DrugRequest Setter
// 结果集
func (r *AlibabaAlihealthOutflowDrugSaveorupdateRequest) SetDrugRequest(drugRequest *DrugRequest) error {
    r.drugRequest = drugRequest
    r.Set("drug_request", drugRequest)
    return nil
}

// DrugRequest Getter
func (r AlibabaAlihealthOutflowDrugSaveorupdateRequest) GetDrugRequest() *DrugRequest {
    return r.drugRequest
}
