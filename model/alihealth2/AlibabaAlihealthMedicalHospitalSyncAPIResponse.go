package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalHospitalSyncAPIResponse 阿里将康支付宝挂号数据医院回传接口 API返回值
// alibaba.alihealth.medical.hospital.sync
//
// 阿里将康支付宝挂号数据医院回传接口
type AlibabaAlihealthMedicalHospitalSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalHospitalSyncAPIResponseModel
}

// AlibabaAlihealthMedicalHospitalSyncAPIResponseModel is 阿里将康支付宝挂号数据医院回传接口 成功返回结果
type AlibabaAlihealthMedicalHospitalSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_hospital_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
