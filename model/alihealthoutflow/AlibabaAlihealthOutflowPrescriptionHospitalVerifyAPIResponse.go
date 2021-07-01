package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse
处方同步至医院返回校验结果 API返回值
alibaba.alihealth.outflow.prescription.hospital.verify

处方同步至医院返回校验结果 */
type AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponseModel
}

// AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponseModel is 处方同步至医院返回校验结果 成功返回结果
type AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_outflow_prescription_hospital_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServiceResult
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
