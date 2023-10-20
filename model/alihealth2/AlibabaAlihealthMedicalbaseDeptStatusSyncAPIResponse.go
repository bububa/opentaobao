package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse 挂号科室上下线 API返回值
// alibaba.alihealth.medicalbase.dept.status.sync
//
// 挂号医院上下线
type AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponseModel is 挂号科室上下线 成功返回结果
type AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_dept_status_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse
func GetAlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse() *AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse {
	return poolAlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse.Get().(*AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse 将 AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse(v *AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse.Put(v)
}
