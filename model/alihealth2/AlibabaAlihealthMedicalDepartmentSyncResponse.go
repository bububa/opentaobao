package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康预约挂号科室同步接口 APIResponse
alibaba.alihealth.medical.department.sync

阿里健康预约挂号科室同步接口
*/
type AlibabaAlihealthMedicalDepartmentSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalDepartmentSyncResponse
}

type AlibabaAlihealthMedicalDepartmentSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_department_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
