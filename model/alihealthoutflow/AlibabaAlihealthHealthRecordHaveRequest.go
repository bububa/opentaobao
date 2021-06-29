package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断用户的慢健康健康档案是否建设完成 APIRequest
alibaba.alihealth.health.record.have

判断用户的慢健康健康档案是否建设完成
*/
type AlibabaAlihealthHealthRecordHaveRequest struct {
    model.Params

    // 入参
    request1   *HaveRecordRequest 

}

func NewAlibabaAlihealthHealthRecordHaveRequest() *AlibabaAlihealthHealthRecordHaveRequest{
    return &AlibabaAlihealthHealthRecordHaveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthHealthRecordHaveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.health.record.have"
}

func (r AlibabaAlihealthHealthRecordHaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthHealthRecordHaveRequest) SetRequest1(request1 *HaveRecordRequest) error {
    r.request1 = request1
    r.Set("request1", request1)
    return nil
}

func (r AlibabaAlihealthHealthRecordHaveRequest) GetRequest1() *HaveRecordRequest {
    return r.request1
}

