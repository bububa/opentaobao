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
type AlibabaAlihealthOutflowUsageSaveorupdateAPIRequest struct {
    model.Params
    // 用法数据
    _usageRequest   *UsageRequest
}

// 初始化AlibabaAlihealthOutflowUsageSaveorupdateAPIRequest对象
func NewAlibabaAlihealthOutflowUsageSaveorupdateRequest() *AlibabaAlihealthOutflowUsageSaveorupdateAPIRequest{
    return &AlibabaAlihealthOutflowUsageSaveorupdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowUsageSaveorupdateAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.usage.saveorupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowUsageSaveorupdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UsageRequest Setter
// 用法数据
func (r *AlibabaAlihealthOutflowUsageSaveorupdateAPIRequest) SetUsageRequest(_usageRequest *UsageRequest) error {
    r._usageRequest = _usageRequest
    r.Set("usage_request", _usageRequest)
    return nil
}

// UsageRequest Getter
func (r AlibabaAlihealthOutflowUsageSaveorupdateAPIRequest) GetUsageRequest() *UsageRequest {
    return r._usageRequest
}
