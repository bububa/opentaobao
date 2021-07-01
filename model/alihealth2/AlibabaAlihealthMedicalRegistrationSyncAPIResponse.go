package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康支付宝挂号记录回传接口 API返回值 
alibaba.alihealth.medical.registration.sync

阿里健康支付宝挂号记录回传接口
*/
type AlibabaAlihealthMedicalRegistrationSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalRegistrationSyncAPIResponseModel
}

// 阿里健康支付宝挂号记录回传接口 成功返回结果
type AlibabaAlihealthMedicalRegistrationSyncAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_registration_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
