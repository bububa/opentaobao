package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断用户的慢健康健康档案是否建设完成 API请求
alibaba.alihealth.health.record.have

判断用户的慢健康健康档案是否建设完成
*/
type AlibabaAlihealthHealthRecordHaveRequest struct {
    model.Params
    // 入参
    _request1   *HaveRecordRequest
}

// 初始化AlibabaAlihealthHealthRecordHaveRequest对象
func NewAlibabaAlihealthHealthRecordHaveRequest() *AlibabaAlihealthHealthRecordHaveRequest{
    return &AlibabaAlihealthHealthRecordHaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthHealthRecordHaveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.health.record.have"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthHealthRecordHaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request1 Setter
// 入参
func (r *AlibabaAlihealthHealthRecordHaveRequest) SetRequest1(_request1 *HaveRecordRequest) error {
    r._request1 = _request1
    r.Set("request1", _request1)
    return nil
}

// Request1 Getter
func (r AlibabaAlihealthHealthRecordHaveRequest) GetRequest1() *HaveRecordRequest {
    return r._request1
}