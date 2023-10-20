package alihealthcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse 互联网医院批量导入接口 API返回值
// alibaba.alihealth.medicalbase.hospital.sync
//
// 互联网医院isv批量通过接口批量导入
type AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalbaseHospitalSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalbaseHospitalSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalbaseHospitalSyncAPIResponseModel is 互联网医院批量导入接口 成功返回结果
type AlibabaAlihealthMedicalbaseHospitalSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_hospital_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseHospitalSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalbaseHospitalSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalbaseHospitalSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse
func GetAlibabaAlihealthMedicalbaseHospitalSyncAPIResponse() *AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse {
	return poolAlibabaAlihealthMedicalbaseHospitalSyncAPIResponse.Get().(*AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalbaseHospitalSyncAPIResponse 将 AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseHospitalSyncAPIResponse(v *AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseHospitalSyncAPIResponse.Put(v)
}
