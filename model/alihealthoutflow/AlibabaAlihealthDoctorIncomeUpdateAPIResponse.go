package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDoctorIncomeUpdateAPIResponse 医蝶谷医生收入打款情况回调 API返回值
// alibaba.alihealth.doctor.income.update
//
// 医蝶谷医生收入打款情况回调
type AlibabaAlihealthDoctorIncomeUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDoctorIncomeUpdateAPIResponseModel
}

// AlibabaAlihealthDoctorIncomeUpdateAPIResponseModel is 医蝶谷医生收入打款情况回调 成功返回结果
type AlibabaAlihealthDoctorIncomeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_doctor_income_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ExceptionCode string `json:"exception_code,omitempty" xml:"exception_code,omitempty"`
	// 错误信息
	ExceptionMessage string `json:"exception_message,omitempty" xml:"exception_message,omitempty"`
	// 接口是否成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
