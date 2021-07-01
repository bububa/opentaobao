package alihealthoutflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-患者基础信息同步 API返回值 
alibaba.alihealth.outflow.patientinfo.sync

阿里健康-处方外流-对外提供同步患者基础信息功能
*/
type AlibabaAlihealthOutflowPatientinfoSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthOutflowPatientinfoSyncAPIResponseModel
}

// 处方外流-患者基础信息同步 成功返回结果
type AlibabaAlihealthOutflowPatientinfoSyncAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_outflow_patientinfo_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // ServiceResult
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
