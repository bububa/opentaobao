package alihealthoutflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-用法字典表 API返回值 
alibaba.alihealth.outflow.usage.saveorupdate

阿里健康-处方外流-对外提供用法字典表维护功能
*/
type AlibabaAlihealthOutflowUsageSaveorupdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthOutflowUsageSaveorupdateAPIResponseModel
}

// 处方外流-用法字典表 成功返回结果
type AlibabaAlihealthOutflowUsageSaveorupdateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_outflow_usage_saveorupdate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // ServiceResult
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
