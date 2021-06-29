package alidoc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-药品频次同步接口 API请求
alibaba.alihealth.outflow.frequency.saveorupdate

处方外流-药品频次同步接口
*/
type AlibabaAlihealthOutflowFrequencySaveorupdateRequest struct {
    model.Params
    // 系统自动生成
    frequencyRequest   *FrequencyRequest
}

// 初始化AlibabaAlihealthOutflowFrequencySaveorupdateRequest对象
func NewAlibabaAlihealthOutflowFrequencySaveorupdateRequest() *AlibabaAlihealthOutflowFrequencySaveorupdateRequest{
    return &AlibabaAlihealthOutflowFrequencySaveorupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowFrequencySaveorupdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.frequency.saveorupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowFrequencySaveorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FrequencyRequest Setter
// 系统自动生成
func (r *AlibabaAlihealthOutflowFrequencySaveorupdateRequest) SetFrequencyRequest(frequencyRequest *FrequencyRequest) error {
    r.frequencyRequest = frequencyRequest
    r.Set("frequency_request", frequencyRequest)
    return nil
}

// FrequencyRequest Getter
func (r AlibabaAlihealthOutflowFrequencySaveorupdateRequest) GetFrequencyRequest() *FrequencyRequest {
    return r.frequencyRequest
}
