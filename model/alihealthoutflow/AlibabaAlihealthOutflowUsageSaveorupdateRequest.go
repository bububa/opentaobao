package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-用法字典表 API请求
alibaba.alihealth.outflow.usage.saveorupdate

阿里健康-处方外流-对外提供用法字典表维护功能
*/
type AlibabaAlihealthOutflowUsageSaveorupdateRequest struct {
    model.Params
    // 用法数据
    usageRequest   *UsageRequest
}

// 初始化AlibabaAlihealthOutflowUsageSaveorupdateRequest对象
func NewAlibabaAlihealthOutflowUsageSaveorupdateRequest() *AlibabaAlihealthOutflowUsageSaveorupdateRequest{
    return &AlibabaAlihealthOutflowUsageSaveorupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowUsageSaveorupdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.usage.saveorupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowUsageSaveorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UsageRequest Setter
// 用法数据
func (r *AlibabaAlihealthOutflowUsageSaveorupdateRequest) SetUsageRequest(usageRequest *UsageRequest) error {
    r.usageRequest = usageRequest
    r.Set("usage_request", usageRequest)
    return nil
}

// UsageRequest Getter
func (r AlibabaAlihealthOutflowUsageSaveorupdateRequest) GetUsageRequest() *UsageRequest {
    return r.usageRequest
}
