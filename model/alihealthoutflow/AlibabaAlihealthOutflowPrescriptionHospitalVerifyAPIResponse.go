package alihealthoutflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse 处方同步至医院返回校验结果 API返回值
// alibaba.alihealth.outflow.prescription.hospital.verify
//
// 处方同步至医院返回校验结果
type AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponseModel).Reset()
}

// AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponseModel is 处方同步至医院返回校验结果 成功返回结果
type AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_outflow_prescription_hospital_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServiceResult
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse)
	},
}

// GetAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse 从 sync.Pool 获取 AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse
func GetAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse() *AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse {
	return poolAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse.Get().(*AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse)
}

// ReleaseAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse 将 AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse(v *AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse.Put(v)
}
