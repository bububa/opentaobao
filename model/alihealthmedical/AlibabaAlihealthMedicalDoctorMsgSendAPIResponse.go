package alihealthmedical

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalDoctorMsgSendAPIResponse 三方医生消息写入 API返回值
// alibaba.alihealth.medical.doctor.msg.send
//
// 三方机构医生端发送消息同步写入阿里健康IM
type AlibabaAlihealthMedicalDoctorMsgSendAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalDoctorMsgSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalDoctorMsgSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalDoctorMsgSendAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalDoctorMsgSendAPIResponseModel is 三方医生消息写入 成功返回结果
type AlibabaAlihealthMedicalDoctorMsgSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_doctor_msg_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalDoctorMsgSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalDoctorMsgSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalDoctorMsgSendAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalDoctorMsgSendAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalDoctorMsgSendAPIResponse
func GetAlibabaAlihealthMedicalDoctorMsgSendAPIResponse() *AlibabaAlihealthMedicalDoctorMsgSendAPIResponse {
	return poolAlibabaAlihealthMedicalDoctorMsgSendAPIResponse.Get().(*AlibabaAlihealthMedicalDoctorMsgSendAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalDoctorMsgSendAPIResponse 将 AlibabaAlihealthMedicalDoctorMsgSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalDoctorMsgSendAPIResponse(v *AlibabaAlihealthMedicalDoctorMsgSendAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalDoctorMsgSendAPIResponse.Put(v)
}
