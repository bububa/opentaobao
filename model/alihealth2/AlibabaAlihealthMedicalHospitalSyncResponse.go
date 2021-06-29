package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里将康支付宝挂号数据医院回传接口 APIResponse
alibaba.alihealth.medical.hospital.sync

阿里将康支付宝挂号数据医院回传接口
*/
type AlibabaAlihealthMedicalHospitalSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalHospitalSyncResponse
}

type AlibabaAlihealthMedicalHospitalSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_hospital_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
