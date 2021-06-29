package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-用法字典表 APIRequest
alibaba.alihealth.outflow.usage.saveorupdate

阿里健康-处方外流-对外提供用法字典表维护功能
*/
type AlibabaAlihealthOutflowUsageSaveorupdateRequest struct {
    model.Params

    // 用法数据
    usageRequest   *UsageRequest 

}

func NewAlibabaAlihealthOutflowUsageSaveorupdateRequest() *AlibabaAlihealthOutflowUsageSaveorupdateRequest{
    return &AlibabaAlihealthOutflowUsageSaveorupdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowUsageSaveorupdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.usage.saveorupdate"
}

func (r AlibabaAlihealthOutflowUsageSaveorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowUsageSaveorupdateRequest) SetUsageRequest(usageRequest *UsageRequest) error {
    r.usageRequest = usageRequest
    r.Set("usage_request", usageRequest)
    return nil
}

func (r AlibabaAlihealthOutflowUsageSaveorupdateRequest) GetUsageRequest() *UsageRequest {
    return r.usageRequest
}

