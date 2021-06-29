package alihealthmedical

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三方医生消息写入 APIResponse
alibaba.alihealth.medical.doctor.msg.send

三方机构医生端发送消息同步写入阿里健康IM
*/
type AlibabaAlihealthMedicalDoctorMsgSendAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalDoctorMsgSendResponse
}

type AlibabaAlihealthMedicalDoctorMsgSendResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_doctor_msg_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
