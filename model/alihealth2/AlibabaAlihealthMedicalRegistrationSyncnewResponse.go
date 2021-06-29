package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康新挂号数据回传 API返回值 
alibaba.alihealth.medical.registration.syncnew

阿里健康新挂号记录回传接口
*/
type AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalRegistrationSyncnewResponse
}

// 阿里健康新挂号数据回传 成功返回结果
type AlibabaAlihealthMedicalRegistrationSyncnewResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_registration_syncnew_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
