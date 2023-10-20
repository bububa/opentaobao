package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalDepartmentSyncAPIResponse 阿里健康预约挂号科室同步接口 API返回值
// alibaba.alihealth.medical.department.sync
//
// 阿里健康预约挂号科室同步接口
type AlibabaAlihealthMedicalDepartmentSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalDepartmentSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalDepartmentSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalDepartmentSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalDepartmentSyncAPIResponseModel is 阿里健康预约挂号科室同步接口 成功返回结果
type AlibabaAlihealthMedicalDepartmentSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_department_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalDepartmentSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalDepartmentSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalDepartmentSyncAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalDepartmentSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalDepartmentSyncAPIResponse
func GetAlibabaAlihealthMedicalDepartmentSyncAPIResponse() *AlibabaAlihealthMedicalDepartmentSyncAPIResponse {
	return poolAlibabaAlihealthMedicalDepartmentSyncAPIResponse.Get().(*AlibabaAlihealthMedicalDepartmentSyncAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalDepartmentSyncAPIResponse 将 AlibabaAlihealthMedicalDepartmentSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalDepartmentSyncAPIResponse(v *AlibabaAlihealthMedicalDepartmentSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalDepartmentSyncAPIResponse.Put(v)
}
