package alihealthoutflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-问诊、处方同步 API返回值 
alibaba.alihealth.outflow.visitinfo.sync

阿里健康-处方外流-对外提供问诊、处方功能
*/
type AlibabaAlihealthOutflowVisitinfoSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthOutflowVisitinfoSyncAPIResponseModel
}

// 处方外流-问诊、处方同步 成功返回结果
type AlibabaAlihealthOutflowVisitinfoSyncAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_outflow_visitinfo_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // ServiceResult
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
