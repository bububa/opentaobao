package alihealthmedical

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三方医生消息写入 API返回值 
alibaba.alihealth.medical.doctor.msg.send

三方机构医生端发送消息同步写入阿里健康IM
*/
type AlibabaAlihealthMedicalDoctorMsgSendAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalDoctorMsgSendAPIResponseModel
}

// 三方医生消息写入 成功返回结果
type AlibabaAlihealthMedicalDoctorMsgSendAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_doctor_msg_send_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
