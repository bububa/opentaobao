package alihealthoutflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse 申请单审核结果通知 API返回值
// alibaba.alihealth.doctor.leshui.apply.notify
//
// 申请单审核结果通知
type AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponseModel).Reset()
}

// AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponseModel is 申请单审核结果通知 成功返回结果
type AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_doctor_leshui_apply_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse)
	},
}

// GetAlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse
func GetAlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse() *AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse {
	return poolAlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse.Get().(*AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse)
}

// ReleaseAlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse 将 AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse(v *AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse.Put(v)
}
