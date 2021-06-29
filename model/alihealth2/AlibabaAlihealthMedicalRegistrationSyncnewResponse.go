package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康新挂号数据回传 APIResponse
alibaba.alihealth.medical.registration.syncnew

阿里健康新挂号记录回传接口
*/
type AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalRegistrationSyncnewResponse
}

type AlibabaAlihealthMedicalRegistrationSyncnewResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_registration_syncnew_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
