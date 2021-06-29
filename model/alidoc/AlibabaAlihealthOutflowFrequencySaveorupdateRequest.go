package alidoc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-药品频次同步接口 APIRequest
alibaba.alihealth.outflow.frequency.saveorupdate

处方外流-药品频次同步接口
*/
type AlibabaAlihealthOutflowFrequencySaveorupdateRequest struct {
    model.Params

    // 系统自动生成
    frequencyRequest   *FrequencyRequest 

}

func NewAlibabaAlihealthOutflowFrequencySaveorupdateRequest() *AlibabaAlihealthOutflowFrequencySaveorupdateRequest{
    return &AlibabaAlihealthOutflowFrequencySaveorupdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowFrequencySaveorupdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.frequency.saveorupdate"
}

func (r AlibabaAlihealthOutflowFrequencySaveorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowFrequencySaveorupdateRequest) SetFrequencyRequest(frequencyRequest *FrequencyRequest) error {
    r.frequencyRequest = frequencyRequest
    r.Set("frequency_request", frequencyRequest)
    return nil
}

func (r AlibabaAlihealthOutflowFrequencySaveorupdateRequest) GetFrequencyRequest() *FrequencyRequest {
    return r.frequencyRequest
}

