package alihealth

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthprescriptionauthgetAPIResponse 阿里健康处方平台获取授权码 API返回值
// alibaba.alihealth.prescription.auth.get
//
// 获取处方用户授权
type AlibabaalihealthprescriptionauthgetAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthprescriptionauthgetAPIResponseModel
}

// AlibabaalihealthprescriptionauthgetAPIResponseModel is 阿里健康处方平台获取授权码 成功返回结果
type AlibabaalihealthprescriptionauthgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_prescription_auth_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
