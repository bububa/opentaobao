package alihealthoutflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
判断用户的慢健康健康档案是否建设完成 API返回值 
alibaba.alihealth.health.record.have

判断用户的慢健康健康档案是否建设完成
*/
type AlibabaAlihealthHealthRecordHaveAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthHealthRecordHaveResponse
}

// 判断用户的慢健康健康档案是否建设完成 成功返回结果
type AlibabaAlihealthHealthRecordHaveResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_health_record_have_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result1   *AlibabaAlihealthHealthRecordHaveResult `json:"result1,omitempty" xml:"result1,omitempty"`
}
