package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
微医数据号源回传 APIResponse
alibaba.alihealth.medical.register.weiyi.sync

微医号源数据回传
*/
type AlibabaAlihealthMedicalRegisterWeiyiSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalRegisterWeiyiSyncResponse
}

type AlibabaAlihealthMedicalRegisterWeiyiSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_register_weiyi_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和三方交互最外层model对象
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
