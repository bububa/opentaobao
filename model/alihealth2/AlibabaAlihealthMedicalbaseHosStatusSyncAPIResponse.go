package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse 挂号医院上下线 API返回值
// alibaba.alihealth.medicalbase.hos.status.sync
//
// 挂号医院上下线
type AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponseModel is 挂号医院上下线 成功返回结果
type AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_hos_status_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse
func GetAlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse() *AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse {
	return poolAlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse.Get().(*AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse 将 AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse(v *AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse.Put(v)
}
